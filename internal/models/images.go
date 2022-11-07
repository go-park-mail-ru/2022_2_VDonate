package models

type ResponseImage struct {
	ImgPath string `json:"imgPath" validate:"required" example:"/path/to/image.jpeg"`
}
