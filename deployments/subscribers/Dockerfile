# syntax=docker/dockerfile:1

FROM golang:latest AS builder
RUN mkdir ../build
RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/
WORKDIR ../build
COPY / .
RUN go mod download
EXPOSE 8085
EXPOSE 9085
RUN CGO_ENABLED=0 go build -a -installsuffix cgo cmd/subscribers/subscribers.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/subscribers ./
COPY --from=builder /build/cmd/subscribers/configs/config.yaml ./
CMD ["./subscribers", "-config-path", "config.yaml"]