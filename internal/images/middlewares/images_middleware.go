package imagesMiddleware

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func BucketManager(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := strings.TrimLeft(c.Path(), "api/v1")
		switch {
		case strings.Contains(path, "users"):
			c.Set("bucket", "avatar")
		default:
			c.Set("bucket", "image")
		}
		return next(c)
	}
}
