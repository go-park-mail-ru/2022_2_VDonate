# syntax=docker/dockerfile:1

FROM golang:latest AS builder
RUN mkdir ../build
RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/
WORKDIR ../build
COPY / .
RUN go mod download
EXPOSE 8087
RUN CGO_ENABLED=0 go build -a -installsuffix cgo cmd/users/users.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/users ./
COPY --from=builder /build/cmd/users/configs/config.yaml ./
CMD ["./users", "-config-path", "config.yaml"]