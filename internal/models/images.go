package models

type ResponseImagePosts struct {
	PostID  uint64 `json:"postID" validate:"required" example:"25"`
	ImgPath string `json:"imgPath" validate:"required" example:"/path/to/image.jpeg"`
}

type ResponseImageSubscription struct {
	SubscriptionID uint64 `json:"subscriptionID" validate:"required" example:"25"`
	ImgPath        string `json:"imgPath" validate:"required" example:"/path/to/image.jpeg"`
}
