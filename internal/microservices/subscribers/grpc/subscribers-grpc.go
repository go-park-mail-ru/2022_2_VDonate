package grpcSubscribers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/protobuf"

	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
)

type SubscribersService struct {
	subscribersRepo domain.SubscribersRepository
	protobuf.UnimplementedSubscribersServer
}

func ConvertToModel(s *protobuf.Subscriber) models.Subscription {
	return models.Subscription{
		AuthorID:             s.GetAuthorID(),
		SubscriberID:         s.GetSubscriberID(),
		AuthorSubscriptionID: s.GetAuthorSubscriptionID(),
	}
}

func New(s domain.SubscribersRepository) protobuf.SubscribersServer {
	return &SubscribersService{
		subscribersRepo: s,
	}
}

func (s SubscribersService) GetSubscribers(_ context.Context, id *userProto.UserID) (*userProto.UserIDs, error) {
	sub, err := s.subscribersRepo.GetSubscribers(id.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := make([]*userProto.UserID, 0)

	for _, userID := range sub {
		// Notion: if there is an error while getting user, skip it
		result = append(result, &userProto.UserID{UserId: userID})
	}

	return &userProto.UserIDs{Ids: result}, err
}

func (s SubscribersService) Follow(_ context.Context, pair *userProto.UserAuthorPair) (*emptypb.Empty, error) {
	err := s.subscribersRepo.Follow(pair.GetUserId(), pair.GetAuthorId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s SubscribersService) Subscribe(_ context.Context, payment *protobuf.Payment) (*emptypb.Empty, error) {
	err := s.subscribersRepo.PayAndSubscribe(models.Payment{
		ID:     payment.GetID(),
		ToID:   payment.GetToID(),
		FromID: payment.GetFromID(),
		SubID:  payment.GetSubID(),
		Price:  payment.GetPrice(),
		Status: payment.GetStatus(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s SubscribersService) Unsubscribe(_ context.Context, pair *userProto.UserAuthorPair) (*emptypb.Empty, error) {
	err := s.subscribersRepo.Unsubscribe(pair.GetUserId(), pair.GetAuthorId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s SubscribersService) ChangePaymentStatus(_ context.Context, statusWithID *protobuf.StatusAndID) (*emptypb.Empty, error) {
	err := s.subscribersRepo.UpdateStatus(statusWithID.GetStatus(), statusWithID.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
