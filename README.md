# Books
 Books reviews API written with Go/Jwt/Mux/Docker/Postgres

<br />

# Services

## Services list
- users
- books
- reviews
- genres

## Service struct
- gateway (communication layer)
- service (business logic)
- repository (database connection layer)

<br />

# Usage

## Run with docker
```cmd
cd deployments/docker-compose
docker-compose build --no-cache api
docker-compose up -d db api
```

## View the usage of app
```
go build -o server cmd/main.go
./server --help
```

<br />

# API usage
 API usage descripted in API.md file
