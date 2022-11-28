package models

import (
	"time"
)

type Post struct {
	ID          uint64    `json:"postID" form:"postID" db:"post_id" example:"1"`
	UserID      uint64    `json:"userID" form:"userID" db:"user_id" example:"1"`
	Img         string    `json:"img" form:"img" db:"img" validate:"required" example:"path/to/image.jpeg"`
	Title       string    `json:"title" form:"title" db:"title" validate:"required" example:"some title"`
	Text        string    `json:"text" form:"text" db:"text" validate:"required" example:"some text"`
	Tier        uint64    `json:"tier" form:"tier" db:"tier" example:"5"`
	IsAllowed   bool      `json:"isAllowed" example:"true"`
	DateCreated time.Time `json:"dateCreated" form:"date_created" db:"date_created" example:"2022-11-11"`
	Tags        []string  `json:"tags" form:"tags" db:"tag_name"`

	Author   ResponseImageUsers `json:"author" validate:"required"`
	LikesNum uint64             `json:"likesNum" validate:"required" example:"5"`
	IsLiked  bool               `json:"isLiked" validate:"required" example:"true"`
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
