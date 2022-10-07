package postsRepository

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type API interface {
	GetAllByUserID(userID uint) ([]*models.PostDB, error)
	GetPostByUserID(userID, postID uint) (*models.PostDB, error)
	CreateInUserByID(post models.PostDB) error
	DeleteInUserByID(userID, postID uint) error
	Close() error
}
