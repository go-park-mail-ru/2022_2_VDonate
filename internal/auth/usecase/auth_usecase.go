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

type usecase struct {
	authRepo  Repository
	usersRepo userAPI.Repository
}

func New(authRepo Repository, usersRepo userAPI.Repository) UseCase {
	return &usecase{authRepo: authRepo, usersRepo: usersRepo}
}

func (u *usecase) Login(login, password string) (string, error) {
	user, err := u.usersRepo.GetByUsername(login)
	if err != nil {
		user, err = u.usersRepo.GetByEmail(login)
		if err != nil {
			return "", err
		}
	}
	if password != user.Password {
		return "", errors.New("passwords not the same")
	}
	s, err := u.authRepo.CreateSession(models.CreateCookie(user.ID))
	if err != nil {
		return "", err
	}

	return s.Value, nil
}

func (u *usecase) Auth(sessionID string) (bool, error) {
	_, err := u.authRepo.GetBySessionID(sessionID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *usecase) SignUp(user *models.User) (string, error) {
	_, err := u.usersRepo.GetByUsername(user.Username)
	if err == nil {
		return "", err
	}
	if _, err = u.usersRepo.GetByEmail(user.Email); err == nil {
		return "", err
	}
	if user, err = u.usersRepo.Create(user); err != nil {
		return "", err
	}
	s, err := u.authRepo.CreateSession(models.CreateCookie(user.ID))
	if err != nil {
		return "", err
	}
	return s.Value, nil
}

func (u *usecase) Logout(sessionID string) (bool, error) {
	if err := u.authRepo.DeleteBySessionID(sessionID); err != nil {
		return false, err
	}
	return true, nil
}

func (u *usecase) IsSameSession(sessionID string, userID uint64) bool {
	user, err := u.usersRepo.GetBySessionID(sessionID)
	if err != nil {
		return false
	}
	return user.ID == userID
}
