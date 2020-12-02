# coding=utf-8
__author__ = 'Marcelo Ortiz'

import os

import configparser

__config_file = None
__temp_data_dir = None
__raw_data_dir = None
__backup_data_dir = None
__bovespa_yahoo_data_dir = None
__database_url = None
__bovespa_url = None
__currency_url = None
__project_home = None


def get_project_home():
    global __project_home
    if __project_home == None:
        config = get_config_file()
        bovespa_configs = config['APP']
        __project_home = bovespa_configs['home']
    return __project_home


def get_temp_data_dir(default='bovespa'):
    global __temp_data_dir
    if __temp_data_dir == None:
        __temp_data_dir = '{}/data/{}/temp/'.format(get_project_home(), default)
    return __temp_data_dir


def get_raw_data_dir(default='bovespa'):
    global __raw_data_dir
    if __raw_data_dir == None:
        __raw_data_dir = '{}/data/{}/raw/'.format(get_project_home(), default)
    return __raw_data_dir


def get_backup_data_dir(default='bovespa'):
    global __backup_data_dir
    if __backup_data_dir == None:
        __backup_data_dir = '{}/data/{}/backup/'.format(get_project_home(), default)
    return __backup_data_dir

def get_bovespa_yahoo_dir():
    global __bovespa_yahoo_data_dir
    if __bovespa_yahoo_data_dir == None:
        __bovespa_yahoo_data_dir = '{0}/data/bovespa/y/'.format(get_project_home())
    return __bovespa_yahoo_data_dir


def get_config_file():
    config_file_name = 'config-local.ini'
    if os.environ.get('CONFIG_FILE') != None:
        config_file_name = os.environ.get('CONFIG_FILE')
    global __config_file
    if __config_file == None:
        __config_file = configparser.ConfigParser()
        __config_file.read('{}/{}'.format(os.path.dirname(__file__), config_file_name))
    return __config_file


def get_database_url():
    global __database_url
    if __database_url == None:
        config = get_config_file()
        if os.environ.get('ENV_TEST') != None:
            dbconfigs = config['DB_TEST']
        else:
            dbconfigs = config['DB']
        host = dbconfigs['host']
        port = dbconfigs['port']
        user = dbconfigs['user']
        password = dbconfigs['password']
        database = dbconfigs['database']
        __database_url = 'postgresql://{}:{}@{}:{}/{}'.format(user, password, host, port, database)
    return __database_url


def get_bovespa_url():
    global __bovespa_url
    if __bovespa_url == None:
        config = get_config_file()
        bovespa_configs = config['BOVESPA']
        __bovespa_url = bovespa_configs['url_base']
    return __bovespa_url


def get_currency_url():
    global __currency_url
    if __currency_url == None:
        config = get_config_file()
        currency_configs = config['CURRENCY']
        __currency_url = currency_configs['url_base']
    return __currency_url

