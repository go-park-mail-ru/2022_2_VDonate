# syntax=docker/dockerfile:1

FROM golang:latest AS builder
RUN mkdir ../build
RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/
WORKDIR ../build
COPY / .
RUN go mod download
EXPOSE 8086
EXPOSE 9086
RUN CGO_ENABLED=0 go build -a -installsuffix cgo cmd/subscriptions/subscriptions.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/subscriptions ./
COPY --from=builder /build/cmd/subscriptions/configs/config.yaml ./
CMD ["./subscriptions", "-config-path", "config.yaml"]