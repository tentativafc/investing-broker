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

REDIS_CORPORATES_INFO = "corporates_info"

TOP_50_ASSETS_IBOVESPA = ["ABEV3",
                          "AZUL4",
                          "B3SA3",
                          "BBAS3",
                          "BBDC4",
                          "BBSE3",
                          "BPAC11",
                          "BRDT3",
                          "BRFS3",
                          "BRML3",
                          "BTOW3",
                          "CCRO3",
                          "CIEL3",
                          "CMIG4",
                          "COGN3",
                          "CSNA3",
                          "CVCB3",
                          "CYRE3",
                          "ELET3",
                          "EQTL3",
                          "GGBR4",
                          "GNDI3",
                          "GOAU4",
                          "GOLL4",
                          "IRBR3",
                          "ITSA4",
                          "ITUB4",
                          "JBSS3",
                          "KLBN11",
                          "LAME4",
                          "LREN3",
                          "MGLU3",
                          "MRFG3",
                          "MULT3",
                          "NTCO3",
                          "PETR3",
                          "PETR4",
                          "PRIO3",
                          "RADL3",
                          "RAIL3",
                          "RENT3",
                          "SBSP3",
                          "SULA11",
                          "SUZB3",
                          "TOTS3",
                          "UGPA3",
                          "USIM5",
                          "VALE3",
                          "VVAR3",
                          "WEGE3"]


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
        init_date = datetime.now() + timedelta(days=365)

        dfs = self.data_access.get_rates_from_symbol(TOP_50_ASSETS_IBOVESPA, init_date, final_date, mt5.TIMEFRAME_D1)

        # corporates_info_str = self.cache.get(REDIS_CORPORATES_INFO)
        # corporates_info = json.loads(corporates_info_str)
        # x = {symbol: ci["setorial_classes"] for ci in corporates_info for symbol in ci["assets_code"] if len(symbol) > 0}
        return dfs

        # randomlist = random.sample(
        #     range(0, len(TOP_50_ASSETS_IBOVESPA)), amount_assets)
        # assets = [AssetPortfolio(symbol=TOP_50_ASSETS_IBOVESPA[i])
        #           for i in randomlist]

        # portfolio = Portfolio(user_id=str(uuid.uuid4()),
        #                       assets=assets, capm=1.0, beta=1.0)
        # portfolio.save()

        # portfolioDto = PortfolioDto()
        # portfolioDto.user_id = portfolio.user_id
        # portfolioDto.assets = portfolio.assets
        # portfolioDto.capm = portfolio.capm
        # portfolioDto.beta = portfolio.beta
        # portfolioDto.created_at = portfolio.created_at

        # return portfolioDto
