package httpSubscribers

import (
	"net/http"
	"strconv"

	"github.com/ztrue/tracerr"

	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/http"

	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	subscribersUsecase   domain.SubscribersUseCase
	subscriptionsUsecase domain.SubscriptionsUseCase
	userUsecase          domain.UsersUseCase
}

func NewHandler(s domain.SubscribersUseCase, u domain.UsersUseCase, as domain.SubscriptionsUseCase) *Handler {
	return &Handler{
		subscribersUsecase:   s,
		userUsecase:          u,
		subscriptionsUsecase: as,
	}
}

// GetSubscribers godoc
// @Summary     Returns subscribers by author ID
// @Description Request to server for subscriptions of requested `Author`
// @ID          get_subscribers
// @Tags        subscribers
// @Produce     json
// @Param       author_id path     integer         true "Author ID"
// @Success     200       {object} []models.Author "Get list of author subscribers"
// @Failure     400       {object} echo.HTTPError  "Bad request"
// @Failure     404       {object} echo.HTTPError  "Not found"
// @Security    ApiKeyAuth
// @Router      /subscribers/{author_id} [get]
func (h Handler) GetSubscribers(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("author_id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	s, err := h.subscribersUsecase.GetSubscribers(id)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, s)
}

// CreateSubscriber godoc
// @Summary     Subscribe
// @Description Subscribe user to author with paid subscription
// @ID          create_subscriber
// @Tags        subscribers
// @Produce     json
// @Param       Subscription body     models.QiwiPaymentStatus true "Payment response, documentation: https://developer.qiwi.com/ru/p2p-payments/?shell#create"
// @Success     200          {object} models.Subscription      "Successfully subscribed"
// @Failure     400          {object} echo.HTTPError           "Bad request"
// @Failure     500          {object} echo.HTTPError           "Not created"
// @Security    ApiKeyAuth
// @Router      /subscribers [post]
func (h Handler) CreateSubscriber(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	u, err := h.userUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	var s models.Subscription
	if err = c.Bind(&s); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	sub, err := h.subscriptionsUsecase.GetAuthorSubscriptionByID(s.AuthorSubscriptionID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	response, err := h.subscribersUsecase.Subscribe(s, u.ID, sub)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrCreate, err)
	}

	return c.JSON(http.StatusOK, response)
}

// DeleteSubscriber godoc
// @Summary     Unsubscribe
// @Description Unsubscribe user from author
// @ID          delete_subscriber
// @Tags        subscribers
// @Accept      json
// @Produce     json
// @Param       Subscription body     models.SubscriptionMpfd true "Subscription info with required UserID, AuthorID and Subscription ID"
// @Success     200          {object} models.EmptyStruct      "Subscriber was successfully unsubscribed"
// @Failure     400          {object} echo.HTTPError          "Bad request"
// @Failure     500          {object} echo.HTTPError          "Not deleted"
// @Security    ApiKeyAuth
// @Router      /subscribers [delete]
func (h Handler) DeleteSubscriber(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	u, err := h.userUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	var s models.Subscription
	if err = c.Bind(&s); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	if err = h.subscribersUsecase.Unsubscribe(u.ID, s.AuthorID); err != nil {
		return errorHandling.WrapEcho(domain.ErrDelete, err)
	}

	return c.JSON(http.StatusOK, models.EmptyStruct{})
}

// Withdraw godoc
// @Summary     Withdraw
// @Description ?????????? ?????????????? ???? QIWI ?????????????? ?????? ???????????????????? ??????????
// @ID          withdraw
// @Tags        withdraw
// @Accept      json
// @Produce     json
// @Param       Input body     models.Withdraw      true "???????????????????? ?? ???????????? ??????????????, ?????????????????? ???????? ???? ?????????? Phone ?????? Card, ???????????? ??????????"
// @Success     200   {object} models.WithdrawInfo  "?????????????? ??????????????, ?????????????? ???????????????????? ?? ????????????????"
// @Failure     400   {object} echo.HTTPError       "Bad request (?? ???????????????? ???????????????? ???????????? ????????????)"
// @Failure     500   {object} models.WithdrawError "???????????????????? ???????????? ???????? ???????????? ??????????????, ???????? ???? ??????????????, ?? ?????????????????????? ???? ?????????? ?????????? ???????????? ???????????? ???????? ????????, ???????? ????"
// @Security    ApiKeyAuth
// @Router      /withdraw [post]
func (h Handler) Withdraw(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	u, err := h.userUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	var w models.Withdraw
	if err = c.Bind(&w); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	if w.UserID != u.ID {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	info, err := h.subscribersUsecase.Withdraw(w.UserID, w.Phone, w.Card)
	if err != nil {
		if eTrace, ok := err.(tracerr.Error); ok {
			if e, ok := eTrace.Unwrap().(models.WithdrawError); ok {
				return c.JSON(http.StatusInternalServerError, e)
			}
		}
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, info)
}
