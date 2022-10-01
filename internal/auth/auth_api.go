package auth_api

import (
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models/user"
)

type API interface {
	Login(login string, password string) (sessionId string, err error)
	Logout(sessionId string) error
	SignUp(user *model.User) (sessionId string, err error)
	GetUnauthorizedSession() (sessionId string, err error)
	IsSession(sessionId string) bool
	IsAuthSession(sessionId string) bool
}
