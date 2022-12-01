package imagesMicroservice

import (
	"context"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/images/protobuf"
)

type ImagesMicroservice struct {
	client protobuf.ImagesClient
}

func New(c protobuf.ImagesClient) domain.ImageMicroservice {
	return &ImagesMicroservice{
		client: c,
	}
}

func (m ImagesMicroservice) Create(filename string, file []byte, size int64, oldFilename string) (string, error) {
	name, err := m.client.Create(context.Background(), &protobuf.Image{
		Filename:    filename,
		Content:     file,
		Size:        size,
		OldFilename: oldFilename,
	})
	if err != nil {
		return "", err
	}

	return name.GetFilename(), nil
}

func (m ImagesMicroservice) Get(filename string) (string, error) {
	url, err := m.client.Get(context.Background(), &protobuf.Filename{
		Filename: filename,
	})
	if err != nil {
		return "", err
	}

	return url.GetUrl(), nil
}
