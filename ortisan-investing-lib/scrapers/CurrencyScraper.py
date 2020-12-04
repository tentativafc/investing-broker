# coding=utf-8
__author__ = 'Marcelo Ortiz'

import urllib.request as req
import configs
from urllib.error import HTTPError
import os


class CurrencyScraper(object):
    CURRENCIES = [
        (61, 'Dólar'),
        (222, 'Euro'),
        (101, 'Iene'),
        (178, 'Iuan'),
    ]

    def import_by_period(self, initial_date, final_date, override=False, retries=0):
        print("Importing currenties for dates %s - %s ..." % (initial_date, final_date))

        if retries > 2:
            return
        for currency in self.CURRENCIES:
            try:
                url = configs.get_currency_url().format(currency[0], initial_date.strftime('%d/%m/%Y'), final_date.strftime('%d/%m/%Y'))
                print("Using url %s" % url)
                fileOfDate = req.urlopen(url)
                destFile = '{}{}.csv'.format(configs.get_raw_data_dir('currency'), currency[1])
                if not override and os.path.exists(destFile):
                    raise Exception('File already exists.')
                    return
                with open(destFile, 'wb') as output:
                    output.write(fileOfDate.read())
            except HTTPError as exc:
                print('Data for currency "%s" and dates "%s" and "%s" do not exists...' % currency[1], initial_date,
                    final_date)

            except Exception as exc:
                print('General error. Retrying...')
                self.import_by_period(initial_date, final_date, override, retries + 1)


if __name__ == '__main__':
    import datetime

    init_date = datetime.datetime(2007, 1, 1)
    final_date = datetime.datetime(2020, 12, 3)

    scrapper = CurrencyScraper()
    scrapper.import_by_period(init_date, final_date, override=True)
