import requests
import json

def login(login, password) -> dict:
    resp = requests.post(
        "http://127.0.0.1:8080/user/login/",
        json = {
            "login": "alexdenkk",
            "password": "12345678"
        },
    )
    return resp.json()
