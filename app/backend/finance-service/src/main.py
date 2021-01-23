from flask import Flask
from flask_restful import Api

import routes

app = Flask(__name__)
api = Api(app)

api.add_resource(routes.Portifolio, '/api/portifolio')
api.add_resource(routes.Quotes, '/api/quotes')
api.add_resource(routes.Order, '/api/orders')

if __name__ == '__main__':
    app.run(debug=True)