package domain

import (
	"mime/multipart"
	"net/url"
)

type ImageUseCase interface {
	CreateImage(image *multipart.FileHeader, bucket string) (string, error)
	GetImage(bucket, name string) (*url.URL, error)
}
