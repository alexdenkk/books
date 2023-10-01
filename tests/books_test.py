import auth_test
import requests


token = {
    "Authorization": "Bearer " + auth_test.login("alexdenkk", "123123123")["token"]
}


def create(headers: dict, book: dict):
    resp = requests.post(
        "http://127.0.0.1:8002/book/create/",
        headers=headers,
        json=book
    )
    print(resp.status_code)


"""
create(
    headers=token,
    book={
        "name": "meditations",
        "author": "markus aurelius",
        "year": 70,
        "genre_id": 1
    }
)
"""


def update(headers: dict, book: dict, id: int):
    resp = requests.put(
        f"http://127.0.0.1:8002/book/{id}/",
        headers=headers,
        json=book
    )
    print(resp.status_code)


"""
update(
    headers=token,
    book={
        "name": "fixed meditations"
    },
    id=1
)
"""


def get(id: int):
    resp = requests.get(
        f"http://127.0.0.1:8002/book/{id}/",
    )
    print(resp.status_code)
    print(resp.content)


"""
get(id=1)
"""


def get_all():
    resp = requests.get(
        f"http://127.0.0.1:8002/book/all/",
    )
    print(resp.status_code)
    print(resp.content)


"""
get_all()
"""


def delete(headers: dict, id: int):
    resp = requests.delete(
        f"http://127.0.0.1:8002/book/{id}/",
        headers=headers
    )
    print(resp.status_code)


"""
delete(
    headers=token,
    id=1
)
"""