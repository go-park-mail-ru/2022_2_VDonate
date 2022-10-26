package images

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"mime/multipart"
	"net/url"
	"strings"
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

func (u usecase) GetImage(bucket, name string) (*url.URL, error) {
	switch bucket {
	case "img":
		return u.ImageRepo.GetImage(bucket, name)
	case "avatar":
		urlImage, err := u.ImageRepo.GetPermanentImage(bucket, name)
		if err != nil {
			return nil, err
		}
		return url.Parse(urlImage)
	default:
		return u.ImageRepo.GetImage(bucket, name)
	}
}

func GetFileFromContext(c echo.Context) (*multipart.FileHeader, error) {
	return c.FormFile("file")
}
