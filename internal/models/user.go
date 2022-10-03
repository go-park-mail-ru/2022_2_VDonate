package models

type UserDB struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar,omitempty"`
	Password  string `json:"password"`
	Phone     uint   `json:"phone,omitempty"`
	IsAuthor  bool   `json:"is_author"`
	About     string `json:"about,omitempty"`
}

type Author struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar,omitempty"`
	Password  string `json:"password"`
	Phone     uint   `json:"phone,omitempty"`
	IsAuthor  bool   `json:"is_author"`
	About     string `json:"about,omitempty"`
}

type NotAuthor struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar,omitempty"`
	Password  string `json:"password"`
	Phone     uint   `json:"phone,omitempty"`
	IsAuthor  bool   `json:"is_author"`
}
