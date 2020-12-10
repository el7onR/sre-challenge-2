FROM python:3.9.0-alpine@sha256:df77433749466a68bb599009753c9e5a8efaa3dd9c16450d442bb32f4c1fad4e
LABEL maintainer=eltonribeiro@outlook.com
LABEL app=frontend
EXPOSE 8000
WORKDIR /app
COPY . .
RUN pip install --no-cache-dir -r requirements.txt
RUN apk --no-cache add curl
ENV prometheus_multiproc_dir /tmp
ENV METRICS_PORT 9200 
ENTRYPOINT ["gunicorn", "-c", "gunicorn_config.py", "wsgi:app", "--log-level", "info"]
