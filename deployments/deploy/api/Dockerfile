# syntax=docker/dockerfile:1

FROM golang:latest AS builder
RUN mkdir ../build
RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/
WORKDIR ../build
COPY / .
RUN go mod download
EXPOSE 8080
RUN CGO_ENABLED=0 go build -a -installsuffix cgo cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /build/main ./
COPY --from=builder /build/configs/config_deploy.yaml ./
CMD ["./main", "-config-path", "config_deploy.yaml"]
