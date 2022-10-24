package imagesRepository

import (
	"mime/multipart"
	"net/url"
	"time"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/minio/minio-go"
)

type S3 struct {
	client *minio.Client
}

func New(endpoint, accessKeyID, secretAccessKey string, secure bool) (*S3, error) {
	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, secure)
	if err != nil {
		return nil, err
	}

	return &S3{
		client: client,
	}, nil
}

func (s *S3) CreateImage(image *multipart.FileHeader, bucket string) error {
	exists, err := s.client.BucketExists(bucket)
	if err != nil {
		return err
	}

	if !exists {
		return domain.ErrBucketNotExists
	}

	file, err := image.Open()
	if err != nil {
		return domain.ErrFileOpen
	}

	if _, err = s.client.PutObject(
		bucket,
		image.Filename,
		file,
		image.Size,
		minio.PutObjectOptions{ContentType: image.Header.Get("Content-Type")},
	); err != nil {
		return err
	}

	return nil
}

func (s *S3) GetImage(bucket, filename string, expires time.Duration) (*url.URL, error) {
	image, err := s.client.GetObject(bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	if _, err = image.Stat(); err != nil {
		return nil, err
	}

	object, err := s.client.PresignedGetObject(bucket, filename, expires, url.Values{})
	if err != nil {
		return nil, err
	}

	return object, err
}
