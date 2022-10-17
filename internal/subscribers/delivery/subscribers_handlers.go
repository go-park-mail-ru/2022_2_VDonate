package httpsubscribers

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	subscribersUsecase domain.SubscribersUseCase
	userUsecase        domain.UsersUseCase
}

func NewHandler(s domain.SubscribersUseCase, u domain.UsersUseCase) *Handler {
	return &Handler{
		subscribersUsecase: s,
		userUsecase:        u,
	}
}

func (h Handler) GetSubscribers(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("author_id"), 10, 64)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}
	s, err := h.subscribersUsecase.GetSubscribers(id)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, s)
}

func (h Handler) CreateSubscriber(c echo.Context) error {
	var s models.Subscription
	if err := c.Bind(&s); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	err := h.subscribersUsecase.Subscribe(s)
	if err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	return c.JSON(http.StatusOK, s)
}

func (h Handler) DeleteSubscriber(c echo.Context) error {
	var s models.Subscription
	if err := c.Bind(&s); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}
	err := h.subscribersUsecase.Unsubscribe(s.SubscriberID, s.AuthorID)
	if err != nil {
		return utils.WrapEchoError(domain.ErrDelete, err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}
