# books
Books reviews API written with Go/Jwt/Mux

# services
- users
- books
- reviews

# run with docker
```cmd
cd deployments/docker-compose
docker-compose build --no-cache api
docker-compose up -d db
docker-compose up api
```