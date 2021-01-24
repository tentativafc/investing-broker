# coding=utf-8
__author__ = 'Marcelo Ortiz'


from flask import request
from flask_restful import Resource, Api, reqparse, abort
from marshmallow import Schema, fields
import re

from datetime import datetime
import pandas as pd
from business import Business


class QuotesGetSchema(Schema):
    symbol = fields.Str(required=True)
    init_date = fields.Date(required=True)
    final_date = fields.Date(required=True)    


class Quotes(Resource):
    def __init__(self):
        self.getSchema = QuotesGetSchema()
        self.business = Business()

    def get(self):
        errors = self.getSchema.validate(request.args)
        if errors:            
            abort(400, message="Invalid parameters", cause=str(errors))   
        else:
            search_filter = self.getSchema.load(request.args)
            return self.business.getQuotes(search_filter['symbol'], pd.Timestamp(search_filter['init_date']), pd.Timestamp(search_filter['final_date']))


class PortfolioGetSchema(Schema):
    amount_assets = fields.Integer(required=True)

class Portfolio(Resource):
    def __init__(self):
        self.getSchema = PortfolioGetSchema()
        self.business = Business()

    def get(self):
        errors = self.getSchema.validate(request.args)
        if errors:            
            abort(400, message="Invalid parameters", cause=str(errors))   
        else:
            search_filter = self.getSchema.load(request.args)
            return self.business.createPortfolio(search_filter['amount_assets'])

class Order(Resource):
    def get(self):
        return {'hello': 'world'}