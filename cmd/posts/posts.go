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
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
)

var (
	reg = prometheus.NewRegistry()

	grpcMetrics = grpc_prometheus.NewServerMetrics()
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
	r, err := postsRepository.NewPostgres(cfg.DB.URL)
	if err != nil {
		log.Fatalf("posts: failed to open db: %s", err)
	}

	/*----------------------------grpc----------------------------*/
	metricsHTTP := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog: log,
	}), Addr: "0.0.0.0" + ":" + cfg.Services.Posts.MetricsPort}

	lis, metricsServer := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port, grpcMetrics)
	defer lis.Close()

	protobuf.RegisterPostsServer(metricsServer, grpcPosts.New(r))

	/*---------------------------metric---------------------------*/
	grpcMetrics.InitializeMetrics(metricsServer)

	go func() {
		if err = metricsHTTP.ListenAndServe(); err != nil {
			log.Warnf("posts: prometheus: HTTP server stopped: %s", err)
		}
	}()

	/*---------------------------server---------------------------*/
	if err = metricsServer.Serve(lis); err != nil {
		log.Warnf("posts: %s", "service image stopped")
	}
}
