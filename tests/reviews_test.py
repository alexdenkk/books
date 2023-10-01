import auth_test
import requests


token = {
    "Authorization": "Bearer " + auth_test.login("alexdenkk", "123123123")["token"]
}


def create(headers: dict, review: dict):
    resp = requests.post(
        "http://127.0.0.1:8003/review/create/",
        headers=headers,
        json=review
    )
    print(resp.status_code)
    print(resp.content)


"""
create(
    headers=token,
    review={
        "book_id": 1,
        "text": "Fuck this book. I fucking hate it...",
        "rate": 5
    }
)
"""


def update(headers: dict, review: dict, id: int):
    resp = requests.put(
        f"http://127.0.0.1:8003/review/{id}/",
        headers=headers,
        json=review
    )
    print(resp.status_code)


"""
update(
    headers=token, 
    review={
        "text": "Fix this book. I fucking hate it..."
    },
    id=1
)
"""


def get(id: int):
    resp = requests.get(
        f"http://127.0.0.1:8003/review/{id}/",
    )
    print(resp.status_code)
    print(resp.content)


"""
get(id=1)
"""


def get_for(id: int):
    resp = requests.get(
        f"http://127.0.0.1:8003/review/for/{id}/",
    )
    print(resp.status_code)
    print(resp.content)


"""
get_for(id=1)
"""


def delete_for(headers: dict, id: int):
    resp = requests.delete(
        f"http://127.0.0.1:8003/review/for/{id}/",
        headers=headers
    )
    print(resp.status_code)



delete_for(
    headers=token,
    id=1
)
