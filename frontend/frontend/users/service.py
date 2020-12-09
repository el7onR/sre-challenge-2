import requests, json

def get_all_users(app):
    content = {}
    try:
        backend_url = app.config['BACKEND_URL']+"/users"
        result = requests.get(url=backend_url)
        content['status_code'] = result.status_code
        if len(result.content) == 0:
            content["content"] = None
            return content
        content['content'] = json.loads(result.content)
        return content
    except:
        content['status_code'] = 500
        content['content'] = None
        return content       

def get_user(app, id:int):
    content = {}
    backend_url = app.config['BACKEND_URL']+"/users/"+id
    try:
        result = requests.get(url=backend_url)
        content['status_code'] = result.status_code
        if len(result.content) == 0:
            content["content"] = None
            return content
        content['content'] = json.loads(result.content)
        return content
    except:
        content['status_code'] = 500
        content['content'] = None
        return content

def delete_user(app, id:str):
    content = {}
    backend_url = app.config['BACKEND_URL']+"/users/"+id
    try:
        result = requests.delete(url=backend_url)
        content["status_code"] = result.status_code
        return content
    except:
        content['status_code'] = 500
        return content

def create_user(app, payload:dict):
    content = {}
    backend_url = app.config['BACKEND_URL']+"/users"
    try:
        result = requests.post(url=backend_url, data=json.dumps(payload))
        content['status_code'] = result.status_code
        if len(result.content) == 0:
            content["content"] = None
            return content
        content['content'] = json.loads(result.content)
        return content
    except:
        content['status_code'] = 500
        content['content'] = None
        return content

def update_user(app, payload:str):
    backend_url = app.config['BACKEND_URL']+"/users/"+payload["id"]
    content = {}
    try:
        result = requests.put(url=backend_url, data=json.dumps(payload))
        content['status_code'] = result.status_code
        if len(result.content) == 0:
            content["content"] = None
            return content
        content['content'] = json.loads(result.content)
        return content
    except:
        content['status_code'] = 500
        content['content'] = None
        return content
