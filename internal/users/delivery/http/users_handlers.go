package httpUsers

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/session/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	userAPI     userAPI.UseCase
	sessionRepo sessionRepository.API
}

func NewHandler(userAPI userAPI.UseCase, sessionRepo sessionRepository.API) *Handler {
	return &Handler{userAPI: userAPI, sessionRepo: sessionRepo}
}

func (h *Handler) GetUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrConvertID, err)
	}
	user, err := h.userAPI.GetByID(uint(id))
	if err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrUserNotFound, err)
	}

	return c.JSON(http.StatusOK, user)
}
