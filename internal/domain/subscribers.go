package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type SubscribersUseCase interface {
	GetSubscribers(authorID uint64) ([]models.User, error)
	Subscribe(subscription models.Subscription, userID uint64, as models.AuthorSubscription) (interface{}, error)
	Unsubscribe(userID, authorID uint64) error

	IsSubscriber(userID, authorID uint64) (bool, error)
	Withdraw(userID uint64, phone, card string) (models.WithdrawInfo, error)
	CardValidation(card string) (models.WithdrawValidation, error)
	WithdrawCard(userID uint64, card, provider string) (models.WithdrawInfo, error)
	WithdrawQiwi(userID uint64, phone string) (models.WithdrawInfo, error)
}

type SubscribersMicroservice interface {
	GetSubscribers(authorID uint64) ([]uint64, error)
	Subscribe(payment models.Payment)
	Unsubscribe(subscription models.Subscription) error
}
