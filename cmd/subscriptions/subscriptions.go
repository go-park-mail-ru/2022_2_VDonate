package main

import (
	"flag"
	"log"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"
	grpcSubscriptions "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscriptions/grpc"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscriptions/protobuf"

	subscriptionsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/repository"

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
	r, err := subscriptionsRepository.NewPostgres(cfg.DB.URL)
	if err != nil {
		log.Fatalf("subscriptions: %s", err)
	}

	/*----------------------------grpc----------------------------*/
	s, l := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port)
	defer l.Close()
	protobuf.RegisterSubscriptionsServer(s, grpcSubscriptions.New(r))

	/*---------------------------server---------------------------*/
	if err = s.Serve(l); err != nil {
		log.Printf("subscriptions: %s", "service image stopped")
	}
}
