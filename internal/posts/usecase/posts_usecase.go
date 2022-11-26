package posts

import (
	"sort"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"
	"github.com/jinzhu/copier"
)

type usecase struct {
	postsRepo         domain.PostsRepository
	userRepo          domain.UsersRepository
	subscriptionsRepo domain.SubscriptionsRepository
	imgUseCase        domain.ImageUseCase
}

func New(p domain.PostsRepository, u domain.UsersRepository, i domain.ImageUseCase, s domain.SubscriptionsRepository) domain.PostsUseCase {
	return &usecase{
		postsRepo:         p,
		userRepo:          u,
		subscriptionsRepo: s,
		imgUseCase:        i,
	}
}

func (u usecase) GetPostsByFilter(userID, authorID uint64) ([]models.Post, error) {
	var r []models.Post
	var err error
	validate := true

	switch {
	case authorID == 0:
		if r, err = u.postsRepo.GetPostsBySubscriptions(userID); err != nil {
			return nil, err
		}
		validate = false
	case authorID > 0:
		if r, err = u.postsRepo.GetAllByUserID(authorID); err != nil {
			return nil, err
		}
		if authorID == userID {
			validate = false
		}
	default:
		return nil, domain.ErrInternal
	}

	for i, post := range r {
		if validate {
			as, errSubscriptions := u.subscriptionsRepo.GetSubscriptionByUserAndAuthorID(userID, authorID)
			if errSubscriptions != nil {
				return nil, errSubscriptions
			}

			if as.Tier >= post.Tier {
				r[i].IsAllowed = true
			}
		} else {
			r[i].IsAllowed = true
		}

		if r[i].IsAllowed {
			if r[i].Img, err = u.imgUseCase.GetImage(post.Img); err != nil {
				return nil, err
			}
		} else {
			r[i].Img = ""
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

	sort.Slice(r, func(i, j int) bool {
		return r[i].ID > r[j].ID
	})

	return r, nil
}

func (u usecase) GetPostByID(postID, userID uint64) (models.Post, error) {
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

func (u usecase) Create(post models.Post, userID uint64) (uint64, error) {
	post.UserID = userID
	return u.postsRepo.Create(post)
}

func (u usecase) Update(post models.Post, postID uint64) error {
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

func (u usecase) DeleteByID(postID uint64) error {
	return u.postsRepo.DeleteByID(postID)
}

func (u usecase) GetLikesByPostID(postID uint64) ([]models.Like, error) {
	return u.postsRepo.GetAllLikesByPostID(postID)
}

func (u usecase) GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error) {
	return u.postsRepo.GetLikeByUserAndPostID(userID, postID)
}

func (u usecase) LikePost(userID, postID uint64) error {
	return u.postsRepo.CreateLike(userID, postID)
}

func (u usecase) UnlikePost(userID, postID uint64) error {
	return u.postsRepo.DeleteLikeByID(userID, postID)
}

func (u usecase) GetLikesNum(postID uint64) (uint64, error) {
	likes, err := u.GetLikesByPostID(postID)
	if err != nil {
		return 0, err
	}
	return uint64(len(likes)), nil
}

func (u usecase) IsPostLiked(userID, postID uint64) bool {
	if _, err := u.GetLikeByUserAndPostID(userID, postID); err != nil {
		return false
	}
	return true
}
