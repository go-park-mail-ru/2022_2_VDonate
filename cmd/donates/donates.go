package main

import (
	"flag"
	defaultLogger "log"
	"net/http"

	donatesRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/donates/repository"

	grpcDonate "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/donates/grpc"

	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/donates/protobuf"
)

func main() {
	/*----------------------------flag----------------------------*/
	var configPath string
	config.PathFlag(&configPath)
	flag.Parse()

	/*---------------------------config---------------------------*/
	cfg := config.New()
	if err := cfg.Open(configPath); err != nil {
		defaultLogger.Fatalf("donates: failed to open config: %s", err)
	}

	/*---------------------------logger---------------------------*/
	log := logger.GetInstance()
	log.SetLevel(logger.ToLevel(cfg.Logger.Level))
	log.Info("donates: server started")

	/*----------------------------repo----------------------------*/
	r, err := donatesRepository.NewPostgres(cfg.DB.URL)
	if err != nil {
		log.Fatalf("donates: failed to open db: %s", err)
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
	}), Addr: "localhost" + ":" + cfg.Services.Donates.MetricsPort}

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.EnableClientHandlingTimeHistogram()

	protobuf.RegisterDonatesServer(metricsServer, grpcDonate.New(r))
	grpc_prometheus.Register(metricsServer)

	go func() {
		if err = metricsHTTP.ListenAndServe(); err != nil {
			log.Warnf("donates: prometheus: HTTP server stopped: %s", err)
		}
	}()

	/*---------------------------server---------------------------*/
	if err = metricsServer.Serve(lis); err != nil {
		log.Warnf("donates: %s", "service image stopped")
	}
}
