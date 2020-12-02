# coding=utf-8
__author__ = 'Marcelo Ortiz'

import pandas as pd

class DateUtils:
    @staticmethod
    def get_dates(initial_date, final_date, remove_weekend=True):
        dates = pd.date_range(initial_date, final_date)
        if remove_weekend:
            dates = dates[dates.weekday < 5]
        return dates


from datetime import date

if __name__ == '__main__':
    initial_date = date(2018, 2, 1)
    final_date = date(2018, 3, 1)
    dates = DateUtils.get_dates(initial_date, final_date)
    print(dates)
