# coding=utf-8
__author__ = 'Marcelo Ortiz'

from abc import ABCMeta, abstractmethod

import numpy as np
import pandas as pd
from model.Daos import AssetQuotationDao, CoinQuotationDao
from utils.DateUtils import DateUtils


class DataSource(object):
    POSTGRES = 'Postgres'


class DataItem(object):
    DATE = 'Date'
    OPEN = 'Open'
    CLOSE = 'Close'
    HIGH = 'High'
    LOW = 'Low'
    VOLUME = 'Volume'
    ADJUSTED_CLOSE = 'Adj Close'

    @staticmethod
    def get_list():
        return [DataItem.OPEN, DataItem.CLOSE, DataItem.HIGH, DataItem.LOW, DataItem.ADJUSTED_CLOSE, DataItem.VOLUME]


class DataAccess(object):
    def __init__(self, datasource=DataSource.POSTGRES):
        self.datasource = datasource
        self.data_provider = PostgresDataProvider()

    @staticmethod
    def get_yahoo_dtypes():
        return {DataItem.OPEN: np.float64,
                DataItem.HIGH: np.float64,
                DataItem.LOW: np.float64,
                DataItem.CLOSE: np.float64,
                DataItem.ADJUSTED_CLOSE: np.float64,
                DataItem.VOLUME: np.float64}

    def get_symbols(self, *args, **kwargs):
        return self.data_provider.get_symbols(*args, **kwargs)

    def get_data(self, symbol_list, initial_date, final_date):
        return self.data_provider.get_data_by_period(symbol_list, initial_date, final_date)

    def get_coin_data(self, symbol_list, initial_date, final_date):
        return self.data_provider.get_coin_data_by_period(symbol_list, initial_date, final_date)


class __AbstractDataProvider:
    __metaclass__ = ABCMeta

    @abstractmethod
    def get_symbols(self, *args, **kwargs):
        pass

    @abstractmethod
    def get_data_by_period(self, symbol_list, initial_date, final_date, data_item_list):
        pass

    @abstractmethod
    def get_coin_data_by_period(self, symbol_list, initial_date, final_date):
        pass

class PostgresDataProvider(__AbstractDataProvider):

    def get_symbols(self):
        dao = AssetQuotationDao()
        return dao.get_symbols()

    def get_data_by_period(self, symbol_list, initial_date, final_date):

        columns = []
        columns.append(DataItem.DATE)
        columns.append(DataItem.OPEN)
        columns.append(DataItem.HIGH)
        columns.append(DataItem.LOW)
        columns.append(DataItem.CLOSE)
        columns.append(DataItem.ADJUSTED_CLOSE)
        columns.append(DataItem.VOLUME)

        dict_data = {}
        dao = AssetQuotationDao()

        for symbol in symbol_list:
            quotations = dao.get_by_asset_and_date(symbol, initial_date, final_date)
            data = [(quotation.date, quotation.open_price, quotation.max_price, quotation.min_price,
                     quotation.close_price, quotation.close_price, quotation.volume) for quotation in quotations]

            df_symbol_database = pd.DataFrame(data, columns=columns)
            df_symbol_database.set_index(columns[0], inplace=True)

            dict_data[symbol] = df_symbol_database

        return dict_data

    def get_coin_data_by_period(self, coin_list, initial_date, final_date):
        dao = CoinQuotationDao()

        date_list = DateUtils.get_dates(initial_date, final_date)

        columns = []
        columns.append(DataItem.OPEN)
        columns.append(DataItem.HIGH)
        columns.append(DataItem.LOW)
        columns.append(DataItem.CLOSE)
        columns.append(DataItem.ADJUSTED_CLOSE)
        columns.append(DataItem.VOLUME)

        # Create an empty df
        empty_df = pd.DataFrame(index=pd.to_datetime(date_list, utc=False), columns=columns)

        columns.insert(0, DataItem.DATE)

        dict_data = {}

        for coin in coin_list:
            quotations = dao.get_by_symbol_and_date(coin, initial_date, final_date)
            data = [(quotation.date, quotation.open_price, quotation.max_price, quotation.min_price,
                     quotation.close_price, quotation.close_price, quotation.volume) for quotation in quotations]

            df_coin = pd.DataFrame(data, columns=columns)

            # Index is the date
            df_coin.set_index([DataItem.DATE], inplace=True)

            df_result = empty_df.copy()

            df_result.update(df_coin)

            dict_data[coin] = df_result

        return dict_data


if __name__ == '__main__':
    provider = PostgresDataProvider()
    from datetime import datetime

    data_access = DataAccess(datasource=DataSource.POSTGRES)

    symbols = data_access.get_symbols()
    print(symbols)

    result = data_access.get_data(['AALR3'], datetime(2016, 10, 28, 0, 0, 0, 0),
                                  datetime(2018, 10, 31, 0, 0, 0, 0))
    print(result['AALR3'])

    result = data_access.get_data('AALR3', datetime(2016, 10, 28, 0, 0, 0, 0), datetime(2018, 10, 28, 0, 0, 0, 0))
    print(result)
