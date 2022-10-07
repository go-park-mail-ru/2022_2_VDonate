package authAPI

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type API interface {
	Login(login, password string) (string, error)
	Auth(sessionID string) (bool, error)
	SignUp(user *models.UserDB) (string, error)
	Logout(sessionID string) (bool, error)
}
