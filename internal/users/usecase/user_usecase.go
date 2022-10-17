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
	GetByUsername(username string) (*models.User, error)
	GetByID(id uint64) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetBySessionID(sessionID string) (*models.User, error)
	Update(u *models.User) (*models.User, error)
	DeleteByID(id uint64) error
	Close() error
}

type usecase struct {
	usersRepo Repository
}

func New(usersRepo Repository) UseCase {
	return &usecase{
		usersRepo: usersRepo,
	}
}

func (u *usecase) GetByID(id uint64) (*models.User, error) {
	return u.usersRepo.GetByID(id)
}

func (u *usecase) GetByUsername(username string) (*models.User, error) {
	return u.usersRepo.GetByUsername(username)
}

func (u *usecase) GetByEmail(email string) (*models.User, error) {
	return u.usersRepo.GetByEmail(email)
}

func (u *usecase) GetBySessionID(sessionID string) (*models.User, error) {
	return u.usersRepo.GetBySessionID(sessionID)
}

func (u *usecase) Create(user *models.User) (*models.User, error) {
	return u.usersRepo.Create(user)
}

func (u *usecase) Update(id uint64, user *models.User) (*models.User, error) {
	updateUser, err := u.GetByID(id)
	if err != nil {
		return nil, err
	}

	if err = copier.CopyWithOption(&updateUser, &user, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}
	return u.usersRepo.Update(updateUser)
}

func (u *usecase) DeleteByID(id uint64) error {
	return u.usersRepo.DeleteByID(id)
}
func (u *usecase) DeleteByUsername(username string) error {
	user, err := u.GetByUsername(username)
	if err != nil {
		return err
	}
	return u.DeleteByID(user.ID)
}
func (u *usecase) DeleteByEmail(email string) error {
	user, err := u.GetByEmail(email)
	if err != nil {
		return err
	}
	return u.DeleteByID(user.ID)
}

func (u *usecase) CheckIDAndPassword(id uint64, password string) bool {
	user, err := u.GetByID(id)
	if err != nil {
		return false
	}
	return user.Password == password
}

func (u *usecase) IsExistUsernameAndEmail(username, email string) bool {
	_, err := u.GetByUsername(username)
	if err == nil {
		if _, err = u.GetByEmail(email); err == nil {
			return true
		}
	}
	return false
}
