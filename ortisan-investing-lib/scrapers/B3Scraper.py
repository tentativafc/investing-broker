# coding=utf-8
__author__ = 'Marcelo Ortiz'

import glob
import os
import shutil
import urllib.request as req
import zipfile
from datetime import date
from urllib.error import HTTPError

import os.path

import configs
from utils.DateUtils import DateUtils
from utils.LogUtils import LogUtils
from utils.FileUtils import FileUtils


class B3Scraper(object):
    MAX_LINES_FILES = 1000

    logger = LogUtils.getInstance()

    def import_by_year(self, year, override=False):
        self.logger.info("Importing B3 Quotations of year %d...", year)
        fileName = 'COTAHIST_A{}.zip'.format(year)
        self.__import(fileName, date, override)

    def import_by_moth_year(self, month, year, override=False):
        self.logger.info("Importing B3 Quotations of year %d and month %d...", year, month)
        fileName = 'COTAHIST_M%02d%04d.zip' % (month, year)
        self.__import(fileName, date, override)

    def import_by_date(self, date, override=False):
        self.logger.info("Importing B3 Quotations of date %s...", date)
        fileName = date.strftime('COTAHIST_D%d%m%Y.zip')
        self.__import(fileName, date, override)

    def __import(self, fileName, date, override=False, retries=0):
        if retries > 2:
            return
        try:
            fileOfDate = req.urlopen('{}/{}'.format(configs.get_bovespa_url(), fileName))
            destFile = '{}{}'.format(configs.get_temp_data_dir(), fileName)
            if not override and os.path.exists(destFile):
                self.logger.info("File already exists.")
                return
            with open(destFile, 'wb') as output:
                output.write(fileOfDate.read())
        except HTTPError as exc:
            self.logger.error("Data for date %s do not exists...", date, exc_info=True)
        except Exception as exc:
            self.logger.error("General error. Retrying...", exc_info=True)
            self.__import(date, override, retries + 1)

    def import_by_period(self, initial_date, final_date, override=False):
        dates = DateUtils.get_dates(initial_date, final_date)
        for date in dates:
            self.import_by_date(date)

    def unzip_scraped_files(self):
        dir_bovespa_zip_files = '{0}/COTAHIST*.zip'.format(configs.get_temp_data_dir())
        for file in glob.glob(dir_bovespa_zip_files):
            zip_ref = zipfile.ZipFile(file, 'r')
            zip_ref.extractall(configs.get_raw_data_dir())
            zip_ref.close()
            shutil.copy2(file, configs.get_backup_data_dir())
            os.remove(file)

    def slice_files(self):
        FileUtils.slice_files('{0}/data/bovespa/raw/COTAHIST_*.TXT'.format(configs.get_project_home()),
                              self.MAX_LINES_FILES, '{0}/data/bovespa/slices'.format(configs.get_project_home()),
                              file_encoding='latin-1')


if __name__ == '__main__':
    b3Scraper = B3Scraper()
    for year in range(2007, 2020):
        b3Scraper.import_by_year(year)
    # initial_date = date(2019, 1, 25)
    # final_date = date(2020, 2, 5)
    # b3Scraper.import_by_period(initial_date, final_date)
    b3Scraper.unzip_scraped_files()
    b3Scraper.slice_files()
