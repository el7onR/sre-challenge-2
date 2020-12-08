import requests, json

def get_all_users(app):
    backend_url = app.config['BACKEND_URL']+"/users"
    result = requests.get(url=backend_url)
    if result.status_code == 200:
        return json.loads("{'status': result.status_code, 'content': result.content}")
    return json.loads({'status':result.status_code, 'content': result.reason})
        
def get_user(app, id:int):
    backend_url = app.config['BACKEND_URL']+"/users/"+id
    return requests.get(url=backend_url)

def delete_user(app, id:int):
    backend_url = app.config['BACKEND_URL']+"/users/"+id
    return requests.delete(url=backend_url)

def create_user(app, payload:dict):
    backend_url = app.config['BACKEND_URL']+"/users"
    result = requests.post(url=backend_url, data=json.dumps(payload))
    if result.status_code == 201:
        return json.loads({"status":result.status_code, "content": result.content})
    return json.loads({"status":result.status_code, "content": result.reason})

def update_user(app, id:int, payload:str):
    backend_url = app.config['BACKEND_URL']+"/users"+id
    return requests.put(url=backend_url, data=payload)

