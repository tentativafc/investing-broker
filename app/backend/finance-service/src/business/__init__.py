# coding=utf-8
__author__ = 'Marcelo Ortiz'

import MetaTrader5 as mt5
from ortisan_ta.dataaccess import DataItem, MetaTraderDataAccess
import ortisan_ta.utils.analysis as ortisan_ta
import pandas as pd
import random
from datetime import datetime
from errors import IllegalArgumentException
import uuid

from models import AssetPortfolio, Portfolio

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

    def getQuotes(self, symbol: str, init_date: datetime, final_date: datetime):
        df = self.data_access.get_rates_from_symbol(symbol, init_date, final_date, mt5.TIMEFRAME_H1)
        # to_json() does not serialize the dataframe index
        df_with_date = df.reset_index()
        df_with_date['Date'] = df.index.astype(str)
        list_dict = []
        for index, row in list(df_with_date.iterrows()):
            list_dict.append(dict(row))
        return list_dict

    def createPortfolio(self, amount_assets: int):
        if amount_assets > 10:
            raise IllegalArgumentException("Max amount of assets by portifolio is 10.")
        # create a dummy Portfolio
        # todo use optimization algorithm to create
        #Generate 5 random numbers between 10 and 30
        randomlist = random.sample(range(0, len(TOP_50_ASSETS_IBOVESPA)), amount_assets)
        assets = [AssetPortfolio(symbol=TOP_50_ASSETS_IBOVESPA[i]) for i in randomlist]
        
        portfolio = Portfolio(user_id=str(uuid.uuid4()), assets=assets, capm=1.0, beta=1.0)
        portfolio.save()
        
        return portfolio.to_mongo().to_dict()


