package authMicroservice

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestAuthClient_CreateSession(t *testing.T) {
	type MockCreateSession func(s *mockDomain.MockAuthClient, c context.Context, in *protobuf.Session, opts ...grpc.CallOption)

	tests := []struct {
		name     string
		userID   uint64
		mock     MockCreateSession
		response string
		err      error
	}{
		{
			name:   "OK",
			userID: 1,
			mock: func(s *mockDomain.MockAuthClient, c context.Context, in *protobuf.Session, opts ...grpc.CallOption) {
				s.EXPECT().CreateSession(c, in).Return(&protobuf.SessionID{SessionId: "sessionID"}, nil)
			},
			response: "sessionID",
			err:      nil,
		},
		{
			name:   "Err",
			userID: 123,
			mock: func(s *mockDomain.MockAuthClient, c context.Context, in *protobuf.Session, opts ...grpc.CallOption) {
				s.EXPECT().CreateSession(c, in).Return(&protobuf.SessionID{}, errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockAuthClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.Session{UserId: test.userID})
			authClient := AuthMicroservice{authClient: mock}
			sessionID, err := authClient.CreateSession(test.userID)

			assert.Equal(t, sessionID, test.response)
			assert.Equal(t, err, test.err)
		})
	}
}

func TestAuthClient_DeleteBySessionID(t *testing.T) {
	type MockDeleteBySessionID func(s *mockDomain.MockAuthClient, c context.Context, in *protobuf.SessionID, opts ...grpc.CallOption)

	tests := []struct {
		name     string
		session  string
		mock     MockDeleteBySessionID
		response string
		err      error
	}{
		{
			name:    "OK",
			session: "sessionID",
			mock: func(s *mockDomain.MockAuthClient, c context.Context, in *protobuf.SessionID, opts ...grpc.CallOption) {
				s.EXPECT().DeleteBySessionID(c, in).Return(&emptypb.Empty{}, nil)
			},
			response: "",
			err:      nil,
		},
		{
			name:    "Err",
			session: "sessionID",
			mock: func(s *mockDomain.MockAuthClient, c context.Context, in *protobuf.SessionID, opts ...grpc.CallOption) {
				s.EXPECT().DeleteBySessionID(c, in).Return(&emptypb.Empty{}, errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockAuthClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.SessionID{SessionId: test.session})
			authClient := AuthMicroservice{authClient: mock}
			err := authClient.DeleteBySessionID(test.session)

			assert.Equal(t, err, test.err)
		})
	}
}

func TestAuthClient_GetUserIDBySessionID(t *testing.T) {
	type MockGetUserIDBySessionID func(s *mockDomain.MockAuthClient, c context.Context, in *protobuf.SessionID, opts ...grpc.CallOption)

	tests := []struct {
		name     string
		session  string
		mock     MockGetUserIDBySessionID
		response models.Cookie
		err      error
	}{
		{
			name:    "OK",
			session: "sessionID",
			mock: func(s *mockDomain.MockAuthClient, c context.Context, in *protobuf.SessionID, opts ...grpc.CallOption) {
				s.EXPECT().GetBySessionID(c, in).Return(&protobuf.Session{
					SessionId: "sessionID",
					UserId:    1,
					Expires:   timestamppb.New(time.Time{}),
				}, nil)
			},
			response: models.Cookie{
				UserID: 1,
				Value:  "sessionID",
			},
			err: nil,
		},
		{
			name:    "Err",
			session: "sessionID",
			mock: func(s *mockDomain.MockAuthClient, c context.Context, in *protobuf.SessionID, opts ...grpc.CallOption) {
				s.EXPECT().GetBySessionID(c, in).Return(&protobuf.Session{}, errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockAuthClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.SessionID{SessionId: test.session})
			authClient := AuthMicroservice{authClient: mock}
			userID, err := authClient.GetBySessionID(test.session)

			assert.Equal(t, userID, test.response)
			assert.Equal(t, err, test.err)
		})
	}
}
