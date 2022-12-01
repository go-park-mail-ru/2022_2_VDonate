package subscribers

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
)

type usecase struct {
	subscribersMicroservice domain.SubscribersMicroservice
	userMicroservice        domain.UsersMicroservice
}

func New(s domain.SubscribersMicroservice, u domain.UsersMicroservice) domain.SubscribersUseCase {
	return &usecase{
		subscribersMicroservice: s,
		userMicroservice:        u,
	}
}

func (u usecase) GetSubscribers(authorID uint64) ([]models.User, error) {
	s, err := u.subscribersMicroservice.GetSubscribers(authorID)
	if err != nil {
		return nil, err
	}

	subs := make([]models.User, 0)

	for _, userID := range s {
		// Notion: if there is an error while getting user, skip it
		user, _ := u.userMicroservice.GetByID(userID)
		subs = append(subs, user)
	}

	return subs, nil
}

func (u usecase) Subscribe(subscription models.Subscription, userID uint64) error {
	subscription.SubscriberID = userID
	if utils.Empty(subscription.SubscriberID, subscription.AuthorID, subscription.AuthorSubscriptionID) {
		return domain.ErrBadRequest
	}
	return u.subscribersMicroservice.Subscribe(subscription)
}

func (u usecase) Unsubscribe(userID, authorID uint64) error {
	if userID == 0 || authorID == 0 {
		return domain.ErrBadRequest
	}
	return u.subscribersMicroservice.Unsubscribe(models.Subscription{
		AuthorID:     authorID,
		SubscriberID: userID,
	})
}

func (u usecase) IsSubscriber(userID, authorID uint64) (bool, error) {
	s, err := u.subscribersMicroservice.GetSubscribers(authorID)
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
