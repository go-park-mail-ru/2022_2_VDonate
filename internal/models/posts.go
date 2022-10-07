package models

type PostDB struct {
	ID     uint   `json:"id" db:"post_id"`
	UserID uint   `json:"user_id" db:"user_id"`
	Title  string `json:"title" db:"title"`
	Text   string `json:"text" db:"text"`
}
