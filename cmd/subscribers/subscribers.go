package main

import (
	"flag"
	defaultLogger "log"
	"net/http"

	grpcSubscribers "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/grpc"

	subscribersRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers/repository"

	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/protobuf"

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
		defaultLogger.Fatalf("subscribers: failed to open config: %s", err)
	}

	/*---------------------------logger---------------------------*/
	log := logger.GetInstance()
	log.SetLevel(logger.ToLevel(cfg.Logger.Level))
	log.Info("subscribers: server started")

	/*----------------------------repo----------------------------*/
	r, err := subscribersRepository.NewPostgres(cfg.DB.URL)
	if err != nil {
		log.Fatalf("subscribers: failed to open db: %s", err)
	}
	defer r.DB.Close()

	/*----------------------------grpc----------------------------*/
	metricsHTTP := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog: log,
	}), Addr: "0.0.0.0" + ":" + cfg.Services.Subscribers.MetricsPort}

	lis, metricsServer := app.CreateGRPCServer(cfg.Server.Host, cfg.Server.Port, grpcMetrics)
	defer lis.Close()

	protobuf.RegisterSubscribersServer(metricsServer, grpcSubscribers.New(r))

	/*---------------------------metric---------------------------*/
	grpcMetrics.InitializeMetrics(metricsServer)

	go func() {
		if err = metricsHTTP.ListenAndServe(); err != nil {
			log.Warnf("subscribers: prometheus: HTTP server stopped: %s", err)
		}
	}()

	/*---------------------------server---------------------------*/
	if err = metricsServer.Serve(lis); err != nil {
		log.Warnf("subscribers: %s", "service image stopped")
	}
}
