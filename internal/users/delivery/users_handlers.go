package httpUsers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	sessionUseCase       domain.AuthUseCase
	userUseCase          domain.UsersUseCase
	imageUseCase         domain.ImageUseCase
	subscriptionsUseCase domain.SubscriptionsUseCase
	subscribersUseCase   domain.SubscribersUseCase
}

func NewHandler(
	userUseCase domain.UsersUseCase,
	sessionUseCase domain.AuthUseCase,
	imageUseCase domain.ImageUseCase,
	subscriptionsUseCase domain.SubscriptionsUseCase,
	subscribersUseCase domain.SubscribersUseCase,
) *Handler {
	return &Handler{
		userUseCase:          userUseCase,
		sessionUseCase:       sessionUseCase,
		imageUseCase:         imageUseCase,
		subscriptionsUseCase: subscriptionsUseCase,
		subscribersUseCase:   subscribersUseCase,
	}
}

// GetUser godoc
// @Summary     Get a User
// @Description Get a User information from server
// @ID          get_user
// @Tags        users
// @Produce     json
// @Param       id  path     integer        true "User ID"
// @Success     200 {object} models.Author  "User was successfully received"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     401 {object} echo.HTTPError "No session provided"
// @Failure     404 {object} echo.HTTPError "User not found"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /users/{id} [get]
func (h *Handler) GetUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	user, err := h.userUseCase.GetByID(id)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	if user.Avatar, err = h.imageUseCase.GetImage(fmt.Sprint(c.Get("bucket")), user.Avatar); err != nil {
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	var subscriptions []models.AuthorSubscription
	var subscribers []models.User

	if subscriptions, err = h.subscriptionsUseCase.GetSubscriptionsByUserID(user.ID); err != nil {
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	if user.IsAuthor {
		if subscribers, err = h.subscribersUseCase.GetSubscribers(user.ID); err != nil {
			return errorHandling.WrapEcho(domain.ErrInternal, err)
		}
	}

	user.CountSubscriptions = uint64(len(subscriptions))
	user.CountSubscribers = uint64(len(subscribers))

	return UserResponse(c, user)
}

// PutUser godoc
// @Summary     Update user
// @Description Update user by user id
// @ID          update_user
// @Tags        users
// @Produce     json
// @Param       id path integer true "User ID"
// @Accept      mpfd
// @Param       post formData models.UserMpfd           true  "New Post"
// @Param       file formData file                      false "Uploaded file"
// @Success     200  {object} models.ResponseImageUsers "User was successfully updated"
// @Failure     400  {object} echo.HTTPError            "Bad request"
// @Failure     401  {object} echo.HTTPError            "No session provided"
// @Failure     403  {object} echo.HTTPError            "Not a user"
// @Failure     500  {object} echo.HTTPError            "Internal error / failed to create user"
// @Security    ApiKeyAuth
// @Router      /users/{id} [put]
func (h *Handler) PutUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	file, err := images.GetFileFromContext(c)
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	var updateUser models.User

	if err = c.Bind(&updateUser); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	if file != nil {
		if updateUser.Avatar, err = h.imageUseCase.CreateImage(file, fmt.Sprint(c.Get("bucket"))); err != nil {
			return errorHandling.WrapEcho(domain.ErrCreate, err)
		}
	}

	if err = h.userUseCase.Update(updateUser, id); err != nil {
		return errorHandling.WrapEcho(domain.ErrUpdate, err)
	}

	return c.JSON(http.StatusOK, models.ResponseImageUsers{
		UserID:  id,
		ImgPath: updateUser.Avatar,
	})
}
