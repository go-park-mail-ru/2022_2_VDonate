# syntax=docker/dockerfile:1

FROM golang:latest AS builder
RUN mkdir ../build
RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/
WORKDIR ../build
COPY / .
RUN go mod download
EXPOSE 8079
EXPOSE 8080
ARG TOKEN_SECRET
RUN if [  -z ${TOKEN_SECRET} ];then \
  >&2 echo  "\n******Warning!!!!******\n"; \
  >&2 echo  "TOKEN_SECRET seems empty!" ;\
  fi
RUN CGO_ENABLED=0 TOKEN_SECRET=${TOKEN_SECRET} go build -a -installsuffix cgo cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/main ./
COPY --from=builder /build/cmd/api/configs/config_deploy.yaml ./
CMD ["./main", "-config-path", "config_deploy.yaml"]
