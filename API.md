# API usage

## Authorization
  Get token by login function in users service and put it in request header
  ```
  Authorization: "Bearer SoMeToKeNsYmBoLs="
  ```

## Users service
 Users service includes operations with users and authentication

### Usage
 ### POST /user/register/
  Register user by login and password, if user registered, returns 201.
  #### Example request data
  ```json
  {
      "login": "alexdenkk",
      "password": "12345678"
  }
  ```

 ### POST /user/login/
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

 ### GET /user/id/
  Returns user record with 200 code or http error code.

 ### POST /user/create/
  Creates user (only for admins). If user created, returns 201.

 ### PUT /user/id/
  Update user record (user can only update himself, admins can also update users).
  If user correctly updated, returns 200.
 
 ### DELETE /user/id/
  Deletes user (only for admins). If user deleted, returns 200.
