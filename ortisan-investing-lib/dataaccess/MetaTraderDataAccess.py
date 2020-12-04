import MetaTrader5 as mt5
import pandas as pd
from datetime import datetime
from utils.FileUtils import FileUtils

class MetaTraderDataAccess:
    def __init__(self):
        if not mt5.initialize():
            print("initialize() failed, error code =", mt5.last_error())
            raise Exception("Error to connect to Metatrader.", mt5.last_error())

    def get_rates_from_symbol(self, symbol, date_from, date_to, timeframe=mt5.TIMEFRAME_M5):
        rates = mt5.copy_rates_range(symbol, timeframe, date_from, date_to)
        df = pd.DataFrame(rates)
        if df.empty:
            return df
        df.time = df.time.transform([datetime.fromtimestamp])
        return df


if __name__ == '__main__':
    data_access = MetaTraderDataAccess()
    asset = 'WINM20'
    df = data_access.get_rates_from_symbol(asset, datetime(2020, 4, 10), datetime(2020, 5, 20), mt5.TIMEFRAME_M5)
    FileUtils.save_to_pickle(df, asset)




