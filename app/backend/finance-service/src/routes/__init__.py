# coding=utf-8
__author__ = 'Marcelo Ortiz'


from flask import request
from flask_restful import Resource, Api, reqparse, abort
from marshmallow import Schema, fields
import re

from datetime import datetime
import pandas as pd
import MetaTrader5 as mt5
from ortisan_ta.dataaccess import DataItem, MetaTraderDataAccess
import ortisan_ta.utils.analysis as ortisan_ta

class QuotesSchema(Schema):
    symbol = fields.Str(required=True)
    init_date = fields.Date(required=True)
    final_date = fields.Date(required=True)    

class Quotes(Resource):
    def __init__(self):
        self.schema = QuotesSchema()
        self.data_access = MetaTraderDataAccess()

    def get(self):
        errors = self.schema.validate(request.args)
        if errors:            
            abort(400, message="Invalid parameters", cause=str(errors))   
        else:
            search_filter = self.schema.load(request.args)
            df = self.data_access.get_rates_from_symbol(search_filter['symbol'], pd.Timestamp(search_filter['init_date']), pd.Timestamp(search_filter['final_date']), mt5.TIMEFRAME_H1)
            # to_json() does not serialize the dataframe index
            df_with_date = df.reset_index()
            df_with_date['Date'] = df.index.astype(str)
            list_dict = []
            for index, row in list(df_with_date.iterrows()):
                list_dict.append(dict(row))
            return list_dict


class Portifolio(Resource):
    def get(self):
        return {'hello': 'world'}

class Order(Resource):
    def get(self):
        return {'hello': 'world'}