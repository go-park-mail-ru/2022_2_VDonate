package subscriptions

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"
	"github.com/jinzhu/copier"
)

type usecase struct {
	subRepo    domain.SubscriptionsRepository
	userRepo   domain.UsersRepository
	imgUseCase domain.ImageUseCase
}

func New(s domain.SubscriptionsRepository, u domain.UsersRepository, i domain.ImageUseCase) domain.SubscriptionsUseCase {
	return &usecase{
		subRepo:    s,
		userRepo:   u,
		imgUseCase: i,
	}
}

func (u *usecase) GetSubscriptionsByUserID(userID uint64) ([]models.AuthorSubscription, error) {
	s, err := u.subRepo.GetSubscriptionsByUserID(userID)
	if err != nil {
		return nil, err
	}

	if len(s) == 0 {
		return make([]models.AuthorSubscription, 0), nil
	}

	for i, subscription := range s {
		if s[i].Img, err = u.imgUseCase.GetImage(subscription.Img); err != nil {
			return nil, errorHandling.WrapEcho(domain.ErrInternal, err)
		}

		author, errAuthor := u.userRepo.GetByID(subscription.AuthorID)
		if errAuthor != nil {
			return nil, errorHandling.WrapEcho(domain.ErrInternal, errAuthor)
		}

		s[i].AuthorName = author.Username
		if s[i].AuthorAvatar, err = u.imgUseCase.GetImage(author.Avatar); err != nil {
			return nil, errorHandling.WrapEcho(domain.ErrInternal, err)
		}
	}

	return s, nil
}

func (u *usecase) GetAuthorSubscriptionsByAuthorID(authorID uint64) ([]models.AuthorSubscription, error) {
	s, err := u.subRepo.GetSubscriptionsByAuthorID(authorID)
	if err != nil {
		return nil, err
	}

	if len(s) == 0 {
		return make([]models.AuthorSubscription, 0), nil
	}

	for i, subscription := range s {
		if s[i].Img, err = u.imgUseCase.GetImage(subscription.Img); err != nil {
			return nil, errorHandling.WrapEcho(domain.ErrInternal, err)
		}
	}

	return s, nil
}

func (u *usecase) GetAuthorSubscriptionByID(id uint64) (models.AuthorSubscription, error) {
	s, err := u.subRepo.GetSubscriptionByID(id)
	if err != nil {
		return models.AuthorSubscription{}, err
	}

	if s.Img, err = u.imgUseCase.GetImage(s.Img); err != nil {
		return models.AuthorSubscription{}, errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	return s, nil
}

func (u *usecase) AddAuthorSubscription(sub models.AuthorSubscription, id uint64) (uint64, error) {
	sub.AuthorID = id
	return u.subRepo.AddSubscription(sub)
}

func (u *usecase) UpdateAuthorSubscription(sub models.AuthorSubscription, id uint64) error {
	updateSub, err := u.GetAuthorSubscriptionByID(id)
	if err != nil {
		return err
	}

	if err = copier.CopyWithOption(&updateSub, &sub, copier.Option{IgnoreEmpty: true}); err != nil {
		return err
	}

	return u.subRepo.UpdateSubscription(updateSub)
}

func (u *usecase) DeleteAuthorSubscription(subID uint64) error {
	return u.subRepo.DeleteSubscription(subID)
}
