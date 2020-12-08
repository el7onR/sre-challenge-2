from flask import Flask
from .home import home
from .users import users

def create_app(config:str):
    app = Flask(__name__)
    app.config.from_object(config or 'default')

    with app.app_context():
        app.register_blueprint(users.users_bp) 
        app.register_blueprint(home.home_bp) 

    return app