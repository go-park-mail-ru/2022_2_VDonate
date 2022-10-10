package models

type PostDB struct {
	ID     uint64 `json:"id" db:"post_id"`
	UserID uint64 `json:"user_id" db:"user_id"`
	Img    string `json:"img" db:"img"`
	Title  string `json:"title" db:"title"`
	Text   string `json:"text" db:"text"`
}
