package main

import (
	"flag"
	defaultLogger "log"
	"net/http"

	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"

	grpcImages "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/images/grpc"

	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"

	imagesRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/images/protobuf"

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
		defaultLogger.Fatalf("images: failed to open config: %s", err)
	}

	/*---------------------------logger---------------------------*/
	log := logger.GetInstance()
	log.SetLevel(logger.ToLevel(cfg.Logger.Level))
	log.Info("images: server started")

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

	/*----------------------------grpc----------------------------*/
	metricsHTTP := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog: log,
	}), Addr: "0.0.0.0" + ":" + cfg.Services.Images.MetricsPort}

	listener, grpcServer := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port, grpcMetrics)
	defer listener.Close()

	protobuf.RegisterImagesServer(grpcServer, grpcImages.New(r))

	/*---------------------------metric---------------------------*/
	grpcMetrics.InitializeMetrics(grpcServer)

	go func() {
		if err = metricsHTTP.ListenAndServe(); err != nil {
			log.Warnf("images: prometheus: HTTP server stopped: %s", err)
		}
	}()

	/*---------------------------server---------------------------*/
	if err = grpcServer.Serve(listener); err != nil {
		log.Warnf("images: %s", "service image stopped")
	}
}
