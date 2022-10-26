package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type SubscriptionsUseCase interface {
	GetSubscriptionsByAuthorID(authorID uint64) ([]*models.AuthorSubscription, error)
	GetSubscriptionsByID(ID uint64) (*models.AuthorSubscription, error)
	AddSubscription(sub models.AuthorSubscription) error
	UpdateSubscription(sub models.AuthorSubscription) error
	DeleteSubscription(subID uint64) error
}
