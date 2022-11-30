package grpcSubscribers

import (
	"context"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"google.golang.org/protobuf/types/known/emptypb"

	grpcUsers "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/grpc"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/protobuf"

	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
)

type SubscribersService struct {
	subscribersRepo domain.SubscribersRepository
	usersRepo       domain.UsersRepository
	protobuf.UnimplementedSubscribersServer
}

func ConvertToModel(s *protobuf.Subscriber) models.Subscription {
	return models.Subscription{
		AuthorID:             s.GetAuthorID(),
		SubscriberID:         s.GetSubscriberID(),
		AuthorSubscriptionID: s.GetAuthorSubscriptionID(),
	}
}

func NewSubscribersService(s domain.SubscribersRepository, u domain.UsersRepository) protobuf.SubscribersServer {
	return &SubscribersService{
		subscribersRepo: s,
		usersRepo:       u,
	}
}

func (s SubscribersService) GetSubscribers(_ context.Context, id *userProto.UserID) (*userProto.UsersArray, error) {
	sub, err := s.subscribersRepo.GetSubscribers(id.GetUserId())
	if err != nil {
		return nil, err
	}

	result := make([]*userProto.User, 0)

	for _, userID := range sub {
		// Notion: if there is an error while getting user, skip it
		user, _ := s.usersRepo.GetByID(userID)
		result = append(result, grpcUsers.ConvertToProto(user))
	}

	return &userProto.UsersArray{Users: result}, err
}

func (s SubscribersService) Subscribe(_ context.Context, sub *protobuf.Subscriber) (*emptypb.Empty, error) {
	err := s.subscribersRepo.Subscribe(ConvertToModel(sub))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s SubscribersService) Unsubscribe(_ context.Context, pair *userProto.UserAuthorPair) (*emptypb.Empty, error) {
	err := s.subscribersRepo.Unsubscribe(pair.GetUserId(), pair.GetAuthorId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
