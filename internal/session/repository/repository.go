package sessionRepository

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type API interface {
	GetByValue(value string) (*models.Cookie, error)
	GetByUserID(id uint) (*models.Cookie, error)
	Create(cookie *models.Cookie) (*models.Cookie, error)
	DeleteByValue(value string) error
	DeleteByUserID(id uint) error
	Close() error
}
