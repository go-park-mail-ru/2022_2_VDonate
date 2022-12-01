package grpcImages

import (
	"context"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/images/protobuf"
)

type ImagesService struct {
	imageRepo domain.ImagesRepository
	protobuf.UnimplementedImagesServer
}

func New(r domain.ImagesRepository) protobuf.ImagesServer {
	return &ImagesService{
		imageRepo: r,
	}
}

func (s ImagesService) Create(_ context.Context, img *protobuf.Image) (*protobuf.Filename, error) {
	image, err := s.imageRepo.CreateOrUpdateImage(img.Filename, img.Content, img.Size, img.OldFilename)
	if err != nil {
		return nil, err
	}

	return &protobuf.Filename{Filename: image}, nil
}

func (s ImagesService) Get(_ context.Context, filename *protobuf.Filename) (*protobuf.URL, error) {
	url, err := s.imageRepo.GetPermanentImage(filename.GetFilename())
	if err != nil {
		return nil, err
	}

	return &protobuf.URL{
		Url: url,
	}, nil
}
