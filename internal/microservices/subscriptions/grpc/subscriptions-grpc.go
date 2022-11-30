package grpcSubscriptions

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscriptions/protobuf"
	usersProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
)

type SubscriptionsService struct {
	subscriptionsRepo domain.SubscriptionsRepository
	userRepo          domain.UsersRepository
	imagesUseCase     domain.ImageUseCase
	protobuf.UnimplementedSubscriptionsServer
}

func NewSubscriptionsService(r domain.SubscriptionsRepository, i domain.ImageUseCase, u domain.UsersRepository) protobuf.SubscriptionsServer {
	return &SubscriptionsService{
		subscriptionsRepo: r,
		userRepo:          u,
		imagesUseCase:     i,
	}
}

func convertToProto(s models.AuthorSubscription) *protobuf.AuthorSubscription {
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

func convertToModel(s *protobuf.AuthorSubscription) models.AuthorSubscription {
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
	if err != nil {
		return nil, err
	}

	result := make([]*protobuf.AuthorSubscription, 0)

	for _, subscription := range sub {
		result = append(result, convertToProto(subscription))
	}

	return &protobuf.SubArray{Subscriptions: result}, nil
}

func (s SubscriptionsService) GetAuthorSubscriptionsByAuthorID(_ context.Context, id *usersProto.UserID) (*protobuf.SubArray, error) {
	sub, err := s.subscriptionsRepo.GetSubscriptionsByAuthorID(id.GetUserId())
	if err != nil {
		return nil, err
	}

	result := make([]*protobuf.AuthorSubscription, 0)

	for i, subscription := range sub {
		if sub[i].Img, err = s.imagesUseCase.GetImage(subscription.Img); err != nil {
			return nil, err
		}
		result = append(result, convertToProto(sub[i]))
	}

	return &protobuf.SubArray{Subscriptions: result}, nil
}

func (s SubscriptionsService) GetAuthorSubscriptionByID(_ context.Context, id *protobuf.AuthorSubscriptionID) (*protobuf.AuthorSubscription, error) {
	sub, err := s.subscriptionsRepo.GetSubscriptionByID(id.GetID())
	if err != nil {
		return nil, err
	}

	if sub.Img, err = s.imagesUseCase.GetImage(sub.Img); err != nil {
		return nil, err
	}

	return convertToProto(sub), nil
}

func (s SubscriptionsService) GetSubscriptionByUserAndAuthorID(_ context.Context, pair *usersProto.UserAuthorPair) (*protobuf.AuthorSubscription, error) {
	sub, err := s.subscriptionsRepo.GetSubscriptionByUserAndAuthorID(pair.GetUserId(), pair.GetAuthorId())
	if err != nil {
		return nil, err
	}

	return convertToProto(sub), nil
}

func (s SubscriptionsService) AddSubscription(_ context.Context, sub *protobuf.AuthorSubscription) (*protobuf.AuthorSubscriptionID, error) {
	subscription := convertToModel(sub)
	id, err := s.subscriptionsRepo.AddSubscription(subscription)
	if err != nil {
		return nil, err
	}

	return &protobuf.AuthorSubscriptionID{ID: id}, nil
}

func (s SubscriptionsService) UpdateSubscription(_ context.Context, sub *protobuf.AuthorSubscription) (*emptypb.Empty, error) {
	err := s.subscriptionsRepo.UpdateSubscription(convertToModel(sub))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s SubscriptionsService) DeleteSubscription(_ context.Context, sub *protobuf.AuthorSubscriptionID) (*emptypb.Empty, error) {
	err := s.subscriptionsRepo.DeleteSubscription(sub.GetID())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
