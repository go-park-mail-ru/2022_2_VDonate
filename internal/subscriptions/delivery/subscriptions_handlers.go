package httpSubscriptions

import (
	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"

	"net/http"
	"strconv"
)

type Handler struct {
	subscriptionsUsecase domain.SubscriptionsUseCase
	userUsecase          domain.UsersUseCase
	imageUsecase         domain.ImageUseCase

	bucket string
}

func NewHandler(s domain.SubscriptionsUseCase, u domain.UsersUseCase, i domain.ImageUseCase, bucket string) *Handler {
	return &Handler{
		subscriptionsUsecase: s,
		userUsecase:          u,
		imageUsecase:         i,
		bucket:               bucket,
	}
}

// GetSubscriptions godoc
// @Summary     Get user's subscriptions
// @Description Get user's subscriptions by Cookie
// @ID          get_subscriptions
// @Tags        subscriptions
// @Produce     json
// @Success     200 {object} []models.Subscription "Successfully received subscriptions"
// @Failure     400 {object} echo.HTTPError        "Bad request"
// @Failure     401 {object} echo.HTTPError        "No session"
// @Failure     403 {object} echo.HTTPError        "You are not supposed to make this requests"
// @Failure     500 {object} echo.HTTPError        "Internal error"
// @Security    ApiKeyAuth
// @Router      /subscriptions [get]
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

	for _, subscription := range s {
		url, errGetImage := h.imageUsecase.GetImage(h.bucket, subscription.Img)
		if errGetImage != nil {
			return errGetImage
		}
		subscription.Img = url.String()
	}

	return c.JSON(http.StatusOK, s)
}

// GetSubscription godoc
// @Summary     Get user's subscription
// @Description Get user's subscription by id
// @ID          get_subscription
// @Tags        subscriptions
// @Produce     json
// @Param       id  path     integer             true "Subscription ID"
// @Success     200 {object} models.Subscription "Successfully received subscription"
// @Failure     400 {object} echo.HTTPError      "Bad request"
// @Failure     401 {object} echo.HTTPError      "No session"
// @Failure     403 {object} echo.HTTPError      "You are not supposed to make this requests"
// @Failure     404 {object} echo.HTTPError      "Subscription not found"
// @Failure     500 {object} echo.HTTPError      "Internal error"
// @Security    ApiKeyAuth
// @Router      /subscriptions/{id} [get]
func (h Handler) GetSubscription(c echo.Context) error {
	subID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	s, err := h.subscriptionsUsecase.GetSubscriptionsByID(subID)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	url, err := h.imageUsecase.GetImage(h.bucket, s.Img)
	if err != nil {
		return err
	}
	s.Img = url.String()

	return c.JSON(http.StatusOK, s)
}

// CreateSubscription godoc
// @Summary     Create subscription
// @Description Create subscription by user's Cookie
// @ID          create_subscription
// @Tags        subscriptions
// @Produce     json
// @Accept      mpfd
// @Param       data formData models.AuthorSubscriptionMpfd true  "POST request of all information about `User`"
// @Param       file formData file                          false "Upload image"
// @Success     200  {object} models.ResponseImage          "Successfully created subscription"
// @Failure     400  {object} echo.HTTPError                "Bad request"
// @Failure     401  {object} echo.HTTPError                "No session"
// @Failure     403  {object} echo.HTTPError                "You are not supposed to make this requests"
// @Failure     404  {object} echo.HTTPError                "Subscription not found"
// @Failure     500  {object} echo.HTTPError                "Internal error"
// @Security    ApiKeyAuth
// @Router      /subscriptions [post]
func (h Handler) CreateSubscription(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	author, err := h.userUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	file, err := images.GetFileFromContext(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	var s models.AuthorSubscription

	if err = c.Bind(&s); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	newFile, err := h.imageUsecase.CreateImage(file, h.bucket)
	if err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	s.Img = newFile
	s.AuthorID = author.ID
	if err = h.subscriptionsUsecase.AddSubscription(s); err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	return c.JSON(http.StatusOK, models.ResponseImage{
		ImgPath: s.Img,
	})
}

// UpdateSubscription godoc
// @Summary     Update subscription
// @Description Update subscription by subscription id
// @ID          update_subscription
// @Tags        subscriptions
// @Produce     json
// @Accept      mpfd
// @Param       id   path     integer                       true  "Subscription ID"
// @Param       data formData models.AuthorSubscriptionMpfd true  "POST request of all information about `User`"
// @Param       file formData file                          false "Upload image"
// @Success     200  {object} models.ResponseImage          "Successfully updated subscription"
// @Failure     400  {object} echo.HTTPError                "Bad request"
// @Failure     401  {object} echo.HTTPError                "No session"
// @Failure     403  {object} echo.HTTPError                "You are not supposed to make this requests"
// @Failure     500  {object} echo.HTTPError                "Internal / update error"
// @Security    ApiKeyAuth
// @Router      /subscriptions/{id} [put]
func (h Handler) UpdateSubscription(c echo.Context) error {
	subID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}
	file, err := images.GetFileFromContext(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	var s models.AuthorSubscription

	if err = c.Bind(&s); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	newFile, err := h.imageUsecase.CreateImage(file, h.bucket)
	if err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	s.ID = subID
	s.Img = newFile
	if err = h.subscriptionsUsecase.UpdateSubscription(s); err != nil {
		return utils.WrapEchoError(domain.ErrUpdate, err)
	}

	return c.JSON(http.StatusOK, models.ResponseImage{
		ImgPath: s.Img,
	})
}

// DeleteSubscription godoc
// @Summary     Delete subscription
// @Description Delete subscription by subscription id
// @ID          delete_subscription
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
// @Router      /subscriptions/{id} [delete]
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
