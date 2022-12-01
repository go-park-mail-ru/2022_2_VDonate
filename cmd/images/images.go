package main

import (
	"flag"
	"log"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"

	imagesRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/repository"
	grpcImages "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/images/grpc"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/images/protobuf"

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
		log.Fatalf("images: failed to open config: %s", err)
	}

	/*----------------------------repo----------------------------*/
	r, err := imagesRepository.New(
		cfg.S3.Endpoint,
		cfg.S3.AccessKeyID,
		cfg.S3.SecretAccessKey,
		cfg.S3.UseSSL,
		cfg.S3.Buckets.SymbolsToHash,
		cfg.S3.Buckets.Policy,
		cfg.S3.Buckets.Expire,
	)
	if err != nil {
		log.Fatalf("images: %s", err)
	}
	s, l := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port)
	protobuf.RegisterImagesServer(s, grpcImages.New(r))

	if err = s.Serve(l); err != nil {
		log.Printf("images: %s", "service image stopped")
	}
}
