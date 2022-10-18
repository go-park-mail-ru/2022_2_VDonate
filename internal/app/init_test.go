package app

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer_Start(t *testing.T) {
	e := echo.New()

	cfg := config.New()
	err := cfg.Open("../../configs/config_local.yaml")
	assert.NoError(t, err)

	server := New(e, cfg)
	server.init()
}
