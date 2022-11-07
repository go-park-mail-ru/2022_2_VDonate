package models

import "mime/multipart"

type EmptyStruct struct{}

type Error struct {
	Message string
}

type PostMpfd struct {
	Title string `json:"title" form:"title" db:"title" validate:"required" example:"some title"`
	Text  string `json:"text" form:"text" db:"text" validate:"required" example:"some text"`
	File  multipart.File
}

type UserMpfd struct {
	Username string `json:"username" db:"username" form:"username" validate:"required" example:"admin"`
	Password string `json:"password" db:"password" form:"password" validate:"required" example:"*****"`
	Email    string `json:"email" db:"email" form:"email" validate:"required" example:"admin@mail.ru"`
}

type AuthorSubscriptionMpfd struct {
	ID    uint64 `json:"id" form:"id" db:"id" example:"1"`
	Img   string `json:"img" db:"img" example:"filename.jpeg"`
	Tier  uint64 `json:"tier" form:"tier" db:"tier" validate:"required" example:"15"`
	Title string `json:"title" form:"title" db:"title" validate:"required" example:"some title"`
	Text  string `json:"text" form:"text" db:"text" validate:"required" example:"some text"`
	Price uint64 `json:"price" form:"price" db:"price" validate:"required" example:"2999"`

	Author Author `json:"author" form:"author" validate:"required"`
}

type SubscriptionMpfd struct {
	AuthorID             uint64 `json:"authorID" validate:"required" example:"12"`
	AuthorSubscriptionID uint64 `json:"authorSubscriptionID" validate:"required" example:"13"`
}
