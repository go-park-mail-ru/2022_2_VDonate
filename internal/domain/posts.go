package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type PostsUseCase interface {
	GetPostByID(postID, userID uint64) (models.Post, error)
	GetPostsByFilter(userID, authorID uint64) ([]models.Post, error)
	Create(post models.Post, userID uint64) (uint64, string, error)
	Update(post models.Post, postID uint64) (models.Post, error)
	DeleteByID(postID uint64) error
	GetLikesByPostID(postID uint64) ([]models.Like, error)
	GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error)
	LikePost(userID, postID uint64) error
	UnlikePost(userID, postID uint64) error
	GetLikesNum(postID uint64) (uint64, error)
	IsPostLiked(userID, postID uint64) bool
	CreateTags(tagNames []string, postID uint64) error
	GetTagsByPostID(postID uint64) ([]models.Tag, error)
	DeleteTagDeps(postID uint64) error
	UpdateTags(tagNames []string, postID uint64) error
	ConvertTagsToStrSlice(tags []models.Tag) []string
}
