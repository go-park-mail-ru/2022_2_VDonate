package httpSubscriptions

import (
	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	subscriptionsDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions"
	subscriptionsErrors "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/errors"
	users "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	subscriptionsUsecase subscriptionsDomain.UseCase
	userUsecase          users.UseCase
}

func New(subscriptionsUsecase subscriptionsDomain.UseCase, userUsecase users.UseCase) *Handler {
	return &Handler{
		subscriptionsUsecase: subscriptionsUsecase,
		userUsecase:          userUsecase,
	}
}

func (h *Handler) GetSubscriptions(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return subscriptionsErrors.Wrap(c, subscriptionsErrors.ErrNoSession, err)
	}
	author, err := h.userUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return subscriptionsErrors.Wrap(c, subscriptionsErrors.ErrNoSession, err)
	}
	if err != nil {
		return err
	}
	s, err := h.subscriptionsUsecase.GetSubscriptionsByAuthorID(author.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, s)
}

func (h *Handler) GetSubscription(c echo.Context) error {
	subID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	s, err := h.subscriptionsUsecase.GetSubscriptionsByID(subID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, s)
}

func (h *Handler) CreateSubscription(c echo.Context) error {
	var s models.AuthorSubscription
	if err := c.Bind(&s); err != nil {
		return err
	}
	newS, err := h.subscriptionsUsecase.AddSubscription(s)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newS)
}

func (h *Handler) UpdateSubscription(c echo.Context) error {
	var s models.AuthorSubscription
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	if err := h.subscriptionsUsecase.DeleteSubscription(id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct{}{})
}
