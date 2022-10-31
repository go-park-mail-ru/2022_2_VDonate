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
	Tier  uint64 `json:"tier" form:"tier" db:"tier" validate:"required" example:"15"`
	Title string `json:"title" form:"title" db:"title" validate:"required" example:"some title"`
	Text  string `json:"text" form:"text" db:"text" validate:"required" example:"some text"`
	Price uint64 `json:"price" form:"price" db:"price" validate:"required" example:"2999"`
}
