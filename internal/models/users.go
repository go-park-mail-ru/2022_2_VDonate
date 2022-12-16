package models

type User struct {
	ID       uint64 `json:"id" db:"id" form:"id" example:"1"`
	Username string `json:"username" db:"username" form:"username" validate:"required" example:"admin"`
	Email    string `json:"email" db:"email" form:"email" validate:"required" example:"admin@mail.ru"`
	Avatar   string `json:"avatar" db:"avatar" form:"avatar" example:"filename.jpeg"`
	Password string `json:"password" db:"password" form:"password" validate:"required" example:"*****"`
	IsAuthor bool   `json:"isAuthor" db:"is_author" form:"isAuthor" validate:"required" example:"true"`
	About    string `json:"about" db:"about" form:"about" example:"it's info about myself"`

	CountSubscriptions uint64 `json:"countSubscriptions" example:"25"`
	CountSubscribers   uint64 `json:"countSubscribers" example:"120"`
}

func (u User) GetID() uint64 {
	return u.ID
}

type AuthUser struct {
	Username string `json:"username" validate:"required" example:"admin"`
	Password string `json:"password" validate:"required" example:"*****"`
}

type UserID struct {
	ID uint64 `json:"id" validate:"required" example:"12"`
}

type Author struct {
	Username string `json:"username" validate:"required" example:"admin"`
	Email    string `json:"email" validate:"required" example:"admin@mail.ru"`
	Avatar   string `json:"avatar" example:"filename.jpeg"`
	IsAuthor bool   `json:"isAuthor" validate:"required" example:"true"`
	About    string `json:"about" example:"it's info about myself"`

	CountSubscriptions uint64 `json:"countSubscriptions" example:"25"`
	CountSubscribers   uint64 `json:"countSubscribers" example:"120"`
}

type NotAuthor struct {
	Username string `json:"username" validate:"required" example:"admin"`
	Email    string `json:"email" validate:"required" example:"admin@mail.ru"`
	Avatar   string `json:"avatar" example:"filename.jpeg"`
	IsAuthor bool   `json:"isAuthor" validate:"required" example:"true"`

	CountSubscriptions uint64 `json:"countSubscriptions" example:"120"`
}

func ToAuthor(u User) Author {
	return Author{
		Username: u.Username,
		Email:    u.Email,
		Avatar:   u.Avatar,
		IsAuthor: true,
		About:    u.About,

		CountSubscriptions: u.CountSubscriptions,
		CountSubscribers:   u.CountSubscribers,
	}
}

func ToNonAuthor(u User) NotAuthor {
	return NotAuthor{
		Username: u.Username,
		Email:    u.Email,
		Avatar:   u.Avatar,
		IsAuthor: false,

		CountSubscriptions: u.CountSubscriptions,
	}
}
