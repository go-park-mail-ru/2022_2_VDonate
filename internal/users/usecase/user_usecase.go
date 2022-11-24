package users

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/jinzhu/copier"
)

type hashCreator func(password string) (string, error)

type usecase struct {
	usersRepo  domain.UsersRepository
	imgUseCase domain.ImageUseCase

	hashCreator hashCreator
}

func New(usersRepo domain.UsersRepository, imgUseCase domain.ImageUseCase) domain.UsersUseCase {
	return &usecase{
		usersRepo:  usersRepo,
		imgUseCase: imgUseCase,

		hashCreator: utils.HashPassword,
	}
}

func WithHashCreator(usersRepo domain.UsersRepository, imgUseCase domain.ImageUseCase, hashCreator hashCreator) domain.UsersUseCase {
	return &usecase{
		usersRepo:  usersRepo,
		imgUseCase: imgUseCase,

		hashCreator: hashCreator,
	}
}

func (u *usecase) GetByID(id uint64) (models.User, error) {
	user, err := u.usersRepo.GetByID(id)
	if err != nil {
		return models.User{}, err
	}
	if user.Avatar, err = u.imgUseCase.GetImage(user.Avatar); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *usecase) GetByUsername(username string) (models.User, error) {
	user, err := u.usersRepo.GetByUsername(username)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *usecase) GetByEmail(email string) (models.User, error) {
	user, err := u.usersRepo.GetByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *usecase) GetBySessionID(sessionID string) (models.User, error) {
	user, err := u.usersRepo.GetBySessionID(sessionID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *usecase) GetUserByPostID(postID uint64) (models.User, error) {
	user, err := u.usersRepo.GetUserByPostID(postID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *usecase) Create(user models.User) (uint64, error) {
	return u.usersRepo.Create(user)
}

func (u *usecase) Update(user models.User, id uint64) (models.User, error) {
	updateUser, err := u.GetByID(id)
	if err != nil {
		return models.User{}, err
	}

	if len(user.Password) != 0 {
		if user.Password, err = u.hashCreator(user.Password); err != nil {
			return models.User{}, err
		}
	}

	if err = copier.CopyWithOption(&updateUser, &user, copier.Option{IgnoreEmpty: true}); err != nil {
		return models.User{}, err
	}

	return updateUser, u.usersRepo.Update(updateUser)
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
