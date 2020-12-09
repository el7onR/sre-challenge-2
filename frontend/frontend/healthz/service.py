import requests

def backend_status(app):
    backend_url = app.config['BACKEND_URL']+"/healthz"
    content = {}
    try:
        result = requests.get(backend_url)
        content["status_code"] = result.status_code
        if len(result.content) == 0:
            content["content"] = None
            return content
        content["content"] = result.content
        return content
    except:
        content["status_code"] = 500
        content["content"] = None
        return content