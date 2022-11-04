package subscriptions

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jinzhu/copier"
)

type usecase struct {
	subRepo domain.SubscriptionsRepository
}

func New(s domain.SubscriptionsRepository) domain.SubscriptionsUseCase {
	return &usecase{
		subRepo: s,
	}
}

func (u *usecase) GetSubscriptionsByUserID(userID uint64) ([]models.AuthorSubscription, error) {
	s, err := u.subRepo.GetSubscriptionsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (u *usecase) GetAuthorSubscriptionsByAuthorID(authorID uint64) ([]models.AuthorSubscription, error) {
	s, err := u.subRepo.GetSubscriptionsByAuthorID(authorID)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (u *usecase) GetAuthorSubscriptionByID(id uint64) (models.AuthorSubscription, error) {
	s, err := u.subRepo.GetSubscriptionByID(id)
	if err != nil {
		return models.AuthorSubscription{}, err
	}
	return s, nil
}

func (u *usecase) AddAuthorSubscription(sub models.AuthorSubscription, id uint64) error {
	sub.ID = id
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
