package images

import (
	"mime/multipart"
	"strings"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type usecase struct {
	ImageMicroservice domain.ImageMicroservice
}

func New(i domain.ImageMicroservice) domain.ImageUseCase {
	return &usecase{
		ImageMicroservice: i,
	}
}

func (u usecase) CreateOrUpdateImage(image *multipart.FileHeader, oldFilename string) (string, error) {
	idx := strings.IndexByte(image.Filename, '.')
	if idx == -1 {
		return "", domain.ErrBadRequest
	}
	image.Filename = uuid.New().String() + image.Filename[idx:]

	file, err := image.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, image.Size)
	if _, err = file.Read(buffer); err != nil {
		return "", err
	}

	return u.ImageMicroservice.Create(image.Filename, buffer, image.Size, oldFilename)
}

func (u usecase) GetImage(filename string) (string, error) {
	if len(filename) == 0 {
		return "", nil
	}

	newURL, err := u.ImageMicroservice.Get(filename)
	if err != nil {
		return "", err
	}

	return newURL, nil
}

func (u usecase) GetBlurredImage(filename string) (string, error) {
	filename = "blur_" + filename
	return u.GetImage(filename)
}

func GetFileFromContext(c echo.Context) (*multipart.FileHeader, error) {
	return c.FormFile("file")
}
