# coding=utf-8
__author__ = 'Marcelo Ortiz'

import glob

import pandas as pd
from rx import Observer

import configs
from model.Daos import AssetQuotationDao
from model.Models import AssetQuotation
from utils.FileUtils import FileUtils
from utils.RxUtils import RxUtils
from utils.LogUtils import LogUtils
from dataaccess.DataAccess import DataAccess, DataSource
from datetime import datetime

TIPO_MERCADO_VISTA = 10

class B3AssetsDataImporter(Observer):

    columns = ['TIPREG',
               'DTPREG',
               'CODBDI',
               'CODNEG',
               'TPMERC',
               'NOMRES',
               'ESPECI',
               'PRAZOT',
               'MODREF',
               'PREABE',
               'PREMAX',
               'PREMIN',
               'PREMED',
               'PREULT',
               'PREOFC',
               'PREOFV',
               'TOTNEG',
               'QUATOT',
               'VOLTOT',
               'PREEXE',
               'INDOPC',
               'DATVEN',
               'FATCOT',
               'PTOEXE',
               'CODISI',
               'DISMES']

    colspecs = [(0, 2),
                (2, 10),
                (10, 12),
                (12, 24),
                (24, 27),
                (27, 39),
                (39, 49),
                (49, 52),
                (52, 56),
                (56, 69),
                (69, 82),
                (82, 95),
                (95, 108),
                (108, 121),
                (121, 134),
                (134, 147),
                (147, 152),
                (152, 170),
                (170, 188),
                (188, 201),
                (201, 202),
                (202, 210),
                (210, 217),
                (217, 230),
                (230, 242),
                (242, 245)]

    logger = LogUtils.getInstance()

    def __init__(self):
        pass

    def on_next(self, file):
        self.logger.info("Processing file:{0}".format(file))
        self.process_file(file)

    def on_error(self, error):
        LogUtils.getInstance().error("Error", error)

    def on_completed(self):
        LogUtils.getInstance().info("Completed")

    def process(self):
        dir_bovespa_files = '{0}/data/bovespa/slices/COTAHIST_*'.format(configs.get_project_home())
        observable = RxUtils.get_multithread_observable(glob.glob(dir_bovespa_files))
        observable.subscribe(self)

    def process_file(self, file):
        self.logger.info("Processing file %s ...", file)
        bovespa_df = pd.read_fwf(file, names=self.columns, colspecs=self.colspecs, header=None, infer_nrows=0,
                                 parse_dates=['DTPREG'], date_parser=lambda dt: pd.datetime.strptime(dt, '%Y%m%d'))

        bovespa_df['CODBDI'] = bovespa_df['CODBDI'].astype('str')
        bovespa_df['PRAZOT'] = bovespa_df['PRAZOT'].astype('str')
        bovespa_df['PREABE'] = bovespa_df['PREABE'] / 1e2
        bovespa_df['PREMAX'] = bovespa_df['PREMAX'] / 1e2
        bovespa_df['PREMIN'] = bovespa_df['PREMIN'] / 1e2
        bovespa_df['PREMED'] = bovespa_df['PREMED'] / 1e2
        bovespa_df['PREULT'] = bovespa_df['PREULT'] / 1e2
        bovespa_df['PREOFC'] = bovespa_df['PREOFC'] / 1e2
        bovespa_df['PREOFV'] = bovespa_df['PREOFV'] / 1e2
        bovespa_df['VOLTOT'] = bovespa_df['VOLTOT'] / 1e2
        bovespa_df['PREEXE'] = bovespa_df['PREEXE'] / 1e6

        list_to_insert = []

        try:
            for index, row in bovespa_df.iterrows():
                code = row['CODNEG']
                market_type = row['TPMERC']
                cod_bdi = row['CODBDI']
                asset_id = '{0}_{1}_{2}'.format(market_type, cod_bdi, code)
                date_quote = row['DTPREG']
                days_term_market = row['PRAZOT']
                days_term_market = 0 if days_term_market == 'nan' else int(float(days_term_market))
                min_price = row['PREMIN']
                max_price = row['PREMAX']
                open_price = row['PREABE']
                close_price = row['PREULT']
                volume = row['VOLTOT']
                quotation_id = '{0}_{1}_{2}'.format(date_quote, asset_id, days_term_market)
                quotation = AssetQuotation(id=quotation_id,
                                        date=date_quote,
                                        symbol=code,
                                        market_type=market_type,
                                        bdi_code=cod_bdi,
                                        days_term_market=days_term_market,
                                        min_price=min_price,
                                        max_price=max_price,
                                        open_price=open_price,
                                        close_price=close_price,
                                        volume=volume)
                list_to_insert.append(quotation)

        
            asset_quotation_dao = AssetQuotationDao()
            asset_quotation_dao.insert(list_to_insert)
            FileUtils.delete_files(file)
        except Exception as exc:
            self.logger.error("Error to inserts rows of file: %s", file, exc_info=True)
            pass

    def transform_to_yahoo_format(self):
        data_access = DataAccess(datasource=DataSource.POSTGRES)
        symbols = data_access.get_symbols()
        for symbol in symbols:
            self.logger.info("Generating dataframes for symbol %s", symbol)
            data_frames_dict = data_access.get_data([symbol], datetime(2012, 1, 1, 0, 0, 0, 0),
                                                    datetime(2019, 1, 1, 0, 0, 0, 0))
            for symbol, dataframe in data_frames_dict.items():
                path_dest = '{}{}.csv'.format(configs.get_bovespa_yahoo_dir(), symbol)
                dataframe.to_csv(path_dest, index=False)


if __name__ == '__main__':
    b3Importer = B3AssetsDataImporter()
    b3Importer.process()

    # bovespa_importer.transform_to_yahoo_format()
