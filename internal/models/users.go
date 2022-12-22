package models

type User struct {
	ID       uint64 `json:"id" db:"id" form:"id" example:"1"`
	Username string `json:"username" db:"username" form:"username" validate:"required" example:"admin"`
	Email    string `json:"email" db:"email" form:"email" validate:"required,email" example:"admin@mail.ru"`
	Avatar   string `json:"avatar" db:"avatar" form:"avatar" example:"filename.jpeg"`
	Password string `json:"password" db:"password" form:"password" validate:"required" example:"*****"`
	IsAuthor bool   `json:"isAuthor" db:"is_author" form:"isAuthor" validate:"required,boolean" example:"true"`
	Balance  uint64 `json:"balance" db:"balance" form:"balance" validate:"required" example:"1000"`
	About    string `json:"about" db:"about" form:"about" example:"it's info about myself"`

	CountSubscriptions     uint64 `json:"countSubscriptions" example:"25"`
	CountSubscribers       uint64 `json:"countSubscribers" example:"120"`
	CountPosts             uint64 `json:"countPosts" example:"12"`
	CountSubscribersMounth uint64 `json:"countSubscribersMounth,omitempty" example:"12"`
	CountProfitMounth      uint64 `json:"countProfitMounth,omitempty" example:"12"`
}

func (v User) GetID() uint64 {
	return v.ID
}

type AuthUser struct {
	Username string `json:"username" validate:"required" example:"admin"`
	Password string `json:"password" validate:"required" example:"*****"`
}

type UserID struct {
	ID uint64 `json:"id" validate:"required" example:"12"`
}

type Author struct {
	ID       uint64 `json:"id" db:"id" form:"id" example:"1"`
	Username string `json:"username" validate:"required" example:"admin"`
	Email    string `json:"email" validate:"required" example:"admin@mail.ru"`
	Avatar   string `json:"avatar" example:"filename.jpeg"`
	IsAuthor bool   `json:"isAuthor" validate:"required" example:"true"`
	Balance  uint64 `json:"balance" validate:"required" example:"1000"`
	About    string `json:"about" example:"it's info about myself"`

	CountSubscriptions     uint64 `json:"countSubscriptions" example:"25"`
	CountSubscribers       uint64 `json:"countSubscribers" example:"120"`
	CountPosts             uint64 `json:"countPosts" example:"12"`
	CountSubscribersMounth uint64 `json:"countSubscribersMounth" example:"12"`
	CountProfitMounth      uint64 `json:"countProfitMounth" example:"12"`
}

type NotAuthor struct {
	ID       uint64 `json:"id" db:"id" form:"id" example:"1"`
	Username string `json:"username" validate:"required" example:"admin"`
	Email    string `json:"email" validate:"required" example:"admin@mail.ru"`
	Avatar   string `json:"avatar" example:"filename.jpeg"`
	IsAuthor bool   `json:"isAuthor" validate:"required" example:"true"`

	CountSubscriptions uint64 `json:"countSubscriptions" example:"120"`
}

func ToAuthor(u User) Author {
	return Author{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Avatar:   u.Avatar,
		IsAuthor: true,
		Balance:  u.Balance,
		About:    u.About,

		CountSubscriptions:     u.CountSubscriptions,
		CountSubscribers:       u.CountSubscribers,
		CountPosts:             u.CountPosts,
		CountSubscribersMounth: u.CountSubscribersMounth,
		CountProfitMounth:      u.CountProfitMounth,
	}
}

func ToNonAuthor(u User) NotAuthor {
	return NotAuthor{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Avatar:   u.Avatar,
		IsAuthor: false,

		CountSubscriptions: u.CountSubscriptions,
	}
}
