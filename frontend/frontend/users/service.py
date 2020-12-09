import requests, json

def get_all_users(app):
    content = {}
    backend_url = app.config['BACKEND_URL']+"/users"
    result = requests.get(url=backend_url)
    content["status_code"] = (result.status_code)
    if result.status_code != 404:
        content["content"] = (json.loads(result.content))
    return content
        
def get_user(app, id:int):
    backend_url = app.config['BACKEND_URL']+"/users/"+id
    return requests.get(url=backend_url)

def delete_user(app, id:str):
    backend_url = app.config['BACKEND_URL']+"/users/"+id
    return requests.delete(url=backend_url)

def create_user(app, payload:dict):
    content = {}
    backend_url = app.config['BACKEND_URL']+"/users"
    result = requests.post(url=backend_url, data=json.dumps(payload))
    content["status_code"] = (result.status_code)
    content["content"] = (json.loads(result.content))
    return content

def update_user(app, id:int, payload:str):
    backend_url = app.config['BACKEND_URL']+"/users"+id
    return requests.put(url=backend_url, data=payload)

