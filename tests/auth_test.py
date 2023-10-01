import requests


def register(login: str, password: str):
    resp = requests.post("http://127.0.0.1:8001/user/register/", json={
        "login": login,
        "password": password
    })

    print(resp.status_code)
    print(resp.content)


def login(login: str, password: str) -> dict:
    resp = requests.post("http://127.0.0.1:8001/user/login/", json={
        "login": login,
        "password": password
    })

    print(resp.status_code)
    return resp.json()




