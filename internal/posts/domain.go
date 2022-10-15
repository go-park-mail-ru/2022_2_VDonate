package postsDomain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type UseCase interface {
	GetPostsByUserID(id uint64) ([]*models.PostDB, error)
	GetPostByID(postID uint64) (*models.PostDB, error)
	Create(post models.PostDB) (*models.PostDB, error)
	Update(post models.PostDB) (*models.PostDB, error)
	DeleteByID(postID uint64) error
}

type Repository interface {
	GetAllByUserID(userID uint64) ([]*models.PostDB, error)
	GetPostByUserID(userID, postID uint64) (*models.PostDB, error)
	GetPostByID(postID uint64) (*models.PostDB, error)
	Create(post models.PostDB) (*models.PostDB, error)
	Update(post models.PostDB) (*models.PostDB, error)
	DeleteByID(postID uint64) error
	Close() error
}
