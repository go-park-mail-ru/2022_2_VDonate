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
	defer r.DB.Close()

	/*----------------------------grpc----------------------------*/
	metricsHTTP := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog: log,
	}), Addr: "0.0.0.0" + ":" + cfg.Services.Subscriptions.MetricsPort}

	lis, metricsServer := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port, grpcMetrics)
	defer lis.Close()

	protobuf.RegisterSubscriptionsServer(metricsServer, grpcSubscriptions.New(r))

	/*---------------------------metric---------------------------*/
	grpcMetrics.InitializeMetrics(metricsServer)

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
