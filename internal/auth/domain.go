package authDomain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type UseCase interface {
	Login(login, password string) (string, error)
	Auth(sessionID string) (bool, error)
	SignUp(user *models.User) (string, error)
	Logout(sessionID string) (bool, error)
	IsSameSession(sessionID string, userID uint64) bool
}

type Repository interface {
	GetBySessionID(sessionID string) (*models.Cookie, error)
	GetByUserID(id uint64) (*models.Cookie, error)
	GetByUsername(username string) (*models.Cookie, error)
	CreateSession(cookie models.Cookie) (*models.Cookie, error)
	DeleteBySessionID(sessionID string) error
	DeleteByUserID(id uint64) error
	Close() error
}
