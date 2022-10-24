package main

import (
	"flag"
	"log"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/labstack/echo/v4"
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
	a := app.New(echo.New(), cfg)
	switch cfg.Deploy.Mode {
	case true:
		if err := a.StartTLS(); err != nil {
			a.Echo.Logger.Error(err)
		}
	default:
		if err := a.Start(); err != nil {
			a.Echo.Logger.Error(err)
		}
	}
}
