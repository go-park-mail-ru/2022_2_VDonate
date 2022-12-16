package app

import (
	"log"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func CreateGRPCServer(host, port string, grpcMetrics *grpc_prometheus.ServerMetrics) (net.Listener, *grpc.Server) {
	listen, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("%s:%s: %s", host, port, err)
	}

	metricsServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
	)

	grpcMetrics.EnableHandlingTimeHistogram()

	return listen, metricsServer
}
