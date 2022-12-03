package main

import (
	"flag"
	defaultLogger "log"
	"net/http"

	grpcSubscriptions "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscriptions/grpc"

	subscriptionsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/repository"

	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscriptions/protobuf"

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
		defaultLogger.Fatalf("subscriptions: failed to open config: %s", err)
	}

	/*---------------------------logger---------------------------*/
	log := logger.GetInstance()
	log.SetLevel(logger.ToLevel(cfg.Logger.Level))
	log.Info("subscriptions: server started")

	/*----------------------------repo----------------------------*/
	r, err := subscriptionsRepository.NewPostgres(cfg.DB.URL)
	if err != nil {
		log.Fatalf("subscriptions: failed to open db: %s", err)
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
	}), Addr: "localhost" + ":" + cfg.Services.Subscriptions.MetricsPort}

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.EnableClientHandlingTimeHistogram()

	protobuf.RegisterSubscriptionsServer(metricsServer, grpcSubscriptions.New(r))
	grpc_prometheus.Register(metricsServer)

	go func() {
		if err = metricsHTTP.ListenAndServe(); err != nil {
			log.Warnf("subscriptions: prometheus: HTTP server stopped: %s", err)
		}
	}()

	/*---------------------------server---------------------------*/
	if err = metricsServer.Serve(lis); err != nil {
		log.Warnf("subscriptions: %s", "service image stopped")
	}
}
