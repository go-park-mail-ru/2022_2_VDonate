package users

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jinzhu/copier"
)

type UseCase interface {
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByID(id uint64) (*models.User, error)
	GetBySessionID(sessionID string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(id uint64, user *models.User) (*models.User, error)
	DeleteByID(id uint64) error
	DeleteByUsername(username string) error
	DeleteByEmail(email string) error
	CheckIDAndPassword(id uint64, password string) bool
	IsExistUsernameAndEmail(username, email string) bool
}

type Repository interface {
	Create(u *models.User) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByID(id uint64) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindBySessionID(sessionID string) (*models.User, error)
	Update(u *models.User) (*models.User, error)
	DeleteByID(id uint64) error
	Close() error
}

type api struct {
	usersRepo Repository
}

func New(usersRepo Repository) UseCase {
	return &api{
		usersRepo: usersRepo,
	}
}

func (a *api) GetByID(id uint64) (*models.User, error) {
	return a.usersRepo.FindByID(id)
}

func (a *api) GetByUsername(username string) (*models.User, error) {
	return a.usersRepo.FindByUsername(username)
}

func (a *api) GetByEmail(email string) (*models.User, error) {
	return a.usersRepo.FindByEmail(email)
}

func (a *api) GetBySessionID(sessionID string) (*models.User, error) {
	return a.usersRepo.FindBySessionID(sessionID)
}

func (a *api) Create(user *models.User) (*models.User, error) {
	return a.usersRepo.Create(user)
}

func (a *api) Update(id uint64, user *models.User) (*models.User, error) {
	updateUser, err := a.GetByID(id)
	if err != nil {
		return nil, err
	}

	if err = copier.CopyWithOption(&updateUser, &user, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}
	return a.usersRepo.Update(updateUser)
}

func (a *api) DeleteByID(id uint64) error {
	return a.usersRepo.DeleteByID(id)
}
func (a *api) DeleteByUsername(username string) error {
	user, err := a.GetByUsername(username)
	if err != nil {
		return err
	}
	return a.DeleteByID(user.ID)
}
func (a *api) DeleteByEmail(email string) error {
	user, err := a.GetByEmail(email)
	if err != nil {
		return err
	}
	return a.DeleteByID(user.ID)
}

func (a *api) CheckIDAndPassword(id uint64, password string) bool {
	user, err := a.GetByID(id)
	if err != nil {
		return false
	}
	return user.Password == password
}

func (a *api) IsExistUsernameAndEmail(username, email string) bool {
	_, err := a.GetByUsername(username)
	if err == nil {
		if _, err = a.GetByEmail(email); err == nil {
			return true
		}
	}
	return false
}
