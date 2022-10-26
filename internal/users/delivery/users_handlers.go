package httpUsers

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"

	"net/http"
	"strconv"
)

type Handler struct {
	sessionUseCase domain.AuthUseCase
	userUseCase    domain.UsersUseCase
	imageUseCase   domain.ImageUseCase

	bucket string
}

func NewHandler(
	userUseCase domain.UsersUseCase,
	sessionUseCase domain.AuthUseCase,
	imageUseCase domain.ImageUseCase,
	bucket string,
) *Handler {
	return &Handler{
		userUseCase:    userUseCase,
		sessionUseCase: sessionUseCase,
		imageUseCase:   imageUseCase,

		bucket: bucket,
	}
}

func (h *Handler) GetUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	user, err := h.userUseCase.GetByID(id)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	url, err := h.imageUseCase.GetImage(h.bucket, user.Avatar)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}
	user.Avatar = url.String()

	return UserResponse(c, user)
}

func (h *Handler) PutUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	file, err := images.GetFileFromContext(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	var updateUser models.User

	if err = c.Bind(&updateUser); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	newFile, err := h.imageUseCase.CreateImage(file, h.bucket)
	if err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	updateUser.Avatar = newFile
	updateUser.ID = id
	if err = h.userUseCase.Update(updateUser); err != nil {
		return utils.WrapEchoError(domain.ErrUpdate, err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}
