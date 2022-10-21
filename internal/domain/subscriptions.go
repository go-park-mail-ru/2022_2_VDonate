package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type SubscriptionsUseCase interface {
	GetSubscriptionsByUserID(userID uint64) ([]models.AuthorSubscription, error)

	GetAuthorSubscriptionsByAuthorID(authorID uint64) ([]models.AuthorSubscription, error)
	GetAuthorSubscriptionByID(ID uint64) (models.AuthorSubscription, error)
	AddAuthorSubscription(sub models.AuthorSubscription) (models.AuthorSubscription, error)
	UpdateAuthorSubscription(sub models.AuthorSubscription) (models.AuthorSubscription, error)
	DeleteAuthorSubscription(subID uint64) error
}
