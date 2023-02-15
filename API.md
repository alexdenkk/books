# API usage

<br />

# Authorization
  Get token by login function in users service and put it in request header
  ```
  Authorization: "Bearer SoMeToKeNsYmBoLs="
  ```

<br />

# Users service
 Users service includes operations with users and authentication

### Usage
 ### POST `/user/register/`
  Register user by login and password, if user registered, returns 201.
  #### Example request data
  ```json
  {
      "login": "alexdenkk",
      "password": "12345678"
  }
  ```

 ### POST `/user/login/`
   Authorization by login and password.
 #### Example request data
 ```json
 {
     "login": "alexdenkk",
     "password": "12345678"
 }
 ```
 #### Example response data
 ```json
 {
     "token": "SoMeToKeNsYmBoLs="
 }
 ```

 ### GET `/user/[id]/`
  Returns user record with 200 code or http error code.

 ### POST `/user/create/`
  Creates user (only for admins). If user created, returns 201.
  #### Example request data
  ```json
  {
      "login": "alexdenkk",
      "password": "12345678",
      "role": 2
  }
  ```

 ### PUT `/user/[id]/`
  Updates user record (user can only update himself, admins can also update users).
  If user correctly updated, returns 200.
  #### Example request data
  ```json
  {
      "login": "alexdenkk"
  }
  ```

 ### DELETE `/user/[id]/`
  Deletes user record (only for admins). If user deleted, returns 200.

<br />

# Books service
 Books service includes operations with books (only for admin users)

### Usage
 ### GET `/book/[id]/`
  Returns book record with 200 code or http error code.

 ### GET `/book/all/`
  Returns all book records.
 
 ### POST `/book/create/`
  Creates book (only for admins). If book created, returns 201.
  #### Example request data
  ```json
  {
      "name": "Meditations",
      "author": "Marcus Aurelius",
      "year": 170,
      "genre_id": 228
  }
  ```
 ### PUT `/book/[id]/`
  Updates book record. If book correctly updated, returns 200.
  #### Example request data
  ```json
  {
      "year": 175
  }
  ```
 ### DELETE `/book/[id]/`
  Deletes book record. If book deleted, returns 200.
