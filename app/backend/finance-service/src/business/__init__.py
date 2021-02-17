# coding=utf-8
__author__ = 'Marcelo Ortiz'

import json
import random
import uuid
from datetime import datetime, timedelta

import MetaTrader5 as mt5
import ortisan_ta.utils.analysis as ortisan_ta
import pandas as pd
import redis
from dtos import AssetPortfolioDto, PortfolioDto, QuoteDto
from errors import IllegalArgumentException
from models import AssetPortfolio, Portfolio
from ortisan_ta.dataaccess import DataItem, MetaTraderDataAccess
from sklearn.cluster import KMeans

class Business(object):
    def __init__(self):
        self.data_access = MetaTraderDataAccess()
        self.cache = redis.Redis(host='localhost', port=6379)

    def getQuotes(self, symbol: str, init_date: datetime, final_date: datetime):
        df = self.data_access.get_rates_from_symbol(
            symbol, init_date, final_date, mt5.TIMEFRAME_H1)
        # to_json() does not serialize the dataframe index
        df_with_date = df.reset_index()
        df_with_date['Date'] = df.index
        list_dto = []
        for index, row in list(df_with_date.iterrows()):
            quoteDto = QuoteDto()
            quoteDto.date = row.Date
            quoteDto.min_price = row.Low
            quoteDto.max_price = row.High
            quoteDto.open_price = row.Open
            quoteDto.close_price = row.Close
            quoteDto.volume = row.Volume
            list_dto.append(quoteDto)
        return list_dto

    def createPortfolio(self, amount_assets: int):
        if amount_assets > 10:
            raise IllegalArgumentException(
                "Max amount of assets by portifolio is 10.")

        final_date = datetime.now()
        init_date = datetime.now() - timedelta(days=365)

        dfs = self.data_access.get_rates_from_symbols(
            TOP_50_ASSETS_IBOVESPA, init_date, final_date, mt5.TIMEFRAME_D1)

        # Organize df where coluns are close of each asset
        dict_closes = {symbol: df.Close for (symbol, df) in dfs.items()}
        df_closes = pd.DataFrame.from_dict(dict_closes, orient='columns')
        
        # Create Analitical database
        total_return = df_closes.iloc[-1] / df_closes.iloc[0] - 1
        std = df_closes.std()
        adb = pd.DataFrame({"RY": total_return, "SY": std})

        # Kmeans create n cluster based on return and volatility
        kmeans = KMeans(n_clusters=amount_assets, init='k-means++',
                        n_init=10, random_state=24, max_iter=300)
        pred_y = kmeans.fit_predict(adb)

        # Maximum return over low volatile
        adb['Volatile_Reward'] = adb['RY']/adb['SY']
        adb['Cluster'] = pred_y
        adb.sort_values(by=['Cluster'])

        symbols = [adb[adb.Cluster == i].Volatile_Reward.idxmax()
                     for i in range(0, amount_assets)]

        assets = [AssetPortfolio(symbol=symbol) for symbol in symbols]
        
        portfolio = Portfolio(user_id=str(uuid.uuid4()),
                              assets=assets, capm=1.0, beta=1.0)
        portfolio.save()

        portfolioDto = PortfolioDto()
        portfolioDto.user_id = portfolio.user_id
        portfolioDto.assets = portfolio.assets
        portfolioDto.capm = portfolio.capm
        portfolioDto.beta = portfolio.beta
        portfolioDto.created_at = portfolio.created_at

        return portfolioDto
