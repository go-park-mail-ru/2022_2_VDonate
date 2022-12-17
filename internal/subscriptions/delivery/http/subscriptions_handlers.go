package httpSubscriptions

import (
	"errors"
	"net/http"
	"sort"
	"strconv"

	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/http"

	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	subscriptionsUsecase domain.SubscriptionsUseCase
	userUsecase          domain.UsersUseCase
	imageUsecase         domain.ImageUseCase
}

func NewHandler(s domain.SubscriptionsUseCase, u domain.UsersUseCase, i domain.ImageUseCase) *Handler {
	return &Handler{
		subscriptionsUsecase: s,
		userUsecase:          u,
		imageUsecase:         i,
	}
}

// GetSubscriptions godoc
// @Summary     Get User subscriptions
// @Description Get User subscriptions by Cookie
// @ID          get_user_subscriptions
// @Tags        subscriptions
// @Produce     json
// @Param       user_id query    integer                     true "User ID"
// @Success     200     {object} []models.AuthorSubscription "Successfully received subscriptions"
// @Failure     400     {object} echo.HTTPError              "Bad request"
// @Failure     401     {object} echo.HTTPError              "No session"
// @Failure     403     {object} echo.HTTPError              "You are not supposed to make this requests"
// @Failure     500     {object} echo.HTTPError              "Internal error"
// @Security    ApiKeyAuth
// @Router      /subscriptions [get]
func (h Handler) GetSubscriptions(c echo.Context) error {
	userID, err := strconv.ParseUint(c.QueryParam("user_id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	s, err := h.subscriptionsUsecase.GetSubscriptionsByUserID(userID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, s)
}

// GetAuthorSubscriptions godoc
// @Summary     Get Author subscriptions
// @Description Get Author subscriptions by author ID
// @ID          get_author_subscriptions
// @Tags        subscriptions
// @Produce     json
// @Param       author_id query    integer                         true "Author ID"
// @Success     200       {object} []models.AuthorSubscriptionMpfd "Successfully received subscriptions"
// @Failure     400       {object} echo.HTTPError                  "Bad request"
// @Failure     401       {object} echo.HTTPError                  "No session"
// @Failure     403       {object} echo.HTTPError                  "You are not supposed to make this requests"
// @Failure     500       {object} echo.HTTPError                  "Internal error"
// @Security    ApiKeyAuth
// @Router      /subscriptions/author [get]
func (h Handler) GetAuthorSubscriptions(c echo.Context) error {
	authorID, err := strconv.ParseUint(c.QueryParam("author_id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	s, err := h.subscriptionsUsecase.GetAuthorSubscriptionsByAuthorID(authorID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].Price < s[j].Price
	})

	return c.JSON(http.StatusOK, s)
}

// GetAuthorSubscription godoc
// @Summary     Get Author subscription
// @Description Get Author subscription by id
// @ID          get_author_subscription
// @Tags        subscriptions
// @Produce     json
// @Param       id  path     integer                       true "Subscription ID"
// @Success     200 {object} models.AuthorSubscriptionMpfd "Successfully received subscription"
// @Failure     400 {object} echo.HTTPError                "Bad request"
// @Failure     401 {object} echo.HTTPError                "No session"
// @Failure     403 {object} echo.HTTPError                "You are not supposed to make this requests"
// @Failure     404 {object} echo.HTTPError                "Subscription not found"
// @Failure     500 {object} echo.HTTPError                "Internal error"
// @Security    ApiKeyAuth
// @Router      /subscriptions/author/{id} [get]
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

// CreateAuthorSubscription godoc
// @Summary     Create Author subscription
// @Description Create Author subscription by user's Cookie
// @ID          create_author_subscription
// @Tags        subscriptions
// @Produce     json
// @Accept      mpfd
// @Param       data formData models.AuthorSubscriptionMpfd    true  "POST request of all information about `User`"
// @Param       file formData file                             false "Upload image"
// @Success     200  {object} models.ResponseImageSubscription "Successfully created subscription"
// @Failure     400  {object} echo.HTTPError                   "Bad request"
// @Failure     401  {object} echo.HTTPError                   "No session"
// @Failure     403  {object} echo.HTTPError                   "You are not supposed to make this requests"
// @Failure     500  {object} echo.HTTPError                   "Internal error"
// @Security    ApiKeyAuth
// @Router      /subscriptions/author [post]
func (h Handler) CreateAuthorSubscription(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	author, err := h.userUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	var s models.AuthorSubscription
	if err = c.Bind(&s); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	file, err := images.GetFileFromContext(c)
	if file != nil && !errors.Is(err, http.ErrMissingFile) {
		if s.Img, err = h.imageUsecase.CreateOrUpdateImage(file, ""); err != nil {
			return errorHandling.WrapEcho(domain.ErrCreate, err)
		}
	}

	id, err := h.subscriptionsUsecase.AddAuthorSubscription(s, author.ID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrCreate, err)
	}

	if len(s.Img) != 0 {
		if s.Img, err = h.imageUsecase.GetImage(s.Img); err != nil {
			return errorHandling.WrapEcho(domain.ErrInternal, err)
		}
	}

	return c.JSON(http.StatusOK, models.ResponseImageSubscription{
		SubscriptionID: id,
		ImgPath:        s.Img,
	})
}

// UpdateAuthorSubscription godoc
// @Summary     Update Author subscription
// @Description Update Author subscription by subscription id
// @ID          update_author_subscription
// @Tags        subscriptions
// @Produce     json
// @Accept      mpfd
// @Param       id   path     integer                          true  "Subscription ID"
// @Param       data formData models.AuthorSubscriptionMpfd    true  "POST request of all information about `User`"
// @Param       file formData file                             false "Upload image"
// @Success     200  {object} models.ResponseImageSubscription "Successfully updated subscription"
// @Failure     400  {object} echo.HTTPError                   "Bad request"
// @Failure     401  {object} echo.HTTPError                   "No session"
// @Failure     403  {object} echo.HTTPError                   "You are not supposed to make this requests"
// @Failure     500  {object} echo.HTTPError                   "Internal / update error"
// @Security    ApiKeyAuth
// @Router      /subscriptions/author/{id} [put]
func (h Handler) UpdateAuthorSubscription(c echo.Context) error {
	subID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	var s models.AuthorSubscription

	if err = c.Bind(&s); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	file, err := images.GetFileFromContext(c)
	if file != nil && !errors.Is(err, http.ErrMissingFile) {
		if s.Img, err = h.imageUsecase.CreateOrUpdateImage(file, s.Img); err != nil {
			return errorHandling.WrapEcho(domain.ErrCreate, err)
		}
	}

	if err = h.subscriptionsUsecase.UpdateAuthorSubscription(s, subID); err != nil {
		return errorHandling.WrapEcho(domain.ErrUpdate, err)
	}

	if len(s.Img) != 0 {
		if s.Img, err = h.imageUsecase.GetImage(s.Img); err != nil {
			return errorHandling.WrapEcho(domain.ErrInternal, err)
		}
	}

	return c.JSON(http.StatusOK, models.ResponseImageSubscription{
		SubscriptionID: subID,
		ImgPath:        s.Img,
	})
}

// DeleteAuthorSubscription godoc
// @Summary     Delete Author subscription
// @Description Delete Author subscription by subscription id
// @ID          delete_author_subscription
// @Tags        subscriptions
// @Produce     json
// @Accept      mpfd
// @Param       id  path     integer            true "Subscription ID"
// @Success     200 {object} models.EmptyStruct "Successfully updated subscription"
// @Failure     400 {object} echo.HTTPError     "Bad request"
// @Failure     401 {object} echo.HTTPError     "No session"
// @Failure     403 {object} echo.HTTPError     "You are not supposed to make this requests"
// @Failure     500 {object} echo.HTTPError     "Internal / delete error"
// @Security    ApiKeyAuth
// @Router      /subscriptions/author/{id} [delete]
func (h Handler) DeleteAuthorSubscription(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	if err = h.subscriptionsUsecase.DeleteAuthorSubscription(id); err != nil {
		return errorHandling.WrapEcho(domain.ErrDelete, err)
	}

	return c.JSON(http.StatusOK, models.EmptyStruct{})
}
