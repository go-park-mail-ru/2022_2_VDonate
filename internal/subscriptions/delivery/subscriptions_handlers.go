package httpSubscriptions

import (
	"net/http"
	"strconv"
	"strings"

	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

func (h Handler) CreateSubscription(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	author, err := h.userUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}
	file.Filename = uuid.New().String() + file.Filename[strings.IndexByte(file.Filename, '.'):]

	var s models.AuthorSubscription

	if err = c.Bind(&s); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	if err = h.imageUsecase.CreateImage(file, h.bucket); err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	s.Img = file.Filename
	s.AuthorID = author.ID
	if err = h.subscriptionsUsecase.AddSubscription(s); err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}

func (h Handler) UpdateSubscription(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}
	file.Filename = uuid.New().String() + file.Filename[strings.IndexByte(file.Filename, '.'):]

	var s models.AuthorSubscription

	if err = c.Bind(&s); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	if err = h.imageUsecase.CreateImage(file, h.bucket); err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	if err = h.subscriptionsUsecase.UpdateSubscription(s); err != nil {
		return utils.WrapEchoError(domain.ErrUpdate, err)
	}

	return c.JSON(http.StatusOK, struct{}{})
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
