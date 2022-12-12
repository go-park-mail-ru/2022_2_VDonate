package models

import "time"

type ResponseImagePosts struct {
	PostID      uint64    `json:"postID" validate:"required" example:"25"`
	Content     string    `json:"content" validate:"required" example:"<img src=\"\">"`
	DateCreated time.Time `json:"dateCreated" validate:"required" example:"2020-01-01 00:00:00"`
}

type ResponseImageSubscription struct {
	SubscriptionID uint64 `json:"subscriptionID" validate:"required" example:"25"`
	ImgPath        string `json:"imgPath" validate:"required" example:"/path/to/image.jpeg"`

	AuthorName string `json:"authorName" validate:"required" example:"leo"`
	AuthorImg  string `json:"authorImg" validate:"required" example:"path/to/author/img"`
}

type ResponseImageUsers struct {
	UserID   uint64 `json:"userID" validate:"required" example:"25"`
	Username string `json:"username" validate:"required" example:"leo"`
	ImgPath  string `json:"imgPath" validate:"required" example:"/path/to/image.jpeg"`
}
