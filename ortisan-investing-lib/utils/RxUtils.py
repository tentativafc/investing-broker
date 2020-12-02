import multiprocessing

from rx import Observable
from rx.concurrency import ThreadPoolScheduler


class RxUtils(object):
    @staticmethod
    def get_multithread_observable(iterable):
        optimal_thread_count = multiprocessing.cpu_count() + 1
        pool_scheduler = ThreadPoolScheduler(optimal_thread_count)
        return Observable.from_iterable(iterable).subscribe_on(pool_scheduler)
