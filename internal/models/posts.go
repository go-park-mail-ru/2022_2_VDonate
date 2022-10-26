package models

type Post struct {
	ID     uint64 `json:"id" form:"post_id" db:"post_id"`
	UserID uint64 `json:"user_id" form:"user_id" db:"user_id"`
	Img    string `json:"img" form:"img" db:"img"`
	Title  string `json:"title" form:"title" db:"title"`
	Text   string `json:"text" form:"text" db:"text"`
}

func (p Post) GetID() uint64 {
	return p.ID
}
