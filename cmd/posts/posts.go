package main

import (
	"flag"
	"net/http"

	postsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/repository"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	defaultLogger "log"

	grpcPosts "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/grpc"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
)

var (
	reg = prometheus.NewRegistry()

	grpcMetrics = grpcPrometheus.NewServerMetrics()
)

func main() {
	reg.MustRegister(grpcMetrics)

	/*----------------------------flag----------------------------*/
	var configPath string
	config.PathFlag(&configPath)
	flag.Parse()

	/*---------------------------config---------------------------*/
	cfg := config.New()
	if err := cfg.Open(configPath); err != nil {
		defaultLogger.Fatalf("posts: failed to open config: %s", err)
	}

	/*---------------------------logger---------------------------*/
	log := logger.GetInstance()
	log.SetLevel(logger.ToLevel(cfg.Logger.Level))
	log.Info("posts: server started")

	/*----------------------------repo----------------------------*/
	r, err := postsRepository.NewPostgres(cfg.DB.URL, cfg.DB.MaxOpenConns)
	if err != nil {
		log.Fatalf("posts: failed to open db: %s", err)
	}
	defer r.DB.Close()

	/*----------------------------grpc----------------------------*/
	metricsHTTP := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog: log,
	}), Addr: "0.0.0.0" + ":" + cfg.Services.Posts.MetricsPort}

	listener, grpcServer := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port, grpcMetrics)
	defer listener.Close()

	protobuf.RegisterPostsServer(grpcServer, grpcPosts.New(r))

	/*---------------------------metric---------------------------*/
	grpcMetrics.InitializeMetrics(grpcServer)

	go func() {
		if err = metricsHTTP.ListenAndServe(); err != nil {
			log.Warnf("posts: prometheus: HTTP server stopped: %s", err)
		}
	}()

	/*---------------------------server---------------------------*/
	if err = grpcServer.Serve(listener); err != nil {
		log.Warnf("posts: %s", "service image stopped")
	}
}
