package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type PostsUseCase interface {
	GetPostByID(postID, userID uint64) (models.Post, error)
	GetPostsByFilter(filter string, userID uint64) ([]models.Post, error)
	Create(post models.Post, userID uint64) (uint64, error)
	Update(post models.Post, postID uint64) error
	DeleteByID(postID uint64) error
	GetLikesByPostID(postID uint64) ([]models.Like, error)
	GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error)
	LikePost(userID, postID uint64) error
	UnlikePost(userID, postID uint64) error
	GetLikesNum(postID uint64) (uint64, error)
	IsPostLiked(userID, postID uint64) bool
}
