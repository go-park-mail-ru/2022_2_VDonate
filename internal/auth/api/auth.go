package auth

import (
	auth_api "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth"
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models/user"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository"
)

type Auth struct {
	API      *auth_api.API
	userRepo *user_repo.Repo
}

func (api *Auth) Login(login string, password string) (sessionId string, err error) {
	return "", nil
}
func (api *Auth) Logout(sessionId string) error {
	return nil
}
func (api *Auth) SignUp(user *model.User) (sessionId string, err error) {
	return "", nil
}
func (api *Auth) GetUnauthorizedSession() (sessionId string, err error) {
	return "", nil
}
func (api *Auth) IsSession(sessionId string) bool {
	return false
}
func (api *Auth) IsAuthSession(sessionId string) bool {
	return false
}
