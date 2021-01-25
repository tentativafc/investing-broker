# coding=utf-8
__author__ = " Marcelo Ortiz"

from flask_restful import fields


class QuoteDto(object):

    FIELDS = {
        'date': fields.DateTime(dt_format='iso8601'),
        'min_price': fields.Float,
        'max_price': fields.Float,
        'open_price': fields.Float,
        'close_price': fields.Float,
        'volume': fields.Float
    }

    def __init__(self):
        self.date = None
        self.min_price = 0.0
        self.max_price = 0.0
        self.open_price = 0.0
        self.close_price = 0.0
        self.volume = 0.0


class AssetPortfolioDto(object):
    FIELDS = {
        'symbol': fields.String
    }

    def __init__(self):
        self.symbol = None


class PortfolioDto(object):
    FIELDS = {
        'user_id': fields.String,
        'assets': fields.Nested(AssetPortfolioDto.FIELDS),
        'capm': fields.Float,
        'beta': fields.Float,
        'created_at': fields.DateTime(dt_format='iso8601')
    }

    def __init__(self):
        self.user_id = None
        self.assets = []
        self.capm = 0.0
        self.beta = 0.0
        self.created_at = None
