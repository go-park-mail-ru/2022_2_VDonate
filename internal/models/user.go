package models

type User struct {
	ID        uint64 `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	FirstName string `json:"first_name,omitempty" db:"first_name"`
	LastName  string `json:"last_name,omitempty" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Avatar    string `json:"avatar,omitempty" db:"avatar"`
	Password  string `json:"password" db:"password"`
	Phone     string `json:"phone,omitempty" db:"phone"`
	IsAuthor  bool   `json:"is_author" db:"is_author"`
	About     string `json:"about,omitempty" db:"about"`
}

func (u User) GetID() uint64 {
	return u.ID
}

type model interface {
	*User | *Post
	GetID() uint64
}

func Contains[M model](slice []M, ID uint64) bool {
	for _, el := range slice {
		if el.GetID() == ID {
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
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar,omitempty"`
	Phone     string `json:"phone,omitempty"`
	IsAuthor  bool   `json:"is_author"`
	About     string `json:"about,omitempty"`
}

type NotAuthor struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar,omitempty"`
	Phone     string `json:"phone,omitempty"`
	IsAuthor  bool   `json:"is_author"`
}

func ToAuthor(u *User) *Author {
	return &Author{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Avatar:    u.Avatar,
		Phone:     u.Phone,
		IsAuthor:  true,
		About:     u.About,
	}
}

func ToNonAuthor(u *User) *NotAuthor {
	return &NotAuthor{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Avatar:    u.Avatar,
		Phone:     u.Phone,
		IsAuthor:  false,
	}
}
