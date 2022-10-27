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
// @Failure     404 {object} echo.HTTPError "User / avatars not found"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /users/{id} [get]
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

// PutUser godoc
// @Summary     Update user
// @Description Update user by user id
// @ID          update_user
// @Tags        users
// @Produce     json
// @Param       id path integer true "User ID"
// @Accept      mpfd
// @Param       post formData models.UserMpfd true  "New Post"
// @Param       file formData file            false "Uploaded file"
// @Success     200  {object} models.User     "User was successfully updated"
// @Failure     400  {object} echo.HTTPError  "Bad request"
// @Failure     401  {object} echo.HTTPError  "No session provided"
// @Failure     403  {object} echo.HTTPError  "Not a user"
// @Failure     500  {object} echo.HTTPError  "Internal error / failed to create user"
// @Security    ApiKeyAuth
// @Router      /users/{id} [put]
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
