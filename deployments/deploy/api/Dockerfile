# syntax=docker/dockerfile:1

FROM golang:latest

RUN mkdir ../build

RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/

WORKDIR ../build

COPY / .

RUN go mod download

EXPOSE 8080

CMD [ "go", "run", "cmd/api/main.go", "-config-path", "configs/config_deploy.yaml" ]