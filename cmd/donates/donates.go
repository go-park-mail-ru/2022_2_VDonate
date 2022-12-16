package main

import (
	"flag"
	"net/http"

	donatesRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/donates/repository"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	defaultLogger "log"

	grpcDonate "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/donates/grpc"

	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/donates/protobuf"

	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/prometheus/client_golang/prometheus"
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
	metricsHTTP := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog: log,
	}), Addr: "0.0.0.0" + ":" + cfg.Services.Donates.MetricsPort}

	listener, grpcServer := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port, grpcMetrics)
	defer listener.Close()

	protobuf.RegisterDonatesServer(grpcServer, grpcDonate.New(r))

	/*---------------------------metric---------------------------*/
	grpcMetrics.InitializeMetrics(grpcServer)

	go func() {
		if err = metricsHTTP.ListenAndServe(); err != nil {
			log.Warnf("donates: prometheus: HTTP server stopped: %s", err)
		}
	}()

	/*---------------------------server---------------------------*/
	if err = grpcServer.Serve(listener); err != nil {
		log.Warnf("donates: %s", "service image stopped")
	}
}
