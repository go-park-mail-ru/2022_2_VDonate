package subscribersMicroservice

import (
	"context"

	"github.com/ztrue/tracerr"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/protobuf"

	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type SubscribersMicroservice struct {
	subscribersClient protobuf.SubscribersClient
}

func New(subscribersClient protobuf.SubscribersClient) domain.SubscribersMicroservice {
	return &SubscribersMicroservice{
		subscribersClient: subscribersClient,
	}
}

func (m SubscribersMicroservice) GetSubscribers(userID uint64) ([]uint64, error) {
	subscribers, err := m.subscribersClient.GetSubscribers(context.Background(), &userProto.UserID{
		UserId: userID,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	res := make([]uint64, 0)
	for _, id := range subscribers.GetIds() {
		res = append(res, id.GetUserId())
	}

	return res, nil
}

func (m SubscribersMicroservice) Subscribe(subscriber models.Subscription) error {
	_, err := m.subscribersClient.Subscribe(context.Background(), &protobuf.Subscriber{
		AuthorID:             subscriber.AuthorID,
		SubscriberID:         subscriber.SubscriberID,
		AuthorSubscriptionID: subscriber.AuthorSubscriptionID,
	})

	return tracerr.Wrap(err)
}

func (m SubscribersMicroservice) Unsubscribe(subscriber models.Subscription) error {
	_, err := m.subscribersClient.Unsubscribe(context.Background(), &userProto.UserAuthorPair{
		UserId:   subscriber.SubscriberID,
		AuthorId: subscriber.AuthorID,
	})

	return tracerr.Wrap(err)
}
