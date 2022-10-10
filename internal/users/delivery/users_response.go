package httpUsers

import (
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserResponse(u *model.User, c echo.Context) error {
	if u.IsAuthor {
		return c.JSON(http.StatusOK, model.ToAuthor(u))
	}
	return c.JSON(http.StatusOK, model.ToNonAuthor(u))
}
