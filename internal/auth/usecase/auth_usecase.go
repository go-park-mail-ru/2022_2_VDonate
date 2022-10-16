package auth

import (
	"errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/google/uuid"
	"time"
)

type cookieCreator func(id uint64) models.Cookie

type usecase struct {
	authRepo  domain.AuthRepository
	usersRepo domain.UsersRepository

	// cookieCreator is func for creation cookie,
	// so you can test random sessionID
	cookieCreator cookieCreator
}

func New(authRepo domain.AuthRepository, usersRepo domain.UsersRepository) domain.AuthUseCase {
	return &usecase{authRepo: authRepo, usersRepo: usersRepo, cookieCreator: createCookie}
}

func createCookie(id uint64) models.Cookie {
	return models.Cookie{
		UserID:  id,
		Value:   uuid.New().String(),
		Expires: time.Now().AddDate(0, 1, 0),
	}
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
	s, err := u.authRepo.CreateSession(u.cookieCreator(user.ID))
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
		return "", errors.New("username is already exist")
	}
	if _, err = u.usersRepo.GetByEmail(user.Email); err == nil {
		return "", errors.New("email is already exist")
	}
	if user, err = u.usersRepo.Create(user); err != nil {
		return "", err
	}
	s, err := u.authRepo.CreateSession(u.cookieCreator(user.ID))
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
