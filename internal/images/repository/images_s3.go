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

const (
	expire = time.Hour
)

func getPolicy(client *minio.Client, bucket, policy string) error {
	bucketPolicy, err := client.GetBucketPolicy(bucket)
	if err != nil || bucketPolicy != policy {
		if err = client.SetBucketPolicy(bucket, policy); err != nil {
			return err
		}
	}
	return nil
}

func existBucket(client *minio.Client, bucket, policy string) error {
	exists, err := client.BucketExists(bucket)
	if err == nil || exists {
		if err = getPolicy(client, bucket, policy); err != nil {
			return err
		}
		return nil
	}
	return err
}

func makeBucket(client *minio.Client, bucket, policy string) error {
	if client.MakeBucket(bucket, "") != nil {
		if err := existBucket(client, bucket, policy); err != nil {
			return err
		}
	}
	return nil
}

func New(endpoint, accessKeyID, secretAccessKey string, secure bool, buckets map[string]string) (*S3, error) {
	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, secure)
	if err != nil {
		return nil, err
	}

	for bucket, policy := range buckets {
		if err = makeBucket(client, bucket, policy); err != nil {
			return nil, err
		}
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

	_, err = s.client.PutObject(
		bucket,
		image.Filename,
		file,
		image.Size,
		minio.PutObjectOptions{ContentType: image.Header.Get("Content-Type")},
	)

	return err
}

func (s *S3) GetImage(bucket, filename string) (*url.URL, error) {
	image, err := s.client.GetObject(bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	if _, err = image.Stat(); err != nil {
		return nil, err
	}

	return s.client.PresignedGetObject(bucket, filename, expire, url.Values{})
}

func (s *S3) GetPermanentImage(bucket, filename string) (string, error) {
	urlImage, err := s.client.PresignedGetObject(bucket, filename, expire, url.Values{})
	if err != nil {
		return "", err
	}
	return urlImage.Scheme + "://" + urlImage.Host + "/" + bucket + "/" + filename, nil
}
