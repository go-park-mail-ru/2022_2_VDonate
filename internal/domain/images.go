package domain

import (
	"mime/multipart"
)

type ImageUseCase interface {
	CreateImage(image *multipart.FileHeader, bucket string) (string, error)
	GetImage(bucket, name string) (string, error)
}
