package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type PostsUseCase interface {
	GetPostsByUserID(id uint64) ([]*models.Post, error)
	GetPostByID(postID uint64) (*models.Post, error)
	Create(post models.Post) (*models.Post, error)
	Update(post models.Post) (*models.Post, error)
	DeleteByID(postID uint64) error
}

type PostsRepository interface {
	GetAllByUserID(userID uint64) ([]*models.Post, error)
	GetPostByUserID(userID, postID uint64) (*models.Post, error)
	GetPostByID(postID uint64) (*models.Post, error)
	Create(post models.Post) (*models.Post, error)
	Update(post models.Post) (*models.Post, error)
	DeleteByID(postID uint64) error
	Close() error
}
