package images

import (
	"mime/multipart"
	"strings"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type usecase struct {
	ImageRepo domain.ImagesRepository
}

func New(imageRepo domain.ImagesRepository) domain.ImageUseCase {
	return &usecase{
		ImageRepo: imageRepo,
	}
}

func (u usecase) CreateImage(image *multipart.FileHeader) (string, error) {
	idx := strings.IndexByte(image.Filename, '.')
	if idx == -1 {
		return "", domain.ErrBadRequest
	}
	image.Filename = uuid.New().String() + image.Filename[idx:]
	return u.ImageRepo.CreateImage(image)
}

func (u usecase) GetImage(filename string) (string, error) {
	if len(filename) == 0 {
		return "", nil
	}

	newURL, err := u.ImageRepo.GetPermanentImage(filename)
	if err != nil {
		return "", err
	}

	return "https://wsrv.nl/?url=" + strings.ReplaceAll(newURL, "vdonate.ml", "95.163.209.195"), nil
}

func GetFileFromContext(c echo.Context) (*multipart.FileHeader, error) {
	return c.FormFile("file")
}
