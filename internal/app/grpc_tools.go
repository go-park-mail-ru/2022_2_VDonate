package app

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func CreateGRPCServer(host, port string) (*grpc.Server, net.Listener) {
	listen, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("%s:%s: %s", host, port, err)
	}

	return grpc.NewServer(), listen
}
