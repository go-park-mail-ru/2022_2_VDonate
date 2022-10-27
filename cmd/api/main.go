package main

import (
	"flag"
	"log"

	_ "github.com/go-park-mail-ru/2022_2_VDonate/docs"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/labstack/echo/v4"
)

// @title   VDonate API
// @version 1.0
// @description
//	  ## API of donation web service by VDonate team
//    ### Some useful links:
//    - ### [VDonate repository](https://github.com/go-park-mail-ru/2022_2_VDonate)
//    - ### [Trello](https://trello.com/b/BZHoJsHP/vdonate)
// @termsOfService http://swagger.io/terms/

// @contact.name  VDonate Support
// @contact.email zeronethunter2001@gmail.com

// @host     vdonate.ml:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
// @description                Authorization via Cookie
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
