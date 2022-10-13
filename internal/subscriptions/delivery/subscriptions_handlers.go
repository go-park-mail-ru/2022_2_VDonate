package httpSubscriptions

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/usecase"
	users "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	subscriptionsUsecase subscriptions.UseCase
	userUsecase          users.UseCase
}

func New(subscriptionsUsecase subscriptions.UseCase, userUsecase users.UseCase) *Handler {
	return &Handler{
		subscriptionsUsecase: subscriptionsUsecase,
		userUsecase:          userUsecase,
	}
}

func (h *Handler) GetSubscriptions(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("author_id"), 10, 64)
	if err != nil {
		return err
	}
	s, err := h.subscriptionsUsecase.GetSubscriptions(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, s)
}

func (h *Handler) CreateSubscription(c echo.Context) error {
	var s *models.AuthorSubscription
	if err := c.Bind(&s); err != nil {
		return err
	}
	s, err := h.subscriptionsUsecase.AddSubscription(s)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, s)
}

func (h *Handler) UpdateSubscription(c echo.Context) error {
	var s *models.AuthorSubscription
	if err := c.Bind(&s); err != nil {
		return err
	}

	subscription, err := h.subscriptionsUsecase.UpdateSubscription(s)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, subscription)
}

func (h *Handler) DeleteSubscription(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("subscription_id"), 10, 64)
	if err != nil {
		return err
	}
	if err := h.subscriptionsUsecase.DeleteSubscription(id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct{}{})
}
