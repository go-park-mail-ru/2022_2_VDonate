package httpUsers

import (
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserResponse(c echo.Context, u *model.User) error {
	if u.IsAuthor {
		return c.JSON(http.StatusOK, model.ToAuthor(u))
	}
	return c.JSON(http.StatusOK, model.ToNonAuthor(u))
}

func UsersResponse(c echo.Context, u []*model.User) error {
	var result []interface{}
	for _, user := range u {
		if user.IsAuthor {
			result = append(result, model.ToAuthor(user))
		}
		result = append(result, model.ToNonAuthor(user))
	}
	return c.JSON(http.StatusOK, result)
}
