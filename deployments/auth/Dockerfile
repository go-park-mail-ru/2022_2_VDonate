# syntax=docker/dockerfile:1

FROM golang:latest AS builder
RUN mkdir ../build
RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/
WORKDIR ../build
COPY / .
RUN go mod download
EXPOSE 8081
EXPOSE 9081
RUN CGO_ENABLED=0 go build -a -installsuffix cgo cmd/auth/auth.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/auth ./
COPY --from=builder /build/cmd/auth/configs/config.yaml ./
CMD ["./auth", "-config-path", "config.yaml"]
