# syntax=docker/dockerfile:1

FROM golang:latest AS builder
RUN mkdir ../build
RUN mkdir -p /etc/letsencrypt/live/vdonate.ml/
WORKDIR ../build
COPY / .
RUN go mod download
EXPOSE 8084
EXPOSE 9084
RUN CGO_ENABLED=0 go build -a -installsuffix cgo cmd/posts/posts.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/posts ./
COPY --from=builder /build/cmd/posts/configs/config.yaml ./
CMD ["./posts", "-config-path", "config.yaml"]
