package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{
			http.MethodDelete,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
		},
		AllowOrigins: []string{
			"https://vdonate.ml",
			"http://localhost:8080",
			"http://localhost:4200",
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Content-Type",
			"Content-length",
		},
	})
}
