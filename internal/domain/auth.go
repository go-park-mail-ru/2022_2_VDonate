package domain

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type AuthUseCase interface {
	Login(login, password string) (string, error)
	Auth(sessionID string) (bool, error)
	SignUp(user models.User) (string, error)
	Logout(sessionID string) (bool, error)
	IsSameSession(sessionID string, userID uint64) bool
}

type AuthServiceManager interface {
	CreateSession(userID uint64) (string, error)
	DeleteBySessionID(sessionID string) error
	GetBySessionID(sessionID string) (*protobuf.Session, error)
}
