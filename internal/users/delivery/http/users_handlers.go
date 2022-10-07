package httpUsers

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/session/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
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

	return UserResponse(user, c)
}

func (h *Handler) PutUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrConvertID, err)
	}
	var updateUser *models.UserDB
	if err = c.Bind(&updateUser); err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrBadRequest, err)
	}
	user, err := h.userAPI.GetByID(uint(id))
	if err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrUserNotFound, err)
	}

	if err = copier.CopyWithOption(&user, &updateUser, copier.Option{IgnoreEmpty: true}); err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrUpdate, err)
	}

	if updateUser, err = h.userAPI.Update(user); err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrUpdate, err)
	}
	return UserResponse(user, c)
}
