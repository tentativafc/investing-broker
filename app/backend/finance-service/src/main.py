from flask import Flask, make_response
from flask_restful import Api

from flask import json
import routes

from flask.json import JSONEncoder
from bson import json_util
from mongoengine.base import BaseDocument
from mongoengine.queryset.base import BaseQuerySet
import datetime

app = Flask(__name__)
api = Api(app)

if __name__ == '__main__':
    api.add_resource(routes.Portfolio, '/api/portfolio')
    api.add_resource(routes.Quotes, '/api/quotes')
    api.add_resource(routes.Order, '/api/orders')

    app.run(debug=True)