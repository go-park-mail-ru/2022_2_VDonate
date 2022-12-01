# syntax=docker/dockerfile:1

FROM golang:latest AS builder
RUN mkdir ../build
RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/
WORKDIR ../build
COPY / .
RUN go mod download
EXPOSE 8083
RUN CGO_ENABLED=0 go build -a -installsuffix cgo cmd/images/images.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/images ./
COPY --from=builder /build/cmd/images/configs/config.yaml ./
CMD ["./images", "-config-path", "config.yaml"]