package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewCSRF() echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(middleware.CSRFConfig{
		Skipper:      middleware.DefaultBasicAuthConfig.Skipper,
		TokenLength:  32,
		TokenLookup:  "header:" + echo.HeaderXCSRFToken,
		ContextKey:   "csrf",
		CookieName:   "csrf_token",
		CookieMaxAge: 86400,
	})
}
