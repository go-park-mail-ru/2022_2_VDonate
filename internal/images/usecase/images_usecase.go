package images

import (
	"mime/multipart"
	"net/url"
	"time"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
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
	return u.ImageRepo.GetImage(bucket, name, time.Hour)
}
