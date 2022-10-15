package subscribersDomain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type UseCase interface {
	GetSubscribers(authorID uint64) ([]*models.User, error)
	Subscribe(subscription models.Subscription) error
	Unsubscribe(userID, authorID uint64) error

	IsSubscriber(userID, authorID uint64) (bool, error)
}

type Repository interface {
	GetSubscribers(authorID uint64) ([]uint64, error)
	Subscribe(subscription models.Subscription) error
	Unsubscribe(userID, authorID uint64) error
}
