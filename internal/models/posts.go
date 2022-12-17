package models

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Post struct {
	ID          uint64    `json:"postID" form:"postID" db:"post_id" example:"1"`
	UserID      uint64    `json:"userID" form:"userID" db:"user_id" example:"1"`
	Content     string    `json:"content" form:"content" db:"content" validate:"required" example:"Hello <img src=...>"`
	Tier        uint64    `json:"tier" form:"tier" db:"tier" example:"5"`
	IsAllowed   bool      `json:"isAllowed" example:"true"`
	DateCreated time.Time `json:"dateCreated" form:"date_created" db:"date_created" example:"2022-11-11"`
	Tags        []string  `json:"tags" form:"tags" db:"tag_name"`

	Author          ResponseImageUsers `json:"author" validate:"required"`
	LikesNum        uint64             `json:"likesNum" validate:"required" example:"5"`
	IsLiked         bool               `json:"isLiked" validate:"required" example:"true"`

	CommentsNum     uint64             `json:"commentsNum" validate:"required" example:"5"`
}

type Comment struct {
	ID          uint64                 `json:"id" form:"id" db:"id" validate:"required" example:"1"`
	PostID      uint64                 `json:"postID" form:"postID" db:"post_id" validate:"required" example:"1"`
	AuthorID    uint64                 `json:"authorID" form:"authorID" db:"author_id" validate:"required" example:"1"`
	UserID      uint64                 `json:"userID" form:"userID" db:"user_id" validate:"required" example:"1"`
	DateCreated *timestamppb.Timestamp `json:"dateCreated" form:"date_created" db:"date_created" example:"2022-11-11"`
	Content     string                 `json:"content" form:"content" db:"content" validate:"required" example:"Looks great!"`
}

type Tag struct {
	ID      uint64 `json:"id" form:"id" db:"id" validate:"required" example:"1"`
	TagName string `json:"tagName" form:"tagName" db:"tag_name" validate:"required" example:"sport"`
}

type TagDep struct {
	PostID uint64 `json:"postId" form:"postId" db:"post_id"`
	TagID  uint64 `json:"tagId" form:"tagId" db:"tag_id"`
}

type Like struct {
	UserID uint64 `json:"userID" db:"user_id" validate:"required" example:"100"`
	PostID uint64 `json:"postID" db:"post_id" validate:"required" example:"222"`
}

func (p Post) GetID() uint64 {
	return p.ID
}
