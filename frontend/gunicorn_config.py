import os

from prometheus_flask_exporter.multiprocess import GunicornPrometheusMetrics

loglevel = 'debug'
accesslog = "-"
errorlog = "-"
workers = 2
timeout = 30
bind = ':8000'
threads = 8



def when_ready(server):
    GunicornPrometheusMetrics.start_http_server_when_ready(8001)


def child_exit(server, worker):
    GunicornPrometheusMetrics.mark_process_dead_on_child_exit(worker.pid)
    