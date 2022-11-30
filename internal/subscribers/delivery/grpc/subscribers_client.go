package subscribersClient

import (
	"context"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/protobuf"

	grpcUsers "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/grpc"
	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type subscribersServiceClient struct {
	subscribersClient protobuf.SubscribersClient
}

func NewSubcribersClient(subscribersClient protobuf.SubscribersClient) domain.SubscribersService {
	return subscribersServiceClient{
		subscribersClient: subscribersClient,
	}
}

func (ss subscribersServiceClient) GetSubscribers(userID uint64) ([]models.User, error) {
	subscribers, err := ss.subscribersClient.GetSubscribers(context.Background(), &userProto.UserID{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	res := make([]models.User, 0)
	for _, u := range subscribers.GetUsers() {
		res = append(res, grpcUsers.ConvertToModel(u))
	}

	return res, nil
}

func (ss subscribersServiceClient) Subscribe(subscriber models.Subscription) error {
	_, err := ss.subscribersClient.Subscribe(context.Background(), &protobuf.Subscriber{
		AuthorID:             subscriber.AuthorID,
		SubscriberID:         subscriber.SubscriberID,
		AuthorSubscriptionID: subscriber.AuthorSubscriptionID,
	})

	return err
}

func (ss subscribersServiceClient) Unsubscribe(subscriber models.Subscription) error {
	_, err := ss.subscribersClient.Unsubscribe(context.Background(), &userProto.UserAuthorPair{
		UserId:   subscriber.SubscriberID,
		AuthorId: subscriber.AuthorID,
	})

	return err
}
