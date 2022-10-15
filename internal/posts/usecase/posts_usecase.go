package posts

import (
	"errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	postsDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts"
	"github.com/jinzhu/copier"
)

type usecase struct {
	postsRepo postsDomain.Repository
}

func New(repo postsDomain.Repository) postsDomain.UseCase {
	return &usecase{postsRepo: repo}
}

func (u *usecase) GetPostsByUserID(id uint64) ([]*models.PostDB, error) {
	r, err := u.postsRepo.GetAllByUserID(id)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("no posts were found")
	}
	return r, nil
}
func (u *usecase) GetPostByID(postID uint64) (*models.PostDB, error) {
	return u.postsRepo.GetPostByID(postID)
}
func (u *usecase) Create(post models.PostDB) (*models.PostDB, error) {
	return u.postsRepo.Create(post)
}
func (u *usecase) Update(post models.PostDB) (*models.PostDB, error) {
	updatePost, err := u.GetPostByID(post.ID)
	if err != nil {
		return nil, err
	}

	if err = copier.CopyWithOption(updatePost, &post, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}

	if updatePost, err = u.postsRepo.Update(*updatePost); err != nil {
		return nil, err
	}

	return updatePost, nil
}
func (u *usecase) DeleteByID(postID uint64) error {
	return u.postsRepo.DeleteByID(postID)
}
