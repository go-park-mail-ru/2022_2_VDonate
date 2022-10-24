package httpImages

import (
	"net/http"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	imagesUseCase domain.ImageUseCase
}

func NewHandler(imagesUseCase domain.ImageUseCase) *Handler {
	return &Handler{
		imagesUseCase: imagesUseCase,
	}
}

func (h Handler) CreateImage(c echo.Context) error {
	bucket := c.Param("bucket")
	file, err := c.FormFile("file")
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	if err = h.imagesUseCase.CreateImage(file, bucket); err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	c.SetParamNames("bucket", "filename")
	c.SetParamValues(bucket, file.Filename)

	return h.GetImage(c)
}

func (h Handler) GetImage(c echo.Context) error {
	bucket := c.Param("bucket")
	filename := c.Param("filename")
	img, err := h.imagesUseCase.GetImage(bucket, filename)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoContent, err)
	}

	return c.JSON(http.StatusOK, struct {
		URL string `json:"url"`
	}{
		URL: img.Host + img.Path + "?" + img.RawQuery,
	})
}
