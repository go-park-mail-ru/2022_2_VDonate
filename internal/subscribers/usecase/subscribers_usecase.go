package subscribers

import (
	"errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	subscribersDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers"
	users "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
)

type usecase struct {
	subscribersRepo subscribersDomain.Repository
	userRepo        users.Repository
}

func New(subscribersRepo subscribersDomain.Repository, userRepo users.Repository) subscribersDomain.UseCase {
	return &usecase{
		subscribersRepo: subscribersRepo,
		userRepo:        userRepo,
	}
}

func (u *usecase) GetSubscribers(authorID uint64) ([]*models.User, error) {
	s, err := u.subscribersRepo.GetSubscribers(authorID)
	if err != nil {
		return nil, err
	}
	if len(s) == 0 {
		return nil, errors.New("no subscribers")
	}

	var subs []*models.User

	for _, userID := range s {
		// Notion: if there is an error while getting user, skip it
		user, _ := u.userRepo.GetByID(userID)
		subs = append(subs, user)
	}

	return subs, nil
}

func (u *usecase) Subscribe(subscription models.Subscription) error {
	return u.subscribersRepo.Subscribe(subscription)
}

func (u *usecase) Unsubscribe(userID, authorID uint64) error {
	return u.subscribersRepo.Unsubscribe(userID, authorID)
}

func (u usecase) IsSubscriber(userID, authorID uint64) (bool, error) {
	s, err := u.GetSubscribers(authorID)
	if err != nil {
		return false, err
	}

	if !models.Contains[*models.User](s, userID) {
		return false, errors.New("user is not a subscriber")
	}

	return true, nil
}
