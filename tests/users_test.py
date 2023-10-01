import auth_test
import requests


token = {
    "Authorization": "Bearer " + auth_test.login("alexdenkk", "123123123")["token"]
}


def create(headers: dict, user: dict):
    resp = requests.post(
        "http://127.0.0.1:8001/user/create/",
        headers=headers,
        json=user
    )
    print(resp.status_code)


"""
create(
    headers=token,
    user={
        "login": "alexdenkk2",
        "password": "123123123",
        "role": 1
    }
)
"""


def update(headers: dict, user: dict, id: int):
    resp = requests.put(
        f"http://127.0.0.1:8001/user/{id}/",
        headers=headers,
        json=user
    )
    print(resp.status_code)


"""
update(
    headers=token,
    user={
        "login": "alexdenkk2fixed"
    },
    id=8
)
"""


def get(id: int):
    resp = requests.get(
        f"http://127.0.0.1:8001/user/{id}/",
    )
    print(resp.status_code)
    print(resp.content)


"""
get(id=1)
"""


def delete(headers: dict, id: int):
    resp = requests.delete(
        f"http://127.0.0.1:8001/user/{id}/",
        headers=headers
    )
    print(resp.status_code)


"""
delete(
    headers=token,
    id=2
)
"""