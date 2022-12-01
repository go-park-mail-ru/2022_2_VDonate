package main

import (
	"flag"
	"log"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/protobuf"

	grpcSubscribers "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/grpc"

	subscribersRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers/repository"

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
	r, err := subscribersRepository.NewPostgres(cfg.DB.URL)
	if err != nil {
		log.Fatalf("subscribers: %s", err)
	}

	/*----------------------------grpc----------------------------*/
	s, l := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port)
	defer l.Close()
	protobuf.RegisterSubscribersServer(s, grpcSubscribers.New(r))

	/*---------------------------server---------------------------*/
	if err = s.Serve(l); err != nil {
		log.Printf("subscribers: %s", "service image stopped")
	}
}
