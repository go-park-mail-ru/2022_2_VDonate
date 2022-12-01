package httpUsers

import (
	"net/http"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
)

func UserResponse(c echo.Context, u model.User) error {
	if u.IsAuthor {
		return c.JSON(http.StatusOK, model.ToAuthor(u))
	}

	return c.JSON(http.StatusOK, model.ToNonAuthor(u))
}

func AuthorsResponse(c echo.Context, u []model.User) error {
	authors := make([]models.Author, 0)
	for _, user := range u {
		authors = append(authors, model.ToAuthor(user))
	}

	return c.JSON(http.StatusOK, authors)
}
