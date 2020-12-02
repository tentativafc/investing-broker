# coding=utf-8
__author__ = 'Marcelo Ortiz'

from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
import configs

__engine = None
__session = None


def get_session():
    global __engine
    global __session
    if __session == None:
        __session = sessionmaker(bind=get_engine())
    return __session()


def get_engine():
    global __engine
    if __engine == None:
        __engine = create_engine(configs.get_database_url(), pool_size=20, max_overflow=0)
    return __engine

def transactional(func):
    def wrapper(*args, **kwargs):
        session = get_session()
        try:
            result = func(*args, **dict(kwargs, session=session))
            session.commit()
            return result
        except Exception as exc:
            session.rollback()
            raise exc
        finally:
            session.close()


def non_transactional(func):
    def wrapper(*args, **kwargs):
        session = get_session()
        try:
            result = func(*args, **dict(kwargs, session=session))
            return result
        except Exception as exc:
            raise exc
        finally:
            session.close()

    return wrapper
