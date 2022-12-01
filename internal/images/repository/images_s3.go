package imagesRepository

import (
	"bufio"
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"net/url"
	"strings"
	"time"

	"github.com/esimov/stackblur-go"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/minio/minio-go"
)

const (
	quality    = 100
	blurRadius = 500
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

func (s S3) CreateOrUpdateImage(filename string, file []byte, size int64, oldFilename string) (string, error) {
	idxNew := strings.Index(filename, ".")
	if idxNew == -1 {
		return "", domain.ErrBadRequest
	}
	bucket := utils.GetMD5OfNumLast(filename[:idxNew], s.symbolsToHash)

	if len(oldFilename) != 0 {
		idxOld := strings.Index(oldFilename, ".")
		if idxOld == -1 {
			return "", domain.ErrInternal
		}

		oldBucket := utils.GetMD5OfNumLast(oldFilename[:idxOld], s.symbolsToHash)
		if err := s.client.RemoveObject(oldBucket, oldFilename); err != nil {
			return "", err
		}
	}

	err := makeBucket(s.client, bucket, s.policy)
	if err != nil {
		return "", err
	}

	fileReader := new(bytes.Buffer)
	fileReader.Write(file)

	r := bufio.NewReader(fileReader)

	if _, err = s.client.PutObject(
		bucket,
		filename,
		r,
		size,
		minio.PutObjectOptions{},
	); err != nil {
		return "", err
	}

	fileReader = new(bytes.Buffer)
	fileReader.Write(file)

	r.Reset(fileReader)

	blurImage, format, err := image.Decode(r)
	if err != nil {
		return "", err
	}

	blurImageNRGBA, _ := stackblur.Process(blurImage, blurRadius)

	blurFile := new(bytes.Buffer)

	switch format {
	case "jpeg", "jpg":
		if err = jpeg.Encode(blurFile, blurImageNRGBA.SubImage(blurImageNRGBA.Rect), &jpeg.Options{Quality: quality}); err != nil {
			return "", err
		}
	case "png":
		if err = png.Encode(blurFile, blurImageNRGBA.SubImage(blurImageNRGBA.Rect)); err != nil {
			return "", err
		}
	default:
		return "", errors.New("unknown format")
	}

	_, err = s.client.PutObject(
		bucket,
		"blur_"+filename,
		blurFile,
		int64(blurFile.Len()),
		minio.PutObjectOptions{},
	)

	return filename, err
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
