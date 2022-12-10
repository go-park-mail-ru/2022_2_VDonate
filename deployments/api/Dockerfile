# syntax=docker/dockerfile:1

FROM golang:latest AS builder
RUN mkdir ../build
RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/
WORKDIR ../build
COPY / .
RUN go mod download
EXPOSE 8079
EXPOSE 8080
RUN CGO_ENABLED=0 go build -a -installsuffix cgo cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/main ./
COPY --from=builder /build/cmd/api/configs/config_local.yaml ./
CMD ["./main", "-config-path", "config_local.yaml"]
