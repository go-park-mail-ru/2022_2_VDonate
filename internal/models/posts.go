package models

type Post struct {
	ID     uint64 `json:"id" form:"post_id" db:"post_id" example:"1"`
	UserID uint64 `json:"user_id" form:"user_id" db:"user_id" example:"1"`
	Img    string `json:"img" form:"img" db:"img" validate:"required" example:"path/to/image.jpeg"`
	Title  string `json:"title" form:"title" db:"title" validate:"required" example:"some title"`
	Text   string `json:"text" form:"text" db:"text" validate:"required" example:"some text"`
}

type Like struct {
	UserID        uint64 `json:"user_id" db:"user_id"`
	PostID        uint64 `json:"post_id" db:"post_id"`
}

func (p Post) GetID() uint64 {
	return p.ID
}
