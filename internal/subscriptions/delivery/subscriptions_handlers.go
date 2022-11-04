package httpSubscriptions

import (
	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}
	user, err := h.userUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}
	s, err := h.subscriptionsUsecase.GetSubscriptionsByUserID(user.ID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}
	if len(s) == 0 {
		return c.JSON(http.StatusOK, struct{}{})
	}

	return c.JSON(http.StatusOK, s)
}

func (h Handler) GetAuthorSubscriptions(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}
	author, err := h.userUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}
	s, err := h.subscriptionsUsecase.GetAuthorSubscriptionsByAuthorID(author.ID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}
	if len(s) == 0 {
		return c.JSON(http.StatusOK, struct{}{})
	}

	return c.JSON(http.StatusOK, s)
}

func (h Handler) GetAuthorSubscription(c echo.Context) error {
	subID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}
	s, err := h.subscriptionsUsecase.GetAuthorSubscriptionByID(subID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, s)
}

func (h Handler) CreateAuthorSubscription(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}
	author, err := h.userUsecase.GetBySessionID(cookie.Value)
	var s models.AuthorSubscription
	if err := c.Bind(&s); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}
	s.AuthorID = author.ID
	newS, err := h.subscriptionsUsecase.AddAuthorSubscription(s)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrCreate, err)
	}

	return c.JSON(http.StatusOK, newS)
}

func (h Handler) UpdateAuthorSubscription(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}
	var s models.AuthorSubscription
	if err := c.Bind(&s); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	s.ID = id
	subscription, err := h.subscriptionsUsecase.UpdateAuthorSubscription(s)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrUpdate, err)
	}

	return c.JSON(http.StatusOK, subscription)
}

func (h Handler) DeleteAuthorSubscription(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}
	if err := h.subscriptionsUsecase.DeleteAuthorSubscription(id); err != nil {
		return errorHandling.WrapEcho(domain.ErrDelete, err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}
