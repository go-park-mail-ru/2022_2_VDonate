package subscriptionsMicroservice

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscriptions/protobuf"
	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestSubscriptionsClient_GetSubscriptionsByUserID(t *testing.T) {
	type mockGetSub func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *userProto.UserID, opts ...interface{})

	tests := []struct {
		name     string
		authorID uint64
		mock     mockGetSub
		response []models.AuthorSubscription
		err      error
	}{
		{
			name:     "OK",
			authorID: 1,
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *userProto.UserID, opts ...interface{}) {
				r.EXPECT().GetSubscriptionsByUserID(c, in).Return(&protobuf.SubArray{
					Subscriptions: []*protobuf.AuthorSubscription{
						{
							ID:       1,
							AuthorID: 2,
						},
					},
				}, nil)
			},
			response: []models.AuthorSubscription{
				{
					ID:       1,
					AuthorID: 2,
				},
			},
			err: nil,
		},
		{
			name:     "Error",
			authorID: 1,
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *userProto.UserID, opts ...interface{}) {
				r.EXPECT().GetSubscriptionsByUserID(c, in).Return(nil, errors.New("error"))
			},
			response: nil,
			err:      errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockSubscriptionsClient(ctrl)

			test.mock(mock, context.Background(), &userProto.UserID{
				UserId: test.authorID,
			})

			client := New(mock)
			res, err := client.GetSubscriptionsByUserID(test.authorID)
			require.Equal(t, test.response, res)
			require.Equal(t, test.err, err)
		})
	}
}

func TestSubscriptionsClient_GetSubscriptionsByAuthorID(t *testing.T) {
	type mockGetSub func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *userProto.UserID, opts ...interface{})

	tests := []struct {
		name     string
		authorID uint64
		mock     mockGetSub
		response []models.AuthorSubscription
		err      error
	}{
		{
			name:     "OK",
			authorID: 1,
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *userProto.UserID, opts ...interface{}) {
				r.EXPECT().GetSubscriptionsByAuthorID(c, in).Return(&protobuf.SubArray{
					Subscriptions: []*protobuf.AuthorSubscription{
						{
							ID:       1,
							AuthorID: 2,
						},
					},
				}, nil)
			},
			response: []models.AuthorSubscription{
				{
					ID:       1,
					AuthorID: 2,
				},
			},
			err: nil,
		},
		{
			name:     "Error",
			authorID: 1,
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *userProto.UserID, opts ...interface{}) {
				r.EXPECT().GetSubscriptionsByAuthorID(c, in).Return(nil, errors.New("error"))
			},
			response: nil,
			err:      errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockSubscriptionsClient(ctrl)

			test.mock(mock, context.Background(), &userProto.UserID{
				UserId: test.authorID,
			})

			client := New(mock)
			res, err := client.GetSubscriptionsByAuthorID(test.authorID)
			require.Equal(t, test.response, res)
			require.Equal(t, test.err, err)
		})
	}
}

func TestSubscriptionsClient_GetSubscriptionByID(t *testing.T) {
	type mockGetSub func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscriptionID, opts ...interface{})

	tests := []struct {
		name     string
		subID    uint64
		mock     mockGetSub
		response models.AuthorSubscription
		err      error
	}{
		{
			name:  "OK",
			subID: 1,
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscriptionID, opts ...interface{}) {
				r.EXPECT().GetSubscriptionByID(c, in).Return(&protobuf.AuthorSubscription{
					ID:       1,
					AuthorID: 2,
				}, nil)
			},
			response: models.AuthorSubscription{
				ID:       1,
				AuthorID: 2,
			},
			err: nil,
		},
		{
			name:  "Error",
			subID: 1,
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscriptionID, opts ...interface{}) {
				r.EXPECT().GetSubscriptionByID(c, in).Return(nil, errors.New("error"))
			},
			response: models.AuthorSubscription{},
			err:      errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockSubscriptionsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.AuthorSubscriptionID{
				ID: test.subID,
			})

			client := New(mock)
			res, err := client.GetSubscriptionByID(test.subID)
			require.Equal(t, test.response, res)
			require.Equal(t, test.err, err)
		})
	}
}

