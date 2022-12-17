package grpcSubscribers

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/protobuf"
	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	mock_domain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConvertToModel(t *testing.T) {
	input := &protobuf.Subscriber{
		AuthorID:             1,
		SubscriberID:         2,
		AuthorSubscriptionID: 3,
	}

	expected := models.Subscription{
		AuthorID:             1,
		SubscriberID:         2,
		AuthorSubscriptionID: 3,
	}

	actual := ConvertToModel(input)

	assert.Equal(t, expected, actual)
}

func TestSubscribersService_GetSubscribers(t *testing.T) {
	type mockBehaviorGetSubscribers func(r *mock_domain.MockSubscribersRepository, id uint64)

	tests := []struct {
		name          string
		input         uint64
		mockBehavior  mockBehaviorGetSubscribers
		expected      *userProto.UserIDs
		expectedError string
	}{
		{
			name:  "OK",
			input: 1,
			mockBehavior: func(r *mock_domain.MockSubscribersRepository, id uint64) {
				r.EXPECT().GetSubscribers(id).Return([]uint64{1, 2, 3}, nil)
			},
			expected: &userProto.UserIDs{
				Ids: []*userProto.UserID{
					{UserId: 1},
					{UserId: 2},
					{UserId: 3},
				},
			},
		},
		{
			name:  "Error",
			input: 1,
			mockBehavior: func(r *mock_domain.MockSubscribersRepository, id uint64) {
				r.EXPECT().GetSubscribers(id).Return([]uint64{}, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockSubscribersRepository(ctrl)
			test.mockBehavior(repo, test.input)

			s := New(repo)

			_, err := s.GetSubscribers(context.Background(), &userProto.UserID{UserId: test.input})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

// func TestSubscribersService_Subscribe(t *testing.T) {
// 	type mockBehaviorSubscribe func(r *mock_domain.MockSubscribersRepository, sub models.Subscription)
//
// 	tests := []struct {
// 		name          string
// 		sub           models.Subscription
// 		mockBehavior  mockBehaviorSubscribe
// 		expected      *protobuf.Subscriber
// 		expectedError string
// 	}{
// 		{
// 			name: "OK",
// 			sub: models.Subscription{
// 				AuthorID:             1,
// 				SubscriberID:         2,
// 				AuthorSubscriptionID: 3,
// 			},
// 			mockBehavior: func(r *mock_domain.MockSubscribersRepository, sub models.Subscription) {
// 				r.EXPECT().PayAndSubscribe(sub).Return(nil)
// 			},
// 		},
// 		{
// 			name: "Error",
// 			sub: models.Subscription{
// 				AuthorID:             1,
// 				SubscriberID:         2,
// 				AuthorSubscriptionID: 3,
// 			},
// 			mockBehavior: func(r *mock_domain.MockSubscribersRepository, sub models.Subscription) {
// 				r.EXPECT().PayAndSubscribe(sub).Return(domain.ErrInternal)
// 			},
// 			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
// 		},
// 	}
//
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
//
// 			repo := mock_domain.NewMockSubscribersRepository(ctrl)
// 			test.mockBehavior(repo, test.sub)
//
// 			s := New(repo)
//
// 			_, err := s.Subscribe(context.Background(), &protobuf.Payment{
// 				ToID:   test.sub.AuthorID,
// 				FromID: test.sub.SubscriberID,
// 				SubID:  test.sub.AuthorSubscriptionID,
// 			})
// 			if err != nil {
// 				assert.Equal(t, test.expectedError, err.Error())
// 			}
// 		})
// 	}
// }

func TestSubscribersService_Unsubscribe(t *testing.T) {
	type mockBehaviorUnsubscribe func(r *mock_domain.MockSubscribersRepository, userID, authorID uint64)

	tests := []struct {
		name          string
		userID        uint64
		authorID      uint64
		mockBehavior  mockBehaviorUnsubscribe
		expectedError string
	}{
		{
			name:     "OK",
			userID:   1,
			authorID: 2,
			mockBehavior: func(r *mock_domain.MockSubscribersRepository, userID, authorID uint64) {
				r.EXPECT().Unsubscribe(userID, authorID).Return(nil)
			},
		},
		{
			name:     "Error",
			userID:   1,
			authorID: 2,
			mockBehavior: func(r *mock_domain.MockSubscribersRepository, userID, authorID uint64) {
				r.EXPECT().Unsubscribe(userID, authorID).Return(domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockSubscribersRepository(ctrl)
			test.mockBehavior(repo, test.userID, test.authorID)

			s := New(repo)

			_, err := s.Unsubscribe(context.Background(), &userProto.UserAuthorPair{
				UserId:   test.userID,
				AuthorId: test.authorID,
			})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}
