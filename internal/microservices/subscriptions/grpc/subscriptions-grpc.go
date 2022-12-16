package grpcSubscriptions

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscriptions/protobuf"
	usersProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
)

type SubscriptionsService struct {
	subscriptionsRepo domain.SubscriptionsRepository
	protobuf.UnimplementedSubscriptionsServer
}

func New(r domain.SubscriptionsRepository) protobuf.SubscriptionsServer {
	return &SubscriptionsService{
		subscriptionsRepo: r,
	}
}

func ConvertToProto(s models.AuthorSubscription) *protobuf.AuthorSubscription {
	return &protobuf.AuthorSubscription{
		ID:           s.ID,
		AuthorID:     s.AuthorID,
		Img:          s.Img,
		Tier:         s.Tier,
		Title:        s.Title,
		Text:         s.Text,
		Price:        s.Price,
		AuthorName:   s.AuthorName,
		AuthorAvatar: s.AuthorAvatar,
	}
}

func ConvertToModel(s *protobuf.AuthorSubscription) models.AuthorSubscription {
	return models.AuthorSubscription{
		ID:           s.GetID(),
		AuthorID:     s.GetAuthorID(),
		Img:          s.GetImg(),
		Tier:         s.GetTier(),
		Title:        s.GetTitle(),
		Text:         s.GetText(),
		Price:        s.GetPrice(),
		AuthorName:   s.GetAuthorName(),
		AuthorAvatar: s.GetAuthorAvatar(),
	}
}

func (s SubscriptionsService) GetSubscriptionsByUserID(_ context.Context, id *usersProto.UserID) (*protobuf.SubArray, error) {
	sub, err := s.subscriptionsRepo.GetSubscriptionsByUserID(id.GetUserId())
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := make([]*protobuf.AuthorSubscription, 0)

	for _, subscription := range sub {
		result = append(result, ConvertToProto(subscription))
	}

	return &protobuf.SubArray{Subscriptions: result}, nil
}

func (s SubscriptionsService) GetSubscriptionsByAuthorID(_ context.Context, id *usersProto.UserID) (*protobuf.SubArray, error) {
	sub, err := s.subscriptionsRepo.GetSubscriptionsByAuthorID(id.GetUserId())
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := make([]*protobuf.AuthorSubscription, 0)

	for _, subscription := range sub {
		result = append(result, ConvertToProto(subscription))
	}

	return &protobuf.SubArray{Subscriptions: result}, nil
}

func (s SubscriptionsService) GetSubscriptionByID(_ context.Context, id *protobuf.AuthorSubscriptionID) (*protobuf.AuthorSubscription, error) {
	sub, err := s.subscriptionsRepo.GetSubscriptionByID(id.GetID())
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ConvertToProto(sub), nil
}

func (s SubscriptionsService) GetSubscriptionByUserAndAuthorID(_ context.Context, pair *usersProto.UserAuthorPair) (*protobuf.AuthorSubscription, error) {
	sub, err := s.subscriptionsRepo.GetSubscriptionByUserAndAuthorID(pair.GetUserId(), pair.GetAuthorId())
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ConvertToProto(sub), nil
}

func (s SubscriptionsService) AddSubscription(_ context.Context, sub *protobuf.AuthorSubscription) (*protobuf.AuthorSubscriptionID, error) {
	subscription := ConvertToModel(sub)
	id, err := s.subscriptionsRepo.AddSubscription(subscription)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.AuthorSubscriptionID{ID: id}, nil
}

func (s SubscriptionsService) UpdateSubscription(_ context.Context, sub *protobuf.AuthorSubscription) (*emptypb.Empty, error) {
	err := s.subscriptionsRepo.UpdateSubscription(ConvertToModel(sub))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s SubscriptionsService) DeleteSubscription(_ context.Context, sub *protobuf.AuthorSubscriptionID) (*emptypb.Empty, error) {
	err := s.subscriptionsRepo.DeleteSubscription(sub.GetID())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
