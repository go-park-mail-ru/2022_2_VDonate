package main

import (
	"flag"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	defaultLogger "log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	sessionsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/repository"
	grpcAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/grpc"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
)

var (
	reg = prometheus.NewRegistry()

	grpcMetrics = grpc_prometheus.NewServerMetrics()
)

func init() {
	reg.MustRegister(grpcMetrics)
}

func main() {
	/*----------------------------flag----------------------------*/
	var configPath string
	config.PathFlag(&configPath)
	flag.Parse()

	/*---------------------------config---------------------------*/
	cfg := config.New()
	if err := cfg.Open(configPath); err != nil {
		defaultLogger.Fatalf("auth: failed to open config: %s", err)
	}

	/*---------------------------logger---------------------------*/
	log := logger.GetInstance()
	log.SetLevel(logger.ToLevel(cfg.Logger.Level))
	log.Info("auth: server started")

	/*----------------------------repo----------------------------*/
	r, err := sessionsRepository.NewPostgres(cfg.DB.URL)
	if err != nil {
		log.Fatalf("auth: failed to open db: %s", err)
	}

	/*----------------------------grpc----------------------------*/
	metricsHTTP := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog: log,
	}), Addr: "0.0.0.0" + ":" + cfg.Services.Auth.MetricsPort}

	lis, metricsServer := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port, grpcMetrics)
	defer lis.Close()

	protobuf.RegisterAuthServer(metricsServer, grpcAuth.New(r))

	/*---------------------------metric---------------------------*/
	grpcMetrics.InitializeMetrics(metricsServer)

	go func() {
		if err = metricsHTTP.ListenAndServe(); err != nil {
			log.Warnf("auth: prometheus: HTTP server stopped: %s", err)
		}
	}()

	/*---------------------------server---------------------------*/
	if err = metricsServer.Serve(lis); err != nil {
		log.Warnf("auth: %s", "service image stopped")
	}
}
