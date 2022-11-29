package httpImages

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	imageUseCase domain.ImageUseCase
}

func NewHandler(i domain.ImageUseCase) *Handler {
	return &Handler{
		imageUseCase: i,
	}
}

func (h Handler) CreateOrUpdateImage(c echo.Context) error {
	file, err := images.GetFileFromContext(c)
	if err != nil || file == nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	image, err := h.imageUseCase.CreateOrUpdateImage(file, "")
	if err != nil {
		return err
	}

	url, err := h.imageUseCase.GetImage(image)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		URL string `json:"url"`
	}{
		URL: url,
	})
}
