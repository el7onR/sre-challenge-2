from os import environ, path
from dotenv import load_dotenv

basedir = path.abspath(path.dirname(__file__))
load_dotenv(path.join(basedir, '.env'))


class DevConfig():
    FLASK_ENV = 'development'
    DEBUG = True
    TESTING = True
    LOGGING_LEVEL = environ.get('FRONTEND_LOGGING_LEVEL', 'DEBUG')
    BACKEND_URL = "http://{}:{}".format(
        environ.get('FRONTEND_BACKEND_URL', 'localhost'),
        environ.get('FRONTEND_BACKEND_PORT', '8080'))
