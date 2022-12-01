package main

import (
	"flag"
	"log"

	grpcPosts "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/grpc"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/protobuf"

	postsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/repository"

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
	r, err := postsRepository.NewPostgres(cfg.DB.URL)
	if err != nil {
		log.Fatalf("posts: %s", err)
	}

	/*----------------------------grpc----------------------------*/
	s, l := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port)
	defer l.Close()
	protobuf.RegisterPostsServer(s, grpcPosts.New(r))

	/*---------------------------server---------------------------*/
	if err = s.Serve(l); err != nil {
		log.Printf("posts: %s", "service image stopped")
	}
}
