package subscribersMicroservice

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestSubscribersClient_GetSubscribers(t *testing.T) {
	type mockGetSubscribers func(r *mockDomain.MockSubscribersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption)

	tests := []struct {
		name     string
		authorID uint64
		mock     mockGetSubscribers
		response []uint64
		err      error
	}{
		{
			name: "OK",
			mock: func(r *mockDomain.MockSubscribersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				r.EXPECT().GetSubscribers(c, in).Return(&userProto.UserIDs{
					Ids: []*userProto.UserID{
						{
							UserId: 1,
						},
					},
				}, nil)
			},
			response: []uint64{1},
			err:      nil,
		},
		{
			name: "Error",
			mock: func(r *mockDomain.MockSubscribersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				r.EXPECT().GetSubscribers(c, in).Return(nil, status.Error(codes.Canceled, "canceled"))
			},
			err: status.Error(codes.Canceled, "canceled"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockSubscribersClient(ctrl)

			test.mock(mock, context.Background(), &userProto.UserID{
				UserId: test.authorID,
			})

			client := New(mock)
			response, err := client.GetSubscribers(test.authorID)
			assert.Equal(t, test.response, response)
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
		})
	}
}

// func TestSubscribersClient_Subscribe(t *testing.T) {
// 	type mockSubscribe func(r *mockDomain.MockSubscribersClient, c context.Context, in *protobuf.Subscriber, opts ...grpc.CallOption)
//
// 	tests := []struct {
// 		name         string
// 		subscription models.Subscription
// 		mock         mockSubscribe
// 		response     bool
// 		err          error
// 	}{
// 		{
// 			name: "OK",
// 			subscription: models.Subscription{
// 				AuthorID:     1,
// 				SubscriberID: 2,
// 			},
// 			mock: func(r *mockDomain.MockSubscribersClient, c context.Context, in *protobuf.Subscriber, opts ...grpc.CallOption) {
// 				r.EXPECT().Subscribe(c, in).Return(&emptypb.Empty{}, nil)
// 			},
// 			response: true,
// 			err:      nil,
// 		},
// 		{
// 			name: "Error",
// 			mock: func(r *mockDomain.MockSubscribersClient, c context.Context, in *protobuf.Subscriber, opts ...grpc.CallOption) {
// 				r.EXPECT().Subscribe(c, in).Return(nil, status.Error(codes.Canceled, "canceled"))
// 			},
// 			err: status.Error(codes.Canceled, "canceled"),
// 		},
// 	}
//
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
// 			mock := mockDomain.NewMockSubscribersClient(ctrl)
//
// 			test.mock(mock, context.Background(), &protobuf.Subscriber{
// 				AuthorID:     test.subscription.AuthorID,
// 				SubscriberID: test.subscription.SubscriberID,
// 			})
//
// 			client := New(mock)
// 			client.Subscribe(test.subscription)
// 		})
// 	}
// }

func TestSubscribersClient_Unsubscribe(t *testing.T) {
	type mockUnsubscribe func(r *mockDomain.MockSubscribersClient, c context.Context, in *userProto.UserAuthorPair, opts ...grpc.CallOption)

	tests := []struct {
		name         string
		subscription models.Subscription
		mock         mockUnsubscribe
		response     bool
		err          error
	}{
		{
			name: "OK",
			subscription: models.Subscription{
				AuthorID:     1,
				SubscriberID: 2,
			},
			mock: func(r *mockDomain.MockSubscribersClient, c context.Context, in *userProto.UserAuthorPair, opts ...grpc.CallOption) {
				r.EXPECT().Unsubscribe(c, in).Return(&emptypb.Empty{}, nil)
			},
			response: true,
			err:      nil,
		},
		{
			name: "Error",
			mock: func(r *mockDomain.MockSubscribersClient, c context.Context, in *userProto.UserAuthorPair, opts ...grpc.CallOption) {
				r.EXPECT().Unsubscribe(c, in).Return(nil, status.Error(codes.Canceled, "canceled"))
			},
			err: status.Error(codes.Canceled, "canceled"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockSubscribersClient(ctrl)

			test.mock(mock, context.Background(), &userProto.UserAuthorPair{
				UserId:   test.subscription.SubscriberID,
				AuthorId: test.subscription.AuthorID,
			})

			client := New(mock)
			err := client.Unsubscribe(test.subscription)
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
		})
	}
}
