package subscriptions

import (
	"errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jinzhu/copier"
)

type UseCase interface {
	GetSubscriptions(authorID uint64) ([]*models.AuthorSubscription, error)
	AddSubscription(sub *models.AuthorSubscription) (*models.AuthorSubscription, error)
	UpdateSubscription(sub *models.AuthorSubscription) (*models.AuthorSubscription, error)
	DeleteSubscription(subID uint64) error
}

type Repository interface {
	GetSubscriptions(authorID uint64) ([]*models.AuthorSubscription, error)
	AddSubscription(sub *models.AuthorSubscription) (*models.AuthorSubscription, error)
	UpdateSubscription(sub *models.AuthorSubscription) (*models.AuthorSubscription, error)
	DeleteSubscription(subID uint64) error
}

type usecase struct {
	subRepo Repository
}

func New(subRepo Repository) UseCase {
	return &usecase{subRepo: subRepo}
}

func (u *usecase) GetSubscriptions(authorID uint64) ([]*models.AuthorSubscription, error) {
	s, err := u.subRepo.GetSubscriptions(authorID)
	if err != nil {
		return nil, err
	}
	if len(s) == 0 {
		return nil, errors.New("no subscribers")
	}
	return s, nil
}

func (u *usecase) AddSubscription(sub *models.AuthorSubscription) (*models.AuthorSubscription, error) {
	return u.subRepo.AddSubscription(sub)
}

func (u *usecase) UpdateSubscription(sub *models.AuthorSubscription) (*models.AuthorSubscription, error) {
	updateSub, err := u.GetSubscriptions(sub.AuthorID)
	if err != nil {
		return nil, err
	}

	if err = copier.CopyWithOption(&updateSub, &sub, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}

	return u.subRepo.UpdateSubscription(sub)
}

func (u *usecase) DeleteSubscription(subID uint64) error {
	return u.subRepo.DeleteSubscription(subID)
}
