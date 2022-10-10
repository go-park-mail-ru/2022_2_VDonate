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

type usecase struct {
	postsRepo Repository
}

func New(repo Repository) UseCase {
	return &usecase{postsRepo: repo}
}

func (u *usecase) GetPostsByUserID(id uint64) ([]*models.PostDB, error) {
	r, err := u.postsRepo.GetAllByUserID(id)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("no posts")
	}
	return r, nil
}
func (u *usecase) GetPostByID(postID uint64) (*models.PostDB, error) {
	return u.postsRepo.GetPostByID(postID)
}
func (u *usecase) Create(post *models.PostDB) (*models.PostDB, error) {
	return u.postsRepo.Create(post)
}
func (u *usecase) DeleteInUserByID(userID, postID uint64) error {
	return u.postsRepo.DeleteInUserByID(userID, postID)
}
