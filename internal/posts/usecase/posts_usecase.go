package posts

import (
	"errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type UseCase interface {
	GetPostsByUserID(id uint64) ([]*models.PostDB, error)
	GetPostByID(postID uint64) (*models.PostDB, error)
	Create(post *models.PostDB) (*models.PostDB, error)
	DeleteInUserByID(userID, postID uint64) error
}

type Repository interface {
	GetAllByUserID(userID uint64) ([]*models.PostDB, error)
	GetPostByUserID(userID, postID uint64) (*models.PostDB, error)
	GetPostByID(postID uint64) (*models.PostDB, error)
	Create(post *models.PostDB) (*models.PostDB, error)
	DeleteInUserByID(userID, postID uint64) error
	Close() error
}

type api struct {
	postsRepo Repository
}

func New(repo Repository) UseCase {
	return &api{postsRepo: repo}
}

func (a *api) GetPostsByUserID(id uint64) ([]*models.PostDB, error) {
	r, err := a.postsRepo.GetAllByUserID(id)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("no posts")
	}
	return r, nil
}
func (a *api) GetPostByID(postID uint64) (*models.PostDB, error) {
	return a.postsRepo.GetPostByID(postID)
}
func (a *api) Create(post *models.PostDB) (*models.PostDB, error) {
	return a.postsRepo.Create(post)
}
func (a *api) DeleteInUserByID(userID, postID uint64) error {
	return a.postsRepo.DeleteInUserByID(userID, postID)
}
