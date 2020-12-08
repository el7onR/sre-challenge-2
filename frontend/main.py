from frontend import create_app
import os

app = create_app(config=os.environ.get('BACKEND_ENV', 'config.DevConfig'))

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000)
