package httpSubscriptions

import (
	"net/http"
	"strconv"

	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	subscriptionsUsecase domain.SubscriptionsUseCase
	userUsecase          domain.UsersUseCase
}

func NewHandler(s domain.SubscriptionsUseCase, u domain.UsersUseCase) *Handler {
	return &Handler{
		subscriptionsUsecase: s,
		userUsecase:          u,
	}
}

func (h Handler) GetSubscriptions(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	author, err := h.userUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	s, err := h.subscriptionsUsecase.GetSubscriptionsByAuthorID(author.ID)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, s)
}

func (h Handler) GetSubscription(c echo.Context) error {
	subID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	s, err := h.subscriptionsUsecase.GetSubscriptionsByID(subID)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, s)
}

func (h Handler) CreateSubscription(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	author, err := h.userUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	var s models.AuthorSubscription

	if err = c.Bind(&s); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	s.AuthorID = author.ID
	newS, err := h.subscriptionsUsecase.AddSubscription(s)
	if err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	return c.JSON(http.StatusOK, newS)
}

func (h Handler) UpdateSubscription(c echo.Context) error {
	var s models.AuthorSubscription
	if err := c.Bind(&s); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	subscription, err := h.subscriptionsUsecase.UpdateSubscription(s)
	if err != nil {
		return utils.WrapEchoError(domain.ErrUpdate, err)
	}

	return c.JSON(http.StatusOK, subscription)
}

func (h Handler) DeleteSubscription(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	if err = h.subscriptionsUsecase.DeleteSubscription(id); err != nil {
		return utils.WrapEchoError(domain.ErrDelete, err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}
