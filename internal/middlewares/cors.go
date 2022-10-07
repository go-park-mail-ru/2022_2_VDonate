package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func NewCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{
			http.MethodDelete,
			http.MethodGet,
			http.MethodPost,
		},
		AllowOrigins: []string{
			"http://localhost:63342",
			"http://localhost:63343",
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Content-Type",
			"Content-length",
		},
	})
}
