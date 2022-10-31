package images

import (
	"errors"
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

func (u usecase) CreateImage(image *multipart.FileHeader, bucket string) (string, error) {
	image.Filename = uuid.New().String() + image.Filename[strings.IndexByte(image.Filename, '.'):]
	return u.ImageRepo.CreateImage(image, bucket)
}

func (u usecase) GetImage(bucket, name string) (string, error) {
	if len(name) == 0 {
		return "", nil
	}
	switch bucket {
	case "image":
		newURL, err := u.ImageRepo.GetImage(bucket, name)
		if err != nil {
			return "", err
		}
		return newURL.String(), err
	case "avatar":
		return u.ImageRepo.GetPermanentImage(bucket, name)
	default:
		return "", errors.New("bad url")
	}
}

func GetFileFromContext(c echo.Context) (*multipart.FileHeader, error) {
	return c.FormFile("file")
}
