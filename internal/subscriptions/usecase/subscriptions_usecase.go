package subscriptions

import (
	"errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	subscriptionsDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions"
	"github.com/jinzhu/copier"
)

type usecase struct {
	subRepo subscriptionsDomain.Repository
}

func New(subRepo subscriptionsDomain.Repository) subscriptionsDomain.UseCase {
	return &usecase{subRepo: subRepo}
}

func (u *usecase) GetSubscriptionsByAuthorID(authorID uint64) ([]*models.AuthorSubscription, error) {
	s, err := u.subRepo.GetSubscriptionsByAuthorID(authorID)
	if err != nil {
		return nil, err
	}
	if len(s) == 0 {
		return nil, errors.New("no subscriptions")
	}
	return s, nil
}

func (u *usecase) GetSubscriptionsByID(ID uint64) (*models.AuthorSubscription, error) {
	s, err := u.subRepo.GetSubscriptionsByID(ID)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (u *usecase) AddSubscription(sub models.AuthorSubscription) (*models.AuthorSubscription, error) {
	return u.subRepo.AddSubscription(sub)
}

func (u *usecase) UpdateSubscription(sub models.AuthorSubscription) (*models.AuthorSubscription, error) {
	updateSub, err := u.GetSubscriptionsByID(sub.ID)
	if err != nil {
		return nil, err
	}

	if err = copier.CopyWithOption(updateSub, &sub, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}

	return u.subRepo.UpdateSubscription(updateSub)
}

func (u *usecase) DeleteSubscription(subID uint64) error {
	return u.subRepo.DeleteSubscription(subID)
}
