from frontend import create_app
import os
from prometheus_flask_exporter.multiprocess import GunicornPrometheusMetrics


app = create_app(config=os.environ.get('BACKEND_ENV', 'config.DevConfig'))
metrics = GunicornPrometheusMetrics(app)
metrics.info('frontend', 'frontend application', version='1.0.0')


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000)
