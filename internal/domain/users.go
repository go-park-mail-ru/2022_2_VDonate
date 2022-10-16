package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type UsersUseCase interface {
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByID(id uint64) (*models.User, error)
	GetBySessionID(sessionID string) (*models.User, error)
	GetUserByPostID(postID uint64) (*models.User, error)
	Create(user models.User) (*models.User, error)
	Update(user models.User) (*models.User, error)
	DeleteByID(id uint64) error
	DeleteByUsername(username string) error
	DeleteByEmail(email string) error
	CheckIDAndPassword(id uint64, password string) bool
	IsExistUsernameAndEmail(username, email string) bool
}

type UsersRepository interface {
	Create(user *models.User) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByID(id uint64) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetBySessionID(sessionID string) (*models.User, error)
	GetUserByPostID(postID uint64) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	DeleteByID(id uint64) error
	Close() error
}