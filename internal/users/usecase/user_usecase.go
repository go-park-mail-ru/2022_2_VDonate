package users

import (
	"database/sql"
	"errors"
	"github.com/ztrue/tracerr"
	"mime/multipart"

	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/jinzhu/copier"
)

type hashCreator func(password string) (string, error)

type usecase struct {
	usersMicroservice domain.UsersMicroservice
	imgUseCase        domain.ImageUseCase

	hashCreator hashCreator
}

func New(usersMicroservice domain.UsersMicroservice, imgUseCase domain.ImageUseCase) domain.UsersUseCase {
	return &usecase{
		usersMicroservice: usersMicroservice,
		imgUseCase:        imgUseCase,

		hashCreator: utils.HashPassword,
	}
}

func WithHashCreator(usersMicroservice domain.UsersMicroservice, imgUseCase domain.ImageUseCase, hashCreator hashCreator) domain.UsersUseCase {
	return &usecase{
		usersMicroservice: usersMicroservice,
		imgUseCase:        imgUseCase,

		hashCreator: hashCreator,
	}
}

func (u usecase) GetByID(id uint64) (models.User, error) {
	user, err := u.usersMicroservice.GetByID(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u usecase) GetByUsername(username string) (models.User, error) {
	return u.usersMicroservice.GetByUsername(username)
}

func (u usecase) GetByEmail(email string) (models.User, error) {
	return u.usersMicroservice.GetByEmail(email)
}

func (u usecase) GetBySessionID(sessionID string) (models.User, error) {
	return u.usersMicroservice.GetBySessionID(sessionID)
}

func (u usecase) GetUserByPostID(postID uint64) (models.User, error) {
	return u.usersMicroservice.GetUserByPostID(postID)
}

func (u usecase) Create(user models.User) (uint64, error) {
	return u.usersMicroservice.Create(user)
}

func (u usecase) Update(user models.User, file *multipart.FileHeader, id uint64) (models.User, error) {
	updateUser, err := u.GetByID(id)
	if err != nil {
		return models.User{}, err
	}

	if file != nil {
		if updateUser.Avatar, err = u.imgUseCase.CreateOrUpdateImage(file, updateUser.Avatar); err != nil {
			return models.User{}, errorHandling.WrapEcho(domain.ErrCreate, err)
		}
	}

	if len(user.Password) != 0 {
		if user.Password, err = u.hashCreator(user.Password); err != nil {
			return models.User{}, err
		}
	}

	if err = copier.CopyWithOption(&updateUser, &user, copier.Option{IgnoreEmpty: true}); err != nil {
		return models.User{}, err
	}

	return updateUser, u.usersMicroservice.Update(updateUser)
}

func (u usecase) CheckIDAndPassword(id uint64, password string) bool {
	user, err := u.GetByID(id)
	if err != nil {
		return false
	}

	return utils.CheckHashPassword(password, user.Password)
}

func (u usecase) IsExistUsernameAndEmail(username, email string) bool {
	_, err := u.GetByUsername(username)
	if err == nil {
		if _, err = u.GetByEmail(email); err == nil {
			return true
		}
	}

	return false
}

func (u usecase) FindAuthors(keyword string) ([]models.User, error) {
	var allAuthors []models.User
	var err error

	if len(keyword) == 0 {
		if allAuthors, err = u.usersMicroservice.GetAllAuthors(); err != nil {
			return nil, tracerr.Wrap(err)
		}
	} else {
		if allAuthors, err = u.usersMicroservice.GetAuthorByUsername(keyword); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, tracerr.Wrap(err)
		}
	}

	result := make([]models.User, 0)

	for _, a := range allAuthors {
		if !slices.Contains(result, a) {
			if a.Avatar, err = u.imgUseCase.GetImage(a.Avatar); err != nil {
				return nil, err
			}
			result = append(result, a)
		}
	}

	return result, nil
}

func (u usecase) GetPostsNum(userID uint64) (uint64, error) {
	postsNum, err := u.usersMicroservice.GetPostsNum(userID)
	if err != nil {
		return 0, err
	}
	return postsNum, nil
}

func (u usecase) GetSubscribersNum(userID uint64) (uint64, error) {
	subscribersNum, err := u.usersMicroservice.GetSubscribersNum(userID)
	if err != nil {
		return 0, err
	}
	return subscribersNum, nil
}

func (u usecase) GetProfitForMounth(userID uint64) (uint64, error) {
	profit, err := u.usersMicroservice.GetProfitForMounth(userID)
	if err != nil {
		return 0, err
	}
	return profit, nil
}
