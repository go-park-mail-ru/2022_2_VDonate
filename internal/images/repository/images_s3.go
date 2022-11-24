package imagesRepository

import (
	"errors"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/minio/minio-go"
)

type S3 struct {
	client *minio.Client

	policy        string
	expire        time.Duration
	symbolsToHash int
}

func setPolicy(client *minio.Client, bucket, policy string) error {
	if err := client.SetBucketPolicy(bucket, strings.ReplaceAll(policy, "$(bucket)", bucket)); err != nil {
		if errRm := client.RemoveBucket(bucket); errRm != nil {
			return errRm
		}
		return err
	}

	return nil
}

func makeBucket(client *minio.Client, bucket, policy string) error {
	exists, err := client.BucketExists(bucket)
	if err != nil {
		return err
	}

	if !exists {
		if err = client.MakeBucket(bucket, ""); err != nil {
			return err
		}

		if err = setPolicy(client, bucket, policy); err != nil {
			return err
		}
	}

	return nil
}

func New(endpoint, accessKeyID, secretAccessKey string, secure bool, sth int, p string, e int) (*S3, error) {
	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, secure)
	if err != nil {
		return nil, err
	}

	return &S3{
		client: client,

		symbolsToHash: sth,
		policy:        p,
		expire:        time.Duration(e) * time.Minute,
	}, nil
}

func (s S3) CreateImage(image *multipart.FileHeader) (string, error) {
	idx := strings.Index(image.Filename, ".")
	if idx == -1 {
		return "", errors.New("bad url")
	}
	bucket := utils.GetMD5OfNumLast(image.Filename[:idx], s.symbolsToHash)

	if err := makeBucket(s.client, bucket, s.policy); err != nil {
		return "", err
	}

	file, err := image.Open()
	if err != nil {
		return "", domain.ErrFileOpen
	}

	_, err = s.client.PutObject(
		bucket,
		image.Filename,
		file,
		image.Size,
		minio.PutObjectOptions{ContentType: image.Header.Get("Content-Type")},
	)

	return image.Filename, err
}

func (s S3) GetPermanentImage(filename string) (string, error) {
	idx := strings.Index(filename, ".")
	if idx == -1 {
		return "", errors.New("bad url")
	}
	bucket := utils.GetMD5OfNumLast(filename[:idx], s.symbolsToHash)
	urlImage, err := s.client.PresignedGetObject(bucket, filename, s.expire, url.Values{})
	if err != nil {
		return "", err
	}
	return urlImage.Scheme + "://" + urlImage.Host + "/" + bucket + "/" + filename, nil
}
