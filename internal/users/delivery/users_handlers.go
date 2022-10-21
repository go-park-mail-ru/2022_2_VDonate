package httpUsers

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Handler struct {
	sessionUseCase       domain.AuthUseCase
	subscriptionsUseCase domain.SubscriptionsUseCase
	userUseCase          domain.UsersUseCase
}

func NewHandler(
	userUseCase domain.UsersUseCase,
	sessionUseCase domain.AuthUseCase,
	subscriptionsUseCase domain.SubscriptionsUseCase,
) *Handler {
	return &Handler{
		userUseCase:          userUseCase,
		sessionUseCase:       sessionUseCase,
		subscriptionsUseCase: subscriptionsUseCase,
	}
}

func (h *Handler) GetUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}
	user, err := h.userUseCase.GetByID(id)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	if user.IsAuthor {
		user.AuthorSubscriptions, err = h.subscriptionsUseCase.GetAuthorSubscriptionsByAuthorID(user.ID)
		if err != nil {
			return errorHandling.WrapEcho(domain.ErrInternal, err)
		}
	}
	user.UserSubscriptions, err = h.subscriptionsUseCase.GetSubscriptionsByUserID(user.ID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	return UserResponse(c, user)
}

func (h *Handler) PutUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}
	var updateUser models.User
	if err = c.Bind(&updateUser); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	updateUser.ID = id
	user, err := h.userUseCase.Update(updateUser)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrUpdate, err)
	}
	return UserResponse(c, user)
}
