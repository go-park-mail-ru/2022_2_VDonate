package imagesRepository

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	"github.com/esimov/stackblur-go"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/interface"
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

func (s S3) CreateOrUpdateImage(img *multipart.FileHeader, oldFilename string) (string, error) {
	idxNew := strings.Index(img.Filename, ".")
	if idxNew == -1 {
		return "", domain.ErrBadRequest
	}
	bucket := utils.GetMD5OfNumLast(img.Filename[:idxNew], s.symbolsToHash)

	if err := makeBucket(s.client, bucket, s.policy); err != nil {
		return "", err
	}

	file, err := img.Open()
	if err != nil {
		return "", domain.ErrFileOpen
	}

	if len(oldFilename) != 0 {
		idxOld := strings.Index(oldFilename, ".")
		if idxOld == -1 {
			return "", domain.ErrInternal
		}

		oldBucket := utils.GetMD5OfNumLast(oldFilename[:idxOld], s.symbolsToHash)
		if err = s.client.RemoveObject(oldBucket, oldFilename); err != nil {
			return "", err
		}
	}

	if _, err = s.client.PutObject(
		bucket,
		img.Filename,
		file,
		img.Size,
		minio.PutObjectOptions{ContentType: img.Header.Get("Content-Type")},
	); err != nil {
		return "", err
	}

	if _, err = file.Seek(0, 0); err != nil {
		return "", err
	}

	blurImage, format, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	blurImageNRGBA, _ := stackblur.Process(blurImage, 500)

	blurFile := new(bytes.Buffer)

	switch format {
	case "jpeg":
		if err = jpeg.Encode(blurFile, blurImageNRGBA.SubImage(blurImageNRGBA.Rect), &jpeg.Options{Quality: 100}); err != nil {
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
		"blur_"+img.Filename,
		blurFile,
		int64(blurFile.Len()),
		minio.PutObjectOptions{ContentType: img.Header.Get("Content-Type")},
	)

	return img.Filename, err
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
