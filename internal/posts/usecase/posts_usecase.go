package posts

import (
	"errors"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jinzhu/copier"
)

type usecase struct {
	postsRepo domain.PostsRepository
}

func New(repo domain.PostsRepository) domain.PostsUseCase {
	return &usecase{postsRepo: repo}
}

func (u *usecase) GetPostsByUserID(id uint64) ([]*models.Post, error) {
	r, err := u.postsRepo.GetAllByUserID(id)
	if err != nil {
		return nil, err
	}

	if len(r) == 0 {
		return nil, errors.New("no posts were found")
	}

	return r, nil
}

func (u *usecase) GetPostByID(postID uint64) (*models.Post, error) {
	return u.postsRepo.GetPostByID(postID)
}

func (u *usecase) Create(post models.Post) error {
	return u.postsRepo.Create(post)
}

func (u *usecase) Update(post models.Post) error {
	updatePost, err := u.GetPostByID(post.ID)
	if err != nil {
		return err
	}

	if err = copier.CopyWithOption(updatePost, &post, copier.Option{IgnoreEmpty: true}); err != nil {
		return err
	}

	return u.postsRepo.Update(*updatePost)
}

func (u *usecase) DeleteByID(postID uint64) error {
	return u.postsRepo.DeleteByID(postID)
}
