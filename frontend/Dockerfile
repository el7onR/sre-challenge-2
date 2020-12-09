FROM python:3.9.0-alpine@sha256:df77433749466a68bb599009753c9e5a8efaa3dd9c16450d442bb32f4c1fad4e
LABEL maintainer=eltonribeiro@outlook.com
LABEL app=frontend
EXPOSE 8000
WORKDIR /app
COPY . .
RUN pip install --no-cache-dir -r requirements.txt
RUN apk --no-cache add curl
ENTRYPOINT ["gunicorn", "-b", ":8000", "wsgi:app", "-w", "4", "-t", "30", "--threads", "4"]
