package users

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/jinzhu/copier"
)

type usecase struct {
	usersRepo domain.UsersRepository
}

func New(usersRepo domain.UsersRepository) domain.UsersUseCase {
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

func (u *usecase) GetUserByPostID(postID uint64) (*models.User, error) {
	return u.usersRepo.GetUserByPostID(postID)
}

func (u *usecase) Create(user models.User) (*models.User, error) {
	return u.usersRepo.Create(&user)
}

func (u *usecase) Update(user models.User) (*models.User, error) {
	updateUser, err := u.GetByID(user.ID)
	if err != nil {
		return nil, err
	}
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	if err = copier.CopyWithOption(updateUser, &user, copier.Option{IgnoreEmpty: true}); err != nil {
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
	return utils.CheckHashPassword(password, user.Password)
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
