package subscribers

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/interface"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
)

type usecase struct {
	subscribersRepo domain.SubscribersRepository
	userRepo        domain.UsersRepository
}

func New(s domain.SubscribersRepository, u domain.UsersRepository) domain.SubscribersUseCase {
	return &usecase{
		subscribersRepo: s,
		userRepo:        u,
	}
}

func (u usecase) GetSubscribers(authorID uint64) ([]models.User, error) {
	s, err := u.subscribersRepo.GetSubscribers(authorID)
	if err != nil {
		return nil, err
	}

	subs := make([]models.User, 0)

	for _, userID := range s {
		// Notion: if there is an error while getting user, skip it
		user, _ := u.userRepo.GetByID(userID)
		subs = append(subs, user)
	}

	return subs, nil
}

func (u usecase) Subscribe(subscription models.Subscription, userID uint64) error {
	subscription.SubscriberID = userID
	if utils.Empty(subscription.SubscriberID, subscription.AuthorID, subscription.AuthorSubscriptionID) {
		return domain.ErrBadRequest
	}
	return u.subscribersRepo.Subscribe(subscription)
}

func (u usecase) Unsubscribe(userID, authorID uint64) error {
	if userID == 0 || authorID == 0 {
		return domain.ErrBadRequest
	}
	return u.subscribersRepo.Unsubscribe(userID, authorID)
}

func (u usecase) IsSubscriber(userID, authorID uint64) (bool, error) {
	s, err := u.subscribersRepo.GetSubscribers(authorID)
	if err != nil {
		return false, err
	}

	for _, id := range s {
		if id == userID {
			return true, nil
		}
	}

	return false, nil
}
