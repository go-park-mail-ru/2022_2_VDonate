package models

type Post struct {
	ID     uint64 `json:"postID" form:"postID" db:"post_id" example:"1"`
	UserID uint64 `json:"userID" form:"userID" db:"user_id" example:"1"`
	Img    string `json:"img" form:"img" db:"img" validate:"required" example:"path/to/image.jpeg"`
	Title  string `json:"title" form:"title" db:"title" validate:"required" example:"some title"`
	Text   string `json:"text" form:"text" db:"text" validate:"required" example:"some text"`

	Author   ResponseImageUsers `json:"author" validate:"required"`
	LikesNum uint64             `json:"likesNum" validate:"required" example:"5"`
	IsLiked  bool               `json:"isLiked" validate:"required" example:"true"`
}

type Like struct {
	UserID uint64 `json:"userID" db:"user_id" validate:"required" example:"100"`
	PostID uint64 `json:"postID" db:"post_id" validate:"required" example:"222"`
}

func (p Post) GetID() uint64 {
	return p.ID
}
