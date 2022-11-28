package domain

import (
	"mime/multipart"
)

type ImageUseCase interface {
	CreateOrUpdateImage(image *multipart.FileHeader, oldFilename string) (string, error)
	GetImage(filename string) (string, error)
	GetBlurredImage(filename string) (string, error)
}
