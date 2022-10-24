# syntax=docker/dockerfile:1

FROM golang:latest

RUN mkdir ../build

WORKDIR ../build

COPY / .

RUN go mod download

EXPOSE 8080

CMD [ "go", "run", "cmd/api/main.go" ]