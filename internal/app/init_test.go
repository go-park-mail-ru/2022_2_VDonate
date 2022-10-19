package app

import (
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

func TestServer_Start(t *testing.T) {
	e := echo.New()

	cfg := config.New()
	err := cfg.Open("../../configs/config_local.yaml")
	if err != nil {
		t.Error(t, err)
	}

	server := New(e, cfg)
	server.init()

	assert.NotEqual(t, server.authHandler, nil)
	assert.NotEqual(t, server.authMiddleware, nil)
	assert.NotEqual(t, server.postsHandler, nil)
	assert.NotEqual(t, server.userHandler, nil)
	assert.NotEqual(t, server.AuthService, nil)
	assert.NotEqual(t, server.PostsService, nil)
	assert.NotEqual(t, server.UserService, nil)
	assert.NotEqual(t, server.Config, nil)
}
