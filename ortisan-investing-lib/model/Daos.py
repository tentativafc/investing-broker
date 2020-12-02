# coding=utf-8
__author__ = 'Marcelo Ortiz'

import collections

from model.DbUtil import get_engine
from model.DbUtil import get_session
from model.Models import AssetQuotation, CoinQuotation


class GenericDao(object):
    def insert(self, element_or_iterable):
        session = get_session()
        iterable = self.__get_iterable(element_or_iterable)
        try:
            for element in iterable:
                session.add(element)
            session.commit()
        except Exception as exc:
            session.rollback()
            raise exc
        finally:
            session.flush()
            session.close()

    def update(self, element_or_iterable, load=True):
        session = get_session()
        iterable = self.__get_iterable(element_or_iterable)
        try:
            if load:
                for element in iterable:
                    element_on_db = session.query(type(element)).populate_existing().get(element.id)
                    element_on_db = self.__clone_attrs_obj(element, element_on_db)
            session.commit()
        except Exception as exc:
            session.rollback()
            raise exc
        finally:
            session.flush()
            session.close()

    def delete(self, element_or_iterable):
        session = get_session()
        iterable = self.__get_iterable(element_or_iterable)
        try:
            for element in iterable:
                session.delete(element)
            session.commit()
        except Exception as exc:
            session.rollback()
            raise exc
        finally:
            session.flush()
            session.close()

    def find_by_id(self, type, id):
        session = get_session()
        try:
            return session.query(type).populate_existing().get(id)
        finally:
            session.close()

    def execute_native_query(self, query):
        session = get_session()
        try:
            result = session.execute(query)
        finally:
            session.close()
        return result

    def __get_iterable(self, element_or_iterable):
        if not isinstance(element_or_iterable, collections.Iterable):
            element_or_iterable = [element_or_iterable]
        return element_or_iterable

    def __clone_attrs_obj(self, obj_from, obj_to):
        type_of_from = type(obj_from)
        type_of_to = type(obj_to)
        assert type_of_from == type_of_to, 'Tipos devem ser iguais'
        columns = type_of_from.__table__.columns
        for column in columns:
            value = getattr(obj_from, column.name)
            setattr(obj_to, column.name, value)
        return obj_to


class AssetQuotationDao(GenericDao):

    def get_symbols(self, market_type=10, bdi_code=2):

        with get_engine().connect() as con:
            rows = con.execute(
                'SELECT DISTINCT symbol FROM asset_quotation WHERE market_type={} AND bdi_code={} ORDER BY symbol'.format(
                    market_type, bdi_code))
            return [row[0] for row in rows]

    def get_by_asset_and_date(self, symbol, init_date, end_date, market_type=10, bdi_code=2):
        session = get_session()
        try:
            quotations = session.query(AssetQuotation) \
                .filter_by(symbol=symbol, market_type=market_type, bdi_code=bdi_code) \
                .filter(AssetQuotation.date >= init_date) \
                .filter(AssetQuotation.date <= end_date) \
                .order_by(AssetQuotation.date.asc())

            return [quotation for quotation in quotations]
        finally:
            session.close()


class CoinQuotationDao(GenericDao):
    def get_by_symbol_and_date(self, symbol, init_date, end_date):
        session = get_session()
        try:
            quotations = session.query(CoinQuotation) \
                .filter_by(symbol=symbol) \
                .filter(CoinQuotation.date >= init_date) \
                .filter(CoinQuotation.date <= end_date) \
                .order_by(CoinQuotation.date.asc())

            return [quotation for quotation in quotations]
        finally:
            session.close()


if __name__ == '__main__':
    asset_quotation_dao = AssetQuotationDao()
    print(asset_quotation_dao.get_symbols())
