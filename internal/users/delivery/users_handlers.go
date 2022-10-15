package httpUsers

import (
	authDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	usersDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/users"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/errors"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Handler struct {
	userUseCase    usersDomain.UseCase
	sessionUseCase authDomain.UseCase
}

func NewHandler(userUseCase usersDomain.UseCase, sessionUseCase authDomain.UseCase) *Handler {
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

	return UserResponse(c, user)
}

func (h *Handler) PutUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrConvertID, err)
	}
	var updateUser models.User
	if err = c.Bind(&updateUser); err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrBadRequest, err)
	}

	updateUser.ID = id
	user, err := h.userUseCase.Update(updateUser)
	if err != nil {
		return usersErrors.Wrap(c, usersErrors.ErrUpdate, err)
	}
	return UserResponse(c, user)
}
