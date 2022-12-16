package auth

import (
	"time"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/google/uuid"
)

type (
	cookieCreator func(id uint64) models.Cookie
	hashCreator   func(password string) (string, error)
)

type usecase struct {
	authMicroservice  domain.AuthMicroservice
	usersMicroservice domain.UsersMicroservice

	// cookieCreator is func for creation cookie,
	// so you can test random sessionID
	cookieCreator cookieCreator

	// hashCreator is func for creation hash from password,
	// so you can test it
	hashCreator hashCreator
}

func New(authMicroservice domain.AuthMicroservice, usersMicroservice domain.UsersMicroservice) domain.AuthUseCase {
	return &usecase{
		authMicroservice:  authMicroservice,
		usersMicroservice: usersMicroservice,
		cookieCreator:     createCookie,
		hashCreator:       utils.HashPassword,
	}
}

func WithCookieCreator(authMicroservice domain.AuthMicroservice, usersMicroservice domain.UsersMicroservice, cookieCreator cookieCreator) domain.AuthUseCase {
	return &usecase{
		authMicroservice:  authMicroservice,
		usersMicroservice: usersMicroservice,
		cookieCreator:     cookieCreator,
	}
}

func WithCreators(authMicroservice domain.AuthMicroservice, usersMicroservice domain.UsersMicroservice, cookieCreator cookieCreator, hashCreator hashCreator) domain.AuthUseCase {
	return &usecase{
		authMicroservice:  authMicroservice,
		usersMicroservice: usersMicroservice,
		cookieCreator:     cookieCreator,
		hashCreator:       hashCreator,
	}
}

func createCookie(id uint64) models.Cookie {
	return models.Cookie{
		UserID:  id,
		Value:   uuid.New().String(),
		Expires: time.Now().AddDate(0, 1, 0),
	}
}

func (u usecase) Login(login, password string) (string, error) {
	user, err := u.usersMicroservice.GetByUsername(login)
	if err != nil {
		user, err = u.usersMicroservice.GetByEmail(login)
		if err != nil {
			return "", domain.ErrUsernameOrEmailNotExist
		}
	}

	matchPassword := utils.CheckHashPassword(password, user.Password)

	if !matchPassword {
		return "", domain.ErrPasswordsNotEqual
	}

	s, err := u.authMicroservice.CreateSession(user.ID)
	if err != nil {
		return "", err
	}

	return s, nil
}

func (u usecase) Auth(sessionID string) (bool, error) {
	_, err := u.authMicroservice.GetBySessionID(sessionID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u usecase) SignUp(user models.User) (string, error) {
	_, err := u.usersMicroservice.GetByUsername(user.Username)
	if err == nil {
		return "", domain.ErrUsernameExist
	}

	if _, err = u.usersMicroservice.GetByEmail(user.Email); err == nil {
		return "", domain.ErrEmailExist
	}

	if user.Password, err = u.hashCreator(user.Password); err != nil {
		return "", domain.ErrInternal
	}

	if user.ID, err = u.usersMicroservice.Create(user); err != nil {
		return "", err
	}

	s, err := u.authMicroservice.CreateSession(user.ID)
	if err != nil {
		return "", err
	}

	return s, nil
}

func (u usecase) Logout(sessionID string) (bool, error) {
	if err := u.authMicroservice.DeleteBySessionID(sessionID); err != nil {
		return false, err
	}

	return true, nil
}

func (u usecase) IsSameSession(sessionID string, userID uint64) bool {
	user, err := u.usersMicroservice.GetBySessionID(sessionID)
	if err != nil {
		return false
	}

	return user.ID == userID
}
