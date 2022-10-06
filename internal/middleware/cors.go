package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func NewCORS(debugMode bool) *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodDelete,
			http.MethodGet,
			http.MethodPost,
		},
		AllowedOrigins: []string{
			"http://localhost:63342",
			"http://localhost:63343",
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"Content-Type",
			"Content-length",
		},
		Debug: debugMode,
	})
}
