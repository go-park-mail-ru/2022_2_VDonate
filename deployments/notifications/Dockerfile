# syntax=docker/dockerfile:1

FROM golang:latest AS builder
RUN mkdir ../build
RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/
WORKDIR ../build
COPY / .
RUN go mod download
EXPOSE 1234
RUN CGO_ENABLED=0 go build -a -installsuffix cgo cmd/notifications/notifications.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/notifications ./
COPY --from=builder /build/cmd/notifications/configs/config.yaml ./
CMD ["./notifications", "-config-path", "config.yaml"]