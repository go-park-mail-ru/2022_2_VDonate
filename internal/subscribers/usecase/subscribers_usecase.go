package subscribers

import (
	"errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
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
