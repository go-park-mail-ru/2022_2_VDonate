package images

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	"mime/multipart"
	"net/url"
)

type usecase struct {
	ImageRepo domain.ImagesRepository
}

func New(imageRepo domain.ImagesRepository) domain.ImageUseCase {
	return &usecase{
		ImageRepo: imageRepo,
	}
}

func (u usecase) CreateImage(image *multipart.FileHeader, bucket string) error {
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
