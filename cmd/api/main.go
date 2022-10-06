package main

import (
	"flag"
	"github.com/go-park-mail-ru/2022_2_VDonate/init/system"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	storage "github.com/go-park-mail-ru/2022_2_VDonate/internal/storages"
	logger "github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
	"log"
)

func main() {
	/*----------------------------flag----------------------------*/
	var configPath string
	config.PathFlag(&configPath)
	flag.Parse()

	/*---------------------------config---------------------------*/
	cfg := config.New()
	if err := cfg.Open(configPath); err != nil {
		log.Fatalf("failed to open config: %s", err)
	}

	/*---------------------------logrus---------------------------*/
	l, err := logger.NewLogrus(cfg.Server.Host, cfg.Server.Port)
	if err != nil {
		log.Fatalf("failed to init logrus: %s", err)
	}
	l.Logrus.Info("logrus initialized successfully")

	/*---------------------------storage--------------------------*/
	store := storage.New()
	if err = store.Open(cfg.DB.Driver, cfg.DB.URL); err != nil {
		l.Logrus.Fatalf("failed to create storage: %s", err)
	}

	/*----------------------------server--------------------------*/
	s := system.New(store, l, cfg)
	if err = s.Start(); err != nil {
		s.Logger.Logrus.Fatalf("server errors: %s", err)
	}
}
