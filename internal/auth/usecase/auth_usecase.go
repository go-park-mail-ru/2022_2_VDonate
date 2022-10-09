package auth

import (
	"errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	userAPI "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
)

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
	CreateSession(cookie *models.Cookie) (*models.Cookie, error)
	DeleteBySessionID(sessionID string) error
	DeleteByUserID(id uint64) error
	Close() error
}

type api struct {
	authRepo  Repository
	usersRepo userAPI.Repository
}

func New(authRepo Repository, usersRepo userAPI.Repository) UseCase {
	return &api{authRepo: authRepo, usersRepo: usersRepo}
}

func (a *api) Login(login, password string) (string, error) {
	user, err := a.usersRepo.FindByUsername(login)
	if err != nil {
		user, err = a.usersRepo.FindByEmail(login)
		if err != nil {
			return "", err
		}
	}
	if password != user.Password {
		return "", errors.New("passwords not the same")
	}
	s, err := a.authRepo.CreateSession(models.CreateCookie(user.ID))
	if err != nil {
		return "", err
	}

	return s.Value, nil
}

func (a *api) Auth(sessionID string) (bool, error) {
	_, err := a.authRepo.GetBySessionID(sessionID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *api) SignUp(user *models.User) (string, error) {
	_, err := a.usersRepo.FindByUsername(user.Username)
	if err == nil {
		return "", err
	}
	if _, err = a.usersRepo.FindByEmail(user.Email); err == nil {
		return "", err
	}
	if user, err = a.usersRepo.Create(user); err != nil {
		return "", err
	}
	s, err := a.authRepo.CreateSession(models.CreateCookie(user.ID))
	if err != nil {
		return "", err
	}
	return s.Value, nil
}

func (a *api) Logout(sessionID string) (bool, error) {
	if err := a.authRepo.DeleteBySessionID(sessionID); err != nil {
		return false, err
	}
	return true, nil
}

func (a *api) IsSameSession(sessionID string, userID uint64) bool {
	user, err := a.usersRepo.FindBySessionID(sessionID)
	if err != nil {
		return false
	}
	return user.ID == userID
}
