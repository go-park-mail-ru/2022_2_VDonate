package main

import (
	"flag"
	"github.com/go-park-mail-ru/2022_2_VDonate/init/system"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/labstack/echo/v4"
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

	/*----------------------------server--------------------------*/
	s := system.New(echo.New(), cfg)
	if err := s.Start(); err != nil {
		s.Echo.Logger.Error("server errors: %s", err)
	}
}
