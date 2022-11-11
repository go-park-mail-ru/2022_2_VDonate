package posts

import (
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

func (u *usecase) GetPostsByUserID(id uint64) ([]models.Post, error) {
	r, err := u.postsRepo.GetAllByUserID(id)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (u *usecase) GetPostsByFilter(filter string, userID uint64) ([]models.Post, error) {
	switch filter {
	case "subscriptions":
		return u.postsRepo.GetPostsBySubscriptions(userID)
	default:
		return nil, domain.ErrBadRequest
	}
}

func (u *usecase) GetPostByID(postID uint64) (models.Post, error) {
	return u.postsRepo.GetPostByID(postID)
}

func (u *usecase) Create(post models.Post, userID uint64) (uint64, error) {
	post.UserID = userID
	return u.postsRepo.Create(post)
}

func (u *usecase) Update(post models.Post, postID uint64) error {
	updatePost, err := u.GetPostByID(postID)
	if err != nil {
		return err
	}

	if err = copier.CopyWithOption(&updatePost, &post, copier.Option{IgnoreEmpty: true}); err != nil {
		return err
	}

	return u.postsRepo.Update(updatePost)
}

func (u *usecase) DeleteByID(postID uint64) error {
	return u.postsRepo.DeleteByID(postID)
}

func (u *usecase) GetLikesByPostID(postID uint64) ([]models.Like, error) {
	r, err := u.postsRepo.GetAllLikesByPostID(postID)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (u *usecase) GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error) {
	return u.postsRepo.GetLikeByUserAndPostID(userID, postID)
}

func (u *usecase) LikePost(userID, postID uint64) error {
	return u.postsRepo.CreateLike(userID, postID)
}

func (u *usecase) UnlikePost(userID, postID uint64) error {
	return u.postsRepo.DeleteLikeByID(userID, postID)
}

func (u *usecase) GetLikesNum(postID uint64) (uint64, error) {
	likes, err := u.GetLikesByPostID(postID)
	if err != nil {
		return 0, err
	}
	return uint64(len(likes)), nil
}

func (u *usecase) IsPostLiked(userID, postID uint64) bool {
	if _, err := u.GetLikeByUserAndPostID(userID, postID); err != nil {
		return false
	}
	return true
}
