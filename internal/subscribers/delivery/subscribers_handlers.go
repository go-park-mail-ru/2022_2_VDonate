package httpsubscribers

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"
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
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	s, err := h.subscribersUsecase.GetSubscribers(id)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, s)
}

// CreateSubscriber godoc
// @Summary     Subscribe
// @Description Subscribe user to author with paid subscription
// @ID          create_subscriber
// @Tags        subscribers
// @Accept      json
// @Produce     json
// @Param       Subscription body     models.Subscription true "Subscription info with required UserID, AuthorID and Subscription ID"
// @Success     200          {object} models.Subscription "Successfully subscribed"
// @Failure     400          {object} echo.HTTPError      "Bad request"
// @Failure     500          {object} echo.HTTPError      "Not created"
// @Security    ApiKeyAuth
// @Router      /subscribers [post]
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

// DeleteSubscriber godoc
// @Summary     Unsubscribe
// @Description Unsubscribe user from author
// @ID          delete_subscriber
// @Tags        subscribers
// @Accept      json
// @Produce     json
// @Param       Subscription body     models.Subscription true "Subscription info with required UserID, AuthorID and Subscription ID"
// @Success     200          {object} models.EmptyStruct  "Subscriber was successfully unsubscribed"
// @Failure     400          {object} echo.HTTPError      "Bad request"
// @Failure     500          {object} echo.HTTPError      "Not deleted"
// @Security    ApiKeyAuth
// @Router      /subscribers [delete]
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
