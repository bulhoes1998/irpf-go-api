FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN go build cmd/api/main.go

EXPOSE 8080