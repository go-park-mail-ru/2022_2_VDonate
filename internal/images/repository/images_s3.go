package imagesRepository

import (
	"bufio"
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/url"
	"strings"
	"time"

	"github.com/esimov/stackblur-go"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/minio/minio-go"
	"github.com/ztrue/tracerr"
)

const (
	quality      = 100
	blurRadius   = 500
	numColorsGIF = 256
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
	idxNew := strings.LastIndex(filename, ".")
	if idxNew == -1 {
		return "", domain.ErrBadRequest
	}
	bucket := utils.GetMD5OfNumLast(filename[:idxNew], s.symbolsToHash)

	fileFormat := filename[idxNew+1:]

	if len(oldFilename) != 0 {
		idxOld := strings.LastIndex(oldFilename, ".")
		if idxOld == -1 {
			return "", domain.ErrInternal
		}

		oldBucket := utils.GetMD5OfNumLast(oldFilename[:idxOld], s.symbolsToHash)
		if err := s.client.RemoveObject(oldBucket, oldFilename); err != nil {
			return "", tracerr.Wrap(err)
		}
	}

	err := makeBucket(s.client, bucket, s.policy)
	if err != nil {
		return "", tracerr.Wrap(err)
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
		return "", tracerr.Wrap(err)
	}

	fileReader = new(bytes.Buffer)
	fileReader.Write(file)

	r.Reset(fileReader)

	var blurImage image.Image

	switch fileFormat {
	case "png":
		blurImage, err = png.Decode(r)
	case "gif":
		blurImage, err = gif.Decode(r)
	default:
		blurImage, _, err = image.Decode(r)
	}

	if err != nil {
		return "", tracerr.Wrap(err)
	}

	blurImageNRGBA, _ := stackblur.Process(blurImage, blurRadius)

	blurFile := new(bytes.Buffer)

	switch fileFormat {
	case "jpeg", "jpg":
		if err = jpeg.Encode(blurFile, blurImageNRGBA.SubImage(blurImageNRGBA.Rect), &jpeg.Options{Quality: quality}); err != nil {
			return "", tracerr.Wrap(err)
		}
	case "png":
		if err = png.Encode(blurFile, blurImageNRGBA.SubImage(blurImageNRGBA.Rect)); err != nil {
			return "", tracerr.Wrap(err)
		}
	case "gif":
		if err = gif.Encode(blurFile, blurImageNRGBA.SubImage(blurImageNRGBA.Rect), &gif.Options{NumColors: numColorsGIF}); err != nil {
			return "", tracerr.Wrap(err)
		}
	default:
		return "", tracerr.Wrap(errors.New("unknown format"))
	}

	_, err = s.client.PutObject(
		bucket,
		"blur_"+filename,
		blurFile,
		int64(blurFile.Len()),
		minio.PutObjectOptions{},
	)

	return filename, tracerr.Wrap(err)
}

func (s S3) GetPermanentImage(filename string) (string, error) {
	idx := strings.Index(filename, ".")
	if idx == -1 {
		return "", errors.New("bad url")
	}
	bucket := utils.GetMD5OfNumLast(filename[:idx], s.symbolsToHash)
	urlImage, err := s.client.PresignedGetObject(bucket, filename, s.expire, url.Values{})
	if err != nil {
		return "", tracerr.Wrap(err)
	}
	return urlImage.Scheme + "://" + urlImage.Host + "/" + bucket + "/" + filename, nil
}
