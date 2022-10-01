package cors

import (
	"github.com/rs/cors"
	"net/http"
)

func New(debugMode bool) *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodDelete,
			http.MethodGet,
			http.MethodPost,
		},
		AllowedOrigins: []string{
			"http://localhost:63342",
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"Content-Type",
			"Content-length",
		},
		Debug: debugMode,
	})
}
