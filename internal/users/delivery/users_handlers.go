package httpUsers

import (
	"strconv"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	sessionUseCase domain.AuthUseCase
	userUseCase    domain.UsersUseCase
}

func NewHandler(userUseCase domain.UsersUseCase, sessionUseCase domain.AuthUseCase) *Handler {
	return &Handler{
		userUseCase:    userUseCase,
		sessionUseCase: sessionUseCase,
	}
}

func (h *Handler) GetUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	user, err := h.userUseCase.GetByID(id)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	return UserResponse(c, user)
}

func (h *Handler) PutUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	var updateUser models.User

	if err = c.Bind(&updateUser); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	updateUser.ID = id
	user, err := h.userUseCase.Update(updateUser)
	if err != nil {
		return utils.WrapEchoError(domain.ErrUpdate, err)
	}

	return UserResponse(c, user)
}
