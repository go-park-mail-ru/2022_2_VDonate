package models

type User struct {
	ID       uint64 `json:"id" db:"id" form:"id"`
	Username string `json:"username" db:"username" form:"username"`
	Email    string `json:"email" db:"email" form:"email"`
	Avatar   string `json:"avatar,omitempty" db:"avatar" form:"avatar"`
	Password string `json:"password" db:"password" form:"password"`
	IsAuthor bool   `json:"is_author" db:"is_author" form:"is_author"`
	About    string `json:"about,omitempty" db:"about" form:"about"`
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
	Username string `json:"username"`
	Password string `json:"password"`
}

type Author struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar,omitempty"`
	IsAuthor bool   `json:"is_author"`
	About    string `json:"about,omitempty"`
}

type NotAuthor struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar,omitempty"`
	IsAuthor bool   `json:"is_author"`
}

func ToAuthor(u *User) *Author {
	return &Author{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Avatar:   u.Avatar,
		IsAuthor: true,
		About:    u.About,
	}
}

func ToNonAuthor(u *User) *NotAuthor {
	return &NotAuthor{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Avatar:   u.Avatar,
		IsAuthor: false,
	}
}
