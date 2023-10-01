FROM golang:1.21

LABEL maintaner="alexdenkk@yandex.ru"

WORKDIR server/

COPY . .

RUN go mod download

COPY go.mod go.sum ./

RUN go build -o users cmd/main.go
