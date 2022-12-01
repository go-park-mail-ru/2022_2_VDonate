package httpImages

import (
	"net/http"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"

	echo "github.com/labstack/echo/v4"
)

type Handler struct {
	imageUseCase domain.ImageUseCase
}

func NewHandler(i domain.ImageUseCase) *Handler {
	return &Handler{
		imageUseCase: i,
	}
}

// CreateOrUpdateImage godoc
// @Summary     Create image
// @Description Create image POST request
// @ID          create_image
// @Tags        images
// @Produce     json
// @Success     200 {object} models.ImageMpfd "Posts were successfully received"
// @Failure     400 {object} echo.HTTPError   "Bad request"
// @Failure     401 {object} echo.HTTPError   "No session provided"
// @Failure     500 {object} echo.HTTPError   "Internal error"
// @Security    ApiKeyAuth
// @Router      /image [post]
func (h Handler) CreateOrUpdateImage(c echo.Context) error {
	file, err := images.GetFileFromContext(c)
	if err != nil || file == nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	image, err := h.imageUseCase.CreateOrUpdateImage(file, "")
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrCreate, err)
	}

	url, err := h.imageUseCase.GetImage(image)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, models.ImageMpfd{
		URL: url,
	})
}
