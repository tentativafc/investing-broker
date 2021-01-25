# coding=utf-8
__author__ = 'Marcelo Ortiz'

import json
from bson import ObjectId
import datetime

class JSONEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, ObjectId):
            return str(o)
        if isinstance(o, (datetime.datetime, datetime.date)):
            print(o.isoformat())
            return o.isoformat()
        return json.JSONEncoder.default(self, o)