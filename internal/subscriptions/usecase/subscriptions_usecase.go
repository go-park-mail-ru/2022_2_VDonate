package subscriptions

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jinzhu/copier"
)

type usecase struct {
	subscriptionsRepo domain.SubscriptionsRepository
}

func New(s domain.SubscriptionsRepository) domain.SubscriptionsUseCase {
	return &usecase{
		subscriptionsRepo: s,
	}
}

func (u *usecase) GetSubscriptionsByUserID(userID uint64) ([]models.AuthorSubscription, error) {
	s, err := u.subscriptionsRepo.GetSubscriptionsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (u *usecase) GetAuthorSubscriptionsByAuthorID(authorID uint64) ([]models.AuthorSubscription, error) {
	s, err := u.subscriptionsRepo.GetSubscriptionsByAuthorID(authorID)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (u *usecase) GetAuthorSubscriptionByID(ID uint64) (models.AuthorSubscription, error) {
	s, err := u.subscriptionsRepo.GetSubscriptionsByID(ID)
	if err != nil {
		return models.AuthorSubscription{}, err
	}
	return s, nil
}

func (u *usecase) AddAuthorSubscription(sub models.AuthorSubscription) (models.AuthorSubscription, error) {
	return u.subscriptionsRepo.AddSubscription(sub)
}

func (u *usecase) UpdateAuthorSubscription(sub models.AuthorSubscription) (models.AuthorSubscription, error) {
	updateSub, err := u.GetAuthorSubscriptionByID(sub.ID)
	if err != nil {
		return models.AuthorSubscription{}, err
	}

	if err = copier.CopyWithOption(&updateSub, &sub, copier.Option{IgnoreEmpty: true}); err != nil {
		return models.AuthorSubscription{}, err
	}

	return u.subscriptionsRepo.UpdateSubscription(updateSub)
}

func (u *usecase) DeleteAuthorSubscription(subID uint64) error {
	return u.subscriptionsRepo.DeleteSubscription(subID)
}
