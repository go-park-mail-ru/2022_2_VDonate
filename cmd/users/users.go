package main

import (
	"flag"
	"log"

	grpcUsers "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/grpc"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"

	userRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
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
	r, err := userRepository.NewPostgres(cfg.DB.URL)
	if err != nil {
		log.Fatalf("subscriptions: %s", err)
	}

	/*----------------------------grpc----------------------------*/
	s, l := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port)
	defer l.Close()
	protobuf.RegisterUsersServer(s, grpcUsers.New(r))

	/*---------------------------server---------------------------*/
	if err = s.Serve(l); err != nil {
		log.Printf("subscriptions: %s", "service image stopped")
	}
}
