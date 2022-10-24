package httpUsers

import (
	"net/http"

	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
)

func UserResponse(c echo.Context, u *model.User) error {
	if u.IsAuthor {
		return c.JSON(http.StatusOK, model.ToAuthor(u))
	}

	return c.JSON(http.StatusOK, model.ToNonAuthor(u))
}
