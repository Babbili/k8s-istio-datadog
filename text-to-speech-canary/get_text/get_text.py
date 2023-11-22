import requests

def gettext():
    r = requests.get("http://goapi.apps.svc.cluster.local/books")
    if r.status_code == 200:
        data = r.json()
        return str(data)
    else:
        return f"Error: {r.status_code}"
