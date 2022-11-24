package domain

import (
	"mime/multipart"
)

type ImageUseCase interface {
	CreateImage(image *multipart.FileHeader) (string, error)
	GetImage(filename string) (string, error)
}
