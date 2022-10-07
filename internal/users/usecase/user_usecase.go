package userAPI

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/session/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository"
)

type UseCase interface {
	GetByUsername(username string) (*models.UserDB, error)
	GetByEmail(email string) (*models.UserDB, error)
	GetByID(id uint) (*models.UserDB, error)
	Create(user *models.UserDB) (*models.UserDB, error)
	DeleteByID(id uint) error
	DeleteByUsername(username string) error
	DeleteByEmail(email string) error
	CheckIDAndPassword(id uint, password string) bool
	IsExistUsernameAndEmail(username, email string) bool
}

type API struct {
	userRepo    userRepository.API
	sessionRepo sessionRepository.API
}

func New(userRepo userRepository.API, sessionRepo sessionRepository.API) UseCase {
	return &API{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (a *API) GetByID(id uint) (*models.UserDB, error) {
	return a.userRepo.FindByID(id)
}

func (a *API) GetByUsername(username string) (*models.UserDB, error) {
	return a.userRepo.FindByUsername(username)
}

func (a *API) GetByEmail(email string) (*models.UserDB, error) {
	return a.userRepo.FindByEmail(email)
}

func (a *API) Create(user *models.UserDB) (*models.UserDB, error) {
	return a.userRepo.Create(user)
}

func (a *API) DeleteByID(id uint) error {
	return a.userRepo.DeleteByID(id)
}
func (a *API) DeleteByUsername(username string) error {
	user, err := a.GetByUsername(username)
	if err != nil {
		return err
	}
	return a.DeleteByID(user.ID)
}
func (a *API) DeleteByEmail(email string) error {
	user, err := a.GetByEmail(email)
	if err != nil {
		return err
	}
	return a.DeleteByID(user.ID)
}

func (a *API) CheckIDAndPassword(id uint, password string) bool {
	user, err := a.GetByID(id)
	if err != nil {
		return false
	}
	return user.Password == password
}

func (a *API) IsExistUsernameAndEmail(username, email string) bool {
	_, err := a.GetByUsername(username)
	if err == nil {
		if _, err = a.GetByEmail(email); err == nil {
			return true
		}
	}
	return false
}
