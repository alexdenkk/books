import auth_test
import requests


token = {
    "Authorization": "Bearer " + auth_test.login("alexdenkk", "123123123")["token"]
}


def create(headers: dict, genre: dict):
    resp = requests.post(
        "http://127.0.0.1:8004/genre/create/",
        headers=headers,
        json=genre
    )
    print(resp.status_code)


"""
create(
    headers=token,
    genre={
        "name": "horror"
    }
)
"""


def update(headers: dict, genre: dict, id: int):
    resp = requests.put(
        f"http://127.0.0.1:8004/genre/{id}/",
        headers=headers,
        json=genre
    )
    print(resp.status_code)


"""
update(
    headers=token,
    genre={
        "name": "fixed horror"
    },
    id=1
)
"""


def get(id: int):
    resp = requests.get(
        f"http://127.0.0.1:8004/genre/{id}/",
    )
    print(resp.status_code)
    print(resp.content)


"""
get(id=1)
"""


def get_all():
    resp = requests.get(
        f"http://127.0.0.1:8004/genre/all/",
    )
    print(resp.status_code)
    print(resp.content)


"""
get_all()
"""


def delete(headers: dict, id: int):
    resp = requests.delete(
        f"http://127.0.0.1:8004/genre/{id}/",
        headers=headers
    )
    print(resp.status_code)


"""
delete(
    headers=token,
    id=1
)
"""