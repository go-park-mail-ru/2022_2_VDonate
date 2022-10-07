package httpPosts

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	postsAPI postsAPI.UseCase
}

func NewHandler(p postsAPI.UseCase) *Handler {
	return &Handler{postsAPI: p}
}

func (h *Handler) GetPosts(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}
	posts, err := h.postsAPI.GetAllByUserID(uint(id))
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, posts)
}
