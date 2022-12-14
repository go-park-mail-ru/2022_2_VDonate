package domain

import (
	"mime/multipart"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type UsersMicroservice interface {
	Create(user models.User) (uint64, error)
	GetByUsername(username string) (models.User, error)
	GetByID(id uint64) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetBySessionID(sessionID string) (models.User, error)
	GetUserByPostID(postID uint64) (models.User, error)
	Update(user models.User) error
	GetAuthorByUsername(username string) ([]models.User, error)
	GetAllAuthors() ([]models.User, error)
	GetPostsNum(UserID uint64) (uint64, error)
	GetSubscribersNum(UserID uint64) (uint64, error)
	GetProfitForMounth(UserID uint64) (uint64, error)
	DropBalance(userID uint64) error
}

type UsersUseCase interface {
	GetByUsername(username string) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetByID(id uint64) (models.User, error)
	GetBySessionID(sessionID string) (models.User, error)
	GetUserByPostID(postID uint64) (models.User, error)
	Create(user models.User) (uint64, error)
	Update(user models.User, file *multipart.FileHeader, id uint64) (models.User, error)
	CheckIDAndPassword(id uint64, password string) bool
	IsExistUsernameAndEmail(username, email string) bool
	FindAuthors(keyword string) ([]models.User, error)
	GetPostsNum(userID uint64) (uint64, error)
	GetSubscribersNum(userID uint64) (uint64, error)
	GetProfitForMounth(userID uint64) (uint64, error)
}
