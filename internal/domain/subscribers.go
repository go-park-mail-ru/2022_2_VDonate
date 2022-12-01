package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type SubscribersUseCase interface {
	GetSubscribers(authorID uint64) ([]models.User, error)
	Subscribe(subscription models.Subscription, userID uint64) error
	Unsubscribe(userID, authorID uint64) error

	IsSubscriber(userID, authorID uint64) (bool, error)
}

type SubscribersMicroservice interface {
	GetSubscribers(userID uint64) ([]models.User, error)
	Subscribe(subscriber models.Subscription) error
	Unsubscribe(subscriber models.Subscription) error
}
