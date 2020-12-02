# coding=utf-8
__author__ = 'Marcelo Ortiz'

import glob
import os
from os.path import basename
import pickle

class FileUtils(object):

    @classmethod
    def slice_files(cls, pathname_pattern, max_lines, destination_dir, file_encoding='utf-8'):
        for file in glob.glob(pathname_pattern):
            smallfile = None
            with open(file, encoding=file_encoding) as bigfile:
                for lineno, line in enumerate(bigfile):
                    if lineno % max_lines == 0:
                        if smallfile:
                            smallfile.close()
                        file_split = '{0}/{1}.{2:02d}'.format(destination_dir, basename(file), lineno)
                        smallfile = open(file_split, "w")
                    if line.startswith('01'):
                        smallfile.write(line)
                if smallfile:
                    smallfile.close()

    @classmethod
    def delete_files(cls, pathname_pattern):
        for file in glob.glob(pathname_pattern):
            os.remove(file)

    @classmethod
    def save_to_pickle(cls, data, filename='file_example'):
        with open('%s.pickle' %(filename), 'wb') as handle:
            pickle.dump(data, handle)
    
    @classmethod
    def load_from_pickle(cls, filename='file_example'):
        with open('%s.pickle' %(filename), 'rb') as handle:
            return pickle.load(handle)
