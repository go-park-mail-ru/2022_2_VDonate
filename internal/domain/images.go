package domain

import (
	"mime/multipart"
)

type ImageMicroservice interface {
	Create(filename string, file []byte, size int64, oldFilename string) (string, error)
	Get(filename string) (string, error)
}
type ImageUseCase interface {
	CreateOrUpdateImage(image *multipart.FileHeader, oldFilename string) (string, error)
	GetImage(filename string) (string, error)
	GetBlurredImage(filename string) (string, error)
}
