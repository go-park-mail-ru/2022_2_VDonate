package httpUsers

import (
	auth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Handler struct {
	userUseCase    users.UseCase
	sessionUseCase auth.UseCase
}

func NewHandler(userUseCase users.UseCase, sessionUseCase auth.UseCase) *Handler {
	return &Handler{userUseCase: userUseCase, sessionUseCase: sessionUseCase}
}

func (h *Handler) GetUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrConvertID, err)
	}
	user, err := h.userUseCase.GetByID(id)
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
	var updateUser *models.User
	if err = c.Bind(&updateUser); err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrBadRequest, err)
	}

	if updateUser, err = h.userUseCase.Update(id, updateUser); err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrUpdate, err)
	}
	return UserResponse(updateUser, c)
}
