package main

import (
	"flag"
	defaultLogger "log"
	"net/http"

	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"

	grpcImages "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/images/grpc"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"

	imagesRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/repository"
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
	lis, metricsServer := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port)
	defer lis.Close()

	/*---------------------------metric---------------------------*/
	grpcMetrics := grpc_prometheus.NewServerMetrics()
	gatherer := prometheus.NewRegistry()
	gatherer.MustRegister(grpcMetrics)

	grpcMetrics.InitializeMetrics(metricsServer)

	metricsHTTP := &http.Server{Handler: promhttp.HandlerFor(gatherer, promhttp.HandlerOpts{
		ErrorLog: log,
	}), Addr: "localhost" + ":" + cfg.Services.Images.MetricsPort}

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.EnableClientHandlingTimeHistogram()

	protobuf.RegisterImagesServer(metricsServer, grpcImages.New(r))
	grpc_prometheus.Register(metricsServer)

	go func() {
		if err = metricsHTTP.ListenAndServe(); err != nil {
			log.Warnf("images: prometheus: HTTP server stopped: %s", err)
		}
	}()

	/*---------------------------server---------------------------*/
	if err = metricsServer.Serve(lis); err != nil {
		log.Warnf("images: %s", "service image stopped")
	}
}
