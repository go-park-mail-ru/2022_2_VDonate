package app

import (
	"log"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func CreateGRPCServer(host, port string) (net.Listener, *grpc.Server) {
	listen, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("%s:%s: %s", host, port, err)
	}

	return listen, grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
}
