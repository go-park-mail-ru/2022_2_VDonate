# syntax=docker/dockerfile:1

FROM golang:latest AS builder
RUN mkdir ../build
RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/
WORKDIR ../build
COPY / .
RUN go mod download
EXPOSE 8082
RUN CGO_ENABLED=0 go build -a -installsuffix cgo cmd/donates/donates.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/donates ./
COPY --from=builder /build/cmd/donates/configs/config.yaml ./
CMD ["./donates", "-config-path", "config.yaml"]
