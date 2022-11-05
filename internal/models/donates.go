package models

type Donate struct {
	ID       uint64 `json:"id" db:"id"`
	UserID   uint64 `json:"user_id" db:"user_id"`
	AuthorID uint64 `json:"author_id" db:"author_id"`
	Price    uint64 `json:"price" db:"price"`
}
