package main

import (
	"flag"
	defaultLogger "log"
	"net/http"
	"time"

	notificationsWS "github.com/go-park-mail-ru/2022_2_VDonate/internal/notifications/delivery/ws"
	"github.com/gorilla/websocket"

	notificationsUsecase "github.com/go-park-mail-ru/2022_2_VDonate/internal/notifications/usecase"

	notificationsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/notifications/repository"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
)

func main() {
	/*----------------------------flag----------------------------*/
	var configPath string
	config.PathFlag(&configPath)
	flag.Parse()

	/*---------------------------config---------------------------*/
	cfg := config.New()
	if err := cfg.Open(configPath); err != nil {
		defaultLogger.Fatalf("notifications: failed to open config: %s", err)
	}

	/*---------------------------logger---------------------------*/
	log := logger.GetInstance()
	log.SetLevel(logger.ToLevel(cfg.Logger.Level))
	log.Info("notifications: server started")

	/*----------------------------repo----------------------------*/
	r, err := notificationsRepository.New(cfg.DB.URL)
	if err != nil {
		log.Fatalf("notifications: failed to open db: %s", err)
	}
	defer r.DB.Close()

	u := notificationsUsecase.New(r)

	h := notificationsWS.NewHandler(&websocket.Upgrader{
		HandshakeTimeout: time.Minute,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}, u)

	http.HandleFunc("/ws", h.Handler)

	if cfg.Deploy.Mode {
		if err = http.ListenAndServeTLS(cfg.Server.Host+":"+cfg.Server.Port, cfg.Server.CertPath,
			cfg.Server.KeyPath, nil); err != nil {
			log.Warnf("notifications: stop to listen and serve: %s", err)
		}
	} else {
		if err = http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, nil); err != nil {
			log.Warnf("notifications: stop to listen and serve: %s", err)
		}
	}
}
