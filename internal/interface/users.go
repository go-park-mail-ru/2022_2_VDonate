package domain

import (
	"mime/multipart"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type UsersDomain interface {
	GetByUsername(username string) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetByID(id uint64) (models.User, error)
	GetBySessionID(sessionID string) (models.User, error)
	GetUserByPostID(postID uint64) (models.User, error)
	Create(user models.User) (uint64, error)
	Update(user models.User, id uint64) (models.User, error)
	DeleteByID(id uint64) error
	DeleteByUsername(username string) error
	DeleteByEmail(email string) error
	CheckIDAndPassword(id uint64, password string) bool
	IsExistUsernameAndEmail(username, email string) bool
}

type UsersUseCase interface {
	GetByUsername(username string) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetByID(id uint64) (models.User, error)
	GetBySessionID(sessionID string) (models.User, error)
	GetUserByPostID(postID uint64) (models.User, error)
	Create(user models.User) (uint64, error)
	Update(user models.User, file *multipart.FileHeader, id uint64) (models.User, error)
	DeleteByID(id uint64) error
	DeleteByUsername(username string) error
	DeleteByEmail(email string) error
	CheckIDAndPassword(id uint64, password string) bool
	IsExistUsernameAndEmail(username, email string) bool
}
