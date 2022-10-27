package models

type User struct {
	ID       uint64 `json:"id" db:"id" form:"id" example:"1"`
	Username string `json:"username" db:"username" form:"username" validate:"required" example:"admin"`
	Email    string `json:"email" db:"email" form:"email" validate:"required" example:"admin@mail.ru"`
	Avatar   string `json:"avatar,omitempty" db:"avatar" form:"avatar" example:"filename.jpeg"`
	Password string `json:"password" db:"password" form:"password" validate:"required" example:"*****"`
	IsAuthor bool   `json:"is_author" db:"is_author" form:"is_author" validate:"required" example:"true"`
	About    string `json:"about,omitempty" db:"about" form:"about" example:"it's info about myself"`
}

func (u User) GetID() uint64 {
	return u.ID
}

type model interface {
	*User | *Post
	GetID() uint64
}

func Contains[M model](slice []M, id uint64) bool {
	for _, el := range slice {
		if el.GetID() == id {
			return true
		}
	}

	return false
}

type AuthUser struct {
	Username string `json:"username" validate:"required" example:"admin"`
	Password string `json:"password" validate:"required" example:"*****"`
}

type Author struct {
	Username string `json:"username" validate:"required" example:"admin"`
	Email    string `json:"email" validate:"required" example:"admin@mail.ru"`
	Avatar   string `json:"avatar,omitempty" example:"filename.jpeg"`
	IsAuthor bool   `json:"is_author" validate:"required" example:"true"`
	About    string `json:"about" example:"it's info about myself"`
}

type NotAuthor struct {
	Username string `json:"username" validate:"required" example:"admin"`
	Email    string `json:"email" validate:"required" example:"admin@mail.ru"`
	Avatar   string `json:"avatar,omitempty" example:"filename.jpeg"`
	IsAuthor bool   `json:"is_author" validate:"required" example:"true"`
}

func ToAuthor(u *User) *Author {
	return &Author{
		Username: u.Username,
		Email:    u.Email,
		Avatar:   u.Avatar,
		IsAuthor: true,
		About:    u.About,
	}
}

func ToNonAuthor(u *User) *NotAuthor {
	return &NotAuthor{
		Username: u.Username,
		Email:    u.Email,
		Avatar:   u.Avatar,
		IsAuthor: false,
	}
}
