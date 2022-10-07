package userRepository

import model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type API interface {
	Create(u *model.UserDB) (*model.UserDB, error)
	FindByUsername(username string) (*model.UserDB, error)
	FindByID(id uint) (*model.UserDB, error)
	FindByEmail(email string) (*model.UserDB, error)
	Update(u *model.UserDB) (*model.UserDB, error)
	DeleteByID(id uint) error
	Close() error
}
