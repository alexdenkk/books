# Books
Books reviews microservice API written with Go/Jwt/Mux/Docker/Postgres

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

# Run

## Run with docker
```cmd
cd deployments/docker-compose
docker-compose up -d db
docker-compose up -d users books reviews genres
```

<br />

# Usage

## Read the usage
 - godoc documentation
 - API.md file
