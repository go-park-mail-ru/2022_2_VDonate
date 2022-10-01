package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/server"
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
	if _, err := toml.DecodeFile(configPath, cfg); err != nil {
		log.Fatalf("failed to get config: %s", err)
	}

	/*---------------------------logrus---------------------------*/
	l, err := logger.NewLogrus(cfg.Port, cfg.DbName)
	if err != nil {
		log.Fatalf("failed to init logrus: %s", err)
	}
	l.Logrus.Info("logrus initialized successfully")

	/*---------------------------storage--------------------------*/
	store := storage.New(cfg.Storage)
	if err = store.Open(cfg.DbName); err != nil {
		l.Logrus.Fatalf("failed to create storage: %s", err)
	}

	/*----------------------------server--------------------------*/
	s := server.New(store, l, cfg)
	if err = s.Start(); err != nil {
		s.Logger.Logrus.Fatalf("server error: %s", err)
	}
}
