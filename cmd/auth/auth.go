package main

import (
	"flag"
	"log"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"
	sessionsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/grpc"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
)

func main() {
	/*----------------------------flag----------------------------*/
	var configPath string
	config.PathFlag(&configPath)
	flag.Parse()

	/*---------------------------config---------------------------*/
	cfg := config.New()
	if err := cfg.Open(configPath); err != nil {
		log.Fatalf("failed to open config: %s", err)
	}

	/*----------------------------repo----------------------------*/
	r, err := sessionsRepository.NewPostgres(cfg.DB.URL)
	if err != nil {
		log.Fatalf("failed to open db: %s", err)
	}

	/*----------------------------grpc----------------------------*/
	server, lis := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port)
	defer lis.Close()
	protobuf.RegisterAuthServer(server, grpcAuth.New(r))

	/*---------------------------server---------------------------*/
	if err = server.Serve(lis); err != nil {
		log.Printf("images: %s", "service image stopped")
	}
}
