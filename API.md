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
      "year": 70,
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

<br />

# Reviews service
 Reviews service includes operations with reviews

### Usage
 ### GET `/review/[id]/`
  Returns review record with 200 code or http error code.

 ### GET `/review/for/[id]/`
  Returns all reviews for book by id.
 
 ### POST `/review/create/`
  Creates review. If review created, returns 201.
  #### Example request data
  ```json
  {
      "user_id": 7,
      "book_id": 8,
      "text": "This is the text of review",
      "rate": 9
  }
  ```
 ### PUT `/review/[id]/`
  Updates review record. If review correctly updated, returns 200.

  #### Example request data
  ```json
  {
      "text": "This is the updated text of review"
  }
  ```

 ### DELETE `/review/[id]/`
  Deletes review record. If review deleted, returns 200.

 ### DELETE `/review/for/[id]/`
  Deletes all reviews for book by id. If reviews deleted, returns 200.

<br />

# Genres service
 Genres service includes operations with genres

### Usage
 ### GET `/genre/[id]/`
  Returns genre record with 200 code or http error code.

 ### GET `/genre/all/`
  Returns all genres records.
 
 ### POST `/genre/create/`
  Creates genre. If genre created, returns 201.
  #### Example request data
  ```json
  {
      "name": "horror",
  }
  ```
 ### PUT `/genre/[id]/`
  Updates genre record. If genre correctly updated, returns 200.

  #### Example request data
  ```json
  {
      "name": "fixed horror"
  }
  ```

 ### DELETE `/genre/[id]/`
  Deletes genre record. If genre deleted, returns 200.
