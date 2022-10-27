package models

type ResponseImage struct {
	ImgPath string `json:"img_path" validate:"required" example:"/path/to/image.jpeg"`
}
