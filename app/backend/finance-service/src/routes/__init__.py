# coding=utf-8
__author__ = 'Marcelo Ortiz'


from datetime import datetime

import pandas as pd
from business import Business
from dtos import PortfolioDto, QuoteDto
from errors import IllegalArgumentException
from flask import request
from flask_restful import Api, Resource, abort, marshal_with, reqparse
from marshmallow import Schema, fields


class QuotesGetSchema(Schema):
    symbol = fields.Str(required=True)
    init_date = fields.Date(required=True)
    final_date = fields.Date(required=True)


class Quotes(Resource):
    def __init__(self):
        self.getSchema = QuotesGetSchema()
        self.business = Business()

    @marshal_with(QuoteDto.FIELDS)
    def get(self):
        try:
            errors = self.getSchema.validate(request.args)
            if errors:
                abort(400, message="Invalid parameters", cause=str(errors))
            else:
                search_filter = self.getSchema.load(request.args)
                return self.business.getQuotes(search_filter['symbol'], pd.Timestamp(search_filter['init_date']), pd.Timestamp(search_filter['final_date']))
        except Exception as exc:
            abort(500, message="Error to get quotes.", cause=exc)


class PortfolioGetSchema(Schema):
    amount_assets = fields.Integer(required=True)


class Portfolio(Resource):
    def __init__(self):
        self.getSchema = PortfolioGetSchema()
        self.business = Business()

    @marshal_with(PortfolioDto.FIELDS)
    def get(self):
        try:
            errors = self.getSchema.validate(request.args)
            if errors:
                abort(400, message="Invalid parameters", cause=str(errors))
            else:
                search_filter = self.getSchema.load(request.args)
                return self.business.createPortfolio(search_filter['amount_assets'])
        except Exception as exc:
            if isinstance(exc, IllegalArgumentException):
                abort(400, message=str(exc))
            else:
                abort(500, message="Error to create portfolio.", cause=exc)


class Order(Resource):
    def get(self):
        return {'hello': 'world'}
