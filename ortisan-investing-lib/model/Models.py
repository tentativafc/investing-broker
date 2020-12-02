# coding=utf-8
__author__ = 'Marcelo Ortiz'

from sqlalchemy import Column, Integer, String, DateTime, Float
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()

class AssetQuotation(Base):
    __tablename__ = 'asset_quotation'
    id = Column(String, primary_key=True)
    symbol = Column(String)
    market_type = Column(Integer)
    bdi_code = Column(Integer)
    date = Column(DateTime)
    days_term_market = Column(Integer)
    min_price = Column(Float)
    max_price = Column(Float)
    open_price = Column(Float)
    close_price = Column(Float)
    volume = Column(Float)

class CoinQuotation(Base):
    __tablename__ = 'coin_quotation'
    id = Column(String, primary_key=True)
    symbol = Column(String, nullable=False)
    date = Column(DateTime)
    min_price = Column(Float)
    max_price = Column(Float)
    open_price = Column(Float)
    close_price = Column(Float)
    volume = Column(Float)
    quantity = Column(Float)
    amount = Column(Float)
    avg_price = Column(Float)


class CurrencyQuotation(Base):
    __tablename__ = 'currency_quotation'
    id = Column(String, primary_key=True)
    symbol = Column(String)
    date = Column(DateTime)
    buy_price = Column(Float)
    sell_price = Column(Float)
