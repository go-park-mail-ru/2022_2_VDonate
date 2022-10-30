package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type PostsUseCase interface {
	GetPostsByUserID(id uint64) ([]*models.Post, error)
	GetPostByID(postID uint64) (*models.Post, error)
	Create(post models.Post) (*models.Post, error)
	Update(post models.Post) (*models.Post, error)
	DeleteByID(postID uint64) error
	GetLikesByPostID(postID uint64) ([]models.Like, error)
	GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error)
	LikePost(userID, postID uint64) error
	UnlikePost(userID, postID uint64) error
}
