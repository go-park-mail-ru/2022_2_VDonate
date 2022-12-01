package domain

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type SubscriptionMicroservice interface {
	GetSubscriptionsByUserID(userID uint64) ([]models.AuthorSubscription, error)
	GetSubscriptionByUserAndAuthorID(userID, authorID uint64) (models.AuthorSubscription, error)
	GetSubscriptionByID(ID uint64) (models.AuthorSubscription, error)
	AddSubscription(sub models.AuthorSubscription) (uint64, error)
	UpdateSubscription(sub models.AuthorSubscription) error
	DeleteSubscription(subID uint64) error
}

type SubscriptionsUseCase interface {
	GetSubscriptionsByUserID(userID uint64) ([]models.AuthorSubscription, error)

	GetAuthorSubscriptionsByAuthorID(authorID uint64) ([]models.AuthorSubscription, error)
	GetAuthorSubscriptionByID(ID uint64) (models.AuthorSubscription, error)
	AddAuthorSubscription(sub models.AuthorSubscription, id uint64) (uint64, error)
	UpdateAuthorSubscription(sub models.AuthorSubscription, id uint64) error
	DeleteAuthorSubscription(subID uint64) error
}
