# coding=utf-8
__author__ = 'Marcelo Ortiz'

import logging
from logging.handlers import TimedRotatingFileHandler

import configs
from utils.Singleton import Singleton

logger = None

@Singleton
class LogUtils:
    def __init__(self):
        file_log = '%s/logs/bolsa.log' % configs.get_project_home()

        logger = logging.getLogger("Bolsa")
        logger.setLevel(logging.INFO)
        fileHandler = TimedRotatingFileHandler(file_log,
                                           when="h",
                                           interval=1,
                                           backupCount=5)
        fileHandler.setFormatter(logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s'))
        logger.addHandler(fileHandler)

        consoleHandler = logging.StreamHandler()
        consoleHandler.setLevel(logging.INFO)
        consoleHandler.setFormatter(logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s'))
        logger.addHandler(consoleHandler)

        self.logger = logger

    def info(self, msg, *args, **kwargs):
        self.logger.info(msg, *args, **kwargs)

    def error(self, msg, *args, **kwargs):
        self.logger.error(msg, *args, **kwargs)

    def warning(self, msg, *args, **kwargs):
        self.logger.warning(msg, *args, **kwargs)

    def debug(self, msg, *args, **kwargs):
        self.logger.debug(msg, *args, **kwargs)


if __name__ == '__main__':
    LogUtils.getInstance().info("Log")