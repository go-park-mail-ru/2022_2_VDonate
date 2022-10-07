package postsAPI

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	postsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/repository"
)

type UseCase interface {
	GetAllByUserID(id uint) ([]*models.PostDB, error)
	GetByUserIDAndPostID(userID, postID uint) error
	CreateInUserByID(post *models.PostDB, userID uint) error
	DeleteInUserByID(userID, postID uint) error
}

type API struct {
	postsRepo postsRepository.API
}

func New(repo postsRepository.API) UseCase {
	return &API{postsRepo: repo}
}

func (a *API) GetAllByUserID(id uint) ([]*models.PostDB, error) {
	return a.postsRepo.GetAllByUserID(id)
}
func (a *API) GetByUserIDAndPostID(userID, postID uint) error {
	return a.GetByUserIDAndPostID(userID, postID)
}
func (a *API) CreateInUserByID(post *models.PostDB, userID uint) error {
	return a.CreateInUserByID(post, userID)
}
func (a *API) DeleteInUserByID(userID, postID uint) error {
	return a.DeleteInUserByID(userID, postID)
}
