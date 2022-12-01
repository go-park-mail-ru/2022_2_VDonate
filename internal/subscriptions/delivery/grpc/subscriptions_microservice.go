package subscriptionsMicroservice

import (
	"context"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	grpcSubscriptions "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscriptions/grpc"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscriptions/protobuf"
	usersProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type SubscriptionsMicroservice struct {
	client protobuf.SubscriptionsClient
}

func New(c protobuf.SubscriptionsClient) domain.SubscriptionMicroservice {
	return &SubscriptionsMicroservice{
		client: c,
	}
}

func (m SubscriptionsMicroservice) GetSubscriptionsByUserID(userID uint64) ([]models.AuthorSubscription, error) {
	s, err := m.client.GetSubscriptionsByUserID(context.Background(), &usersProto.UserID{UserId: userID})
	if err != nil {
		return nil, err
	}

	result := make([]models.AuthorSubscription, 0)
	for _, subscription := range s.GetSubscriptions() {
		result = append(result, grpcSubscriptions.ConvertToModel(subscription))
	}

	return result, nil
}

func (m SubscriptionsMicroservice) GetSubscriptionsByAuthorID(authorID uint64) ([]models.AuthorSubscription, error) {
	s, err := m.client.GetSubscriptionsByAuthorID(context.Background(), &usersProto.UserID{UserId: authorID})
	if err != nil {
		return nil, err
	}

	result := make([]models.AuthorSubscription, 0)
	for _, subscription := range s.GetSubscriptions() {
		result = append(result, grpcSubscriptions.ConvertToModel(subscription))
	}

	return result, nil
}

func (m SubscriptionsMicroservice) GetSubscriptionByID(id uint64) (models.AuthorSubscription, error) {
	s, err := m.client.GetSubscriptionByID(context.Background(), &protobuf.AuthorSubscriptionID{ID: id})
	if err != nil {
		return models.AuthorSubscription{}, err
	}

	return grpcSubscriptions.ConvertToModel(s), nil
}

func (m SubscriptionsMicroservice) GetSubscriptionByUserAndAuthorID(userID, authorID uint64) (models.AuthorSubscription, error) {
	s, err := m.client.GetSubscriptionByUserAndAuthorID(context.Background(), &usersProto.UserAuthorPair{
		UserId:   userID,
		AuthorId: authorID,
	})
	if err != nil {
		return models.AuthorSubscription{}, err
	}

	return grpcSubscriptions.ConvertToModel(s), nil
}

func (m SubscriptionsMicroservice) AddSubscription(subscription models.AuthorSubscription) (uint64, error) {
	id, err := m.client.AddSubscription(context.Background(), grpcSubscriptions.ConvertToProto(subscription))
	if err != nil {
		return 0, err
	}

	return id.GetID(), nil
}

func (m SubscriptionsMicroservice) UpdateSubscription(subscription models.AuthorSubscription) error {
	_, err := m.client.UpdateSubscription(context.Background(), grpcSubscriptions.ConvertToProto(subscription))

	return err
}

func (m SubscriptionsMicroservice) DeleteSubscription(id uint64) error {
	_, err := m.client.DeleteSubscription(context.Background(), &protobuf.AuthorSubscriptionID{
		ID: id,
	})

	return err
}
