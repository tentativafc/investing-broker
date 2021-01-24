# coding=utf-8
__author__ = " Marcelo Ortiz" 

import os

from mongoengine import register_connection
from mongoengine.document import Document, EmbeddedDocument
from mongoengine.fields import DateTimeField
from mongoengine.fields import FloatField
from mongoengine.fields import IntField
from mongoengine.fields import ReferenceField
from mongoengine.fields import StringField
from mongoengine.fields import ListField
from mongoengine.fields import EmbeddedDocumentField

from datetime import datetime

register_connection(alias="default", name="finance-mongo", host="localhost", port=27021, username="mongo", password="123456", db="finance", authentication_source="admin")

class AssetPortfolio(EmbeddedDocument):
    symbol = StringField(required=True)
    date = DateTimeField()
    min_price = FloatField()
    max_price = FloatField()
    open_price = FloatField()
    close_price = FloatField()
    volume = FloatField()

class Portfolio(Document):
    user_id = StringField(required=True)
    assets = ListField(EmbeddedDocumentField(AssetPortfolio))
    capm = FloatField()
    beta = FloatField()
    created_at = DateTimeField(default=datetime.utcnow)
    # Index on user_id
    meta = {
        " indexes" : [
            " user_id" 
        ]
    }