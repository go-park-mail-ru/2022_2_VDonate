package posts

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"
	"github.com/jinzhu/copier"
)

type usecase struct {
	postsRepo  domain.PostsRepository
	userRepo   domain.UsersRepository
	imgUseCase domain.ImageUseCase
}

func New(postsRepo domain.PostsRepository, userRepo domain.UsersRepository, imgUseCase domain.ImageUseCase) domain.PostsUseCase {
	return &usecase{
		postsRepo:  postsRepo,
		userRepo:   userRepo,
		imgUseCase: imgUseCase,
	}
}

func (u *usecase) GetPostsByFilter(filter string, userID uint64) ([]models.Post, error) {
	r := make([]models.Post, 0)
	var err error

	switch filter {
	case "subscriptions":
		if r, err = u.postsRepo.GetPostsBySubscriptions(userID); err != nil {
			return nil, err
		}
	default:
		if r, err = u.postsRepo.GetAllByUserID(userID); err != nil {
			return nil, err
		}
	}

	for i, post := range r {
		if r[i].Img, err = u.imgUseCase.GetImage(post.Img); err != nil {
			return nil, err
		}

		author, errGetAuthor := u.userRepo.GetByID(post.UserID)
		if errGetAuthor != nil {
			return nil, errGetAuthor
		}

		r[i].Author.UserID = author.ID
		r[i].Author.Username = author.Username
		if r[i].Author.ImgPath, err = u.imgUseCase.GetImage(author.Avatar); err != nil {
			return nil, err
		}

		if r[i].LikesNum, err = u.GetLikesNum(post.ID); err != nil {
			return nil, domain.ErrInternal
		}
		r[i].IsLiked = u.IsPostLiked(userID, post.ID)
	}

	return r, nil
}

func (u *usecase) GetPostByID(postID, userID uint64) (models.Post, error) {
	r, err := u.postsRepo.GetPostByID(postID)
	if err != nil {
		return models.Post{}, err
	}

	if r.Img, err = u.imgUseCase.GetImage(r.Img); err != nil {
		return models.Post{}, err
	}

	author, errGetAuthor := u.userRepo.GetByID(r.UserID)
	if errGetAuthor != nil {
		return models.Post{}, err
	}

	r.Author.UserID = author.ID
	r.Author.Username = author.Username
	if r.Author.ImgPath, err = u.imgUseCase.GetImage(author.Avatar); err != nil {
		return models.Post{}, err
	}

	if r.LikesNum, err = u.GetLikesNum(postID); err != nil {
		return models.Post{}, errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	r.IsLiked = u.IsPostLiked(userID, postID)

	return r, nil
}

func (u *usecase) Create(post models.Post, userID uint64) (uint64, error) {
	post.UserID = userID
	return u.postsRepo.Create(post)
}

func (u *usecase) Update(post models.Post, postID uint64) error {
	var err error

	updatePost, err := u.GetPostByID(postID, post.UserID)
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
	return u.postsRepo.GetAllLikesByPostID(postID)
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