func TestSubscriptionsClient_GetSubscriptionByUserAndAuthorID(t *testing.T) {
	type mockGetSub func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *userProto.UserAuthorPair, opts ...interface{})

	tests := []struct {
		name     string
		userID   uint64
		authorID uint64
		mock     mockGetSub
		response models.AuthorSubscription
		err      error
	}{
		{
			name:     "OK",
			userID:   1,
			authorID: 2,
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *userProto.UserAuthorPair, opts ...interface{}) {
				r.EXPECT().GetSubscriptionByUserAndAuthorID(c, in).Return(&protobuf.AuthorSubscription{
					ID:       1,
					AuthorID: 2,
				}, nil)
			},
			response: models.AuthorSubscription{
				ID:       1,
				AuthorID: 2,
			},
			err: nil,
		},
		{
			name:     "Error",
			userID:   1,
			authorID: 2,
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *userProto.UserAuthorPair, opts ...interface{}) {
				r.EXPECT().GetSubscriptionByUserAndAuthorID(c, in).Return(nil, errors.New("error"))
			},
			response: models.AuthorSubscription{},
			err:      errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockSubscriptionsClient(ctrl)

			test.mock(mock, context.Background(), &userProto.UserAuthorPair{
				UserId:   test.userID,
				AuthorId: test.authorID,
			})

			client := New(mock)
			res, err := client.GetSubscriptionByUserAndAuthorID(test.userID, test.authorID)
			require.Equal(t, test.response, res)
			require.Equal(t, test.err, err)
		})
	}
}

func TestSubscriptionsClient_AddSubscription(t *testing.T) {
	type mockAddSub func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscription, opts ...interface{})

	tests := []struct {
		name     string
		sub      models.AuthorSubscription
		mock     mockAddSub
		response uint64
		err      error
	}{
		{
			name: "OK",
			sub: models.AuthorSubscription{
				ID:       1,
				AuthorID: 2,
			},
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscription, opts ...interface{}) {
				r.EXPECT().AddSubscription(c, in).Return(&protobuf.AuthorSubscriptionID{
					ID: 1,
				}, nil)
			},
			response: 1,
			err:      nil,
		},
		{
			name: "Error",
			sub: models.AuthorSubscription{
				ID:       1,
				AuthorID: 2,
			},
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscription, opts ...interface{}) {
				r.EXPECT().AddSubscription(c, in).Return(nil, errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockSubscriptionsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.AuthorSubscription{
				ID:       test.sub.ID,
				AuthorID: test.sub.AuthorID,
			})

			client := New(mock)
			res, err := client.AddSubscription(test.sub)
			require.Equal(t, test.response, res)
			require.Equal(t, test.err, err)
		})
	}
}

func TestSubscriptionsClient_UpdateSubscription(t *testing.T) {
	type mockUpdateSub func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscription, opts ...interface{})

	tests := []struct {
		name     string
		sub      models.AuthorSubscription
		mock     mockUpdateSub
		response uint64
		err      error
	}{
		{
			name: "OK",
			sub: models.AuthorSubscription{
				ID:       1,
				AuthorID: 2,
			},
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscription, opts ...interface{}) {
				r.EXPECT().UpdateSubscription(c, in).Return(&emptypb.Empty{}, nil)
			},
			response: 1,
			err:      nil,
		},
		{
			name: "Error",
			sub: models.AuthorSubscription{
				ID:       1,
				AuthorID: 2,
			},
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscription, opts ...interface{}) {
				r.EXPECT().UpdateSubscription(c, in).Return(nil, errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockSubscriptionsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.AuthorSubscription{
				ID:       test.sub.ID,
				AuthorID: test.sub.AuthorID,
			})

			client := New(mock)
			err := client.UpdateSubscription(test.sub)
			require.Equal(t, test.err, err)
		})
	}
}

func TestSubscriptionsClient_DeleteSubscription(t *testing.T) {
	type mockDeleteSub func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscriptionID, opts ...interface{})

	tests := []struct {
		name     string
		subID    uint64
		mock     mockDeleteSub
		response uint64
		err      error
	}{
		{
			name:  "OK",
			subID: 1,
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscriptionID, opts ...interface{}) {
				r.EXPECT().DeleteSubscription(c, in).Return(&emptypb.Empty{}, nil)
			},
			response: 1,
			err:      nil,
		},
		{
			name:  "Error",
			subID: 1,
			mock: func(r *mockDomain.MockSubscriptionsClient, c context.Context, in *protobuf.AuthorSubscriptionID, opts ...interface{}) {
				r.EXPECT().DeleteSubscription(c, in).Return(nil, errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockSubscriptionsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.AuthorSubscriptionID{
				ID: test.subID,
			})

			client := New(mock)
			err := client.DeleteSubscription(test.subID)
			require.Equal(t, test.err, err)
		})
	}
}
