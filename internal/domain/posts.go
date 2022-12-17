package domain

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type PostsMicroservice interface {
	GetAllByUserID(authorID uint64) ([]models.Post, error)
	GetPostByID(postID uint64) (models.Post, error)
	Create(post models.Post) (models.Post, error)
	Update(post models.Post) error
	DeleteByID(postID uint64) error
	GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error)
	GetAllLikesByPostID(postID uint64) ([]models.Like, error)
	CreateLike(userID, postID uint64) error
	DeleteLikeByID(userID, postID uint64) error
	GetPostsBySubscriptions(userID uint64) ([]models.Post, error)
	CreateTag(tagName string) (uint64, error)
	CreateDepTag(postID, tagID uint64) error
	DeleteDepTag(tagDep models.TagDep) error
	GetTagById(tagID uint64) (models.Tag, error)
	GetTagDepsByPostId(postID uint64) ([]models.TagDep, error)
	GetTagByName(tagName string) (models.Tag, error)
	CreateComment(comment models.Comment) (models.Comment, error)
	GetCommentByID(commentID uint64) (models.Comment, error)
	GetCommentsByPostID(postID uint64) ([]models.Comment, error)
	UpdateComment(comment models.Comment) error
	DeleteCommentByID(commentID uint64) error
}

type PostsUseCase interface {
	GetPostByID(postID, userID uint64) (models.Post, error)
	GetPostsByFilter(userID, authorID uint64) ([]models.Post, error)
	Create(post models.Post, userID uint64) (models.Post, error)
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
	CreateComment(comment models.Comment) (models.Comment, error)
	GetCommentsByPostID(postID uint64) ([]models.Comment, error)
	GetCommentByID(commentID uint64) (models.Comment, error)
	UpdateComment(commentID uint64, commentMsg string) (models.Comment, error)
	DeleteComment(commentID uint64) error
}
