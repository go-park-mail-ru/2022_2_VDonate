package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type AuthUseCase interface {
	Login(login, password string) (string, error)
	Auth(sessionID string) (bool, error)
	SignUp(user models.User) (string, error)
	Logout(sessionID string) (bool, error)
	IsSameSession(sessionID string, userID uint64) bool
}
