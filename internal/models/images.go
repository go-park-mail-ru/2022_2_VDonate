package models

type ResponseImagePosts struct {
	PostID          uint64 `json:"postID" validate:"required" example:"25"`
	ContentTemplate string `json:"contentTemplate" validate:"required" example:"<img src=\"\">"`
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
