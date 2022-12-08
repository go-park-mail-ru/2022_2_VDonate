package grpcAuth

import (
	"context"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestAuth_CreateSession(t *testing.T) {
	type mockBehaviorCreateSession func(r *mockDomain.MockAuthRepository, id uint64)

	cc := func(id uint64) models.Cookie {
		return models.Cookie{
			UserID:  id,
			Value:   "sessionID",
			Expires: time.Unix(0, 0),
		}
	}

	tests := []struct {
		name                      string
		input                     uint64
		mockBehaviorCreateSession mockBehaviorCreateSession
		expected                  string
		expectedErr               string
	}{
		{
			name:  "OK",
			input: 1,
			mockBehaviorCreateSession: func(r *mockDomain.MockAuthRepository, id uint64) {
				r.EXPECT().CreateSession(cc(id)).Return(models.Cookie{
					Value:   "sessionID",
					UserID:  id,
					Expires: time.Unix(0, 0),
				}, nil)
			},
			expected: "sessionID",
		},
		{
			name:  "Error",
			input: 1,
			mockBehaviorCreateSession: func(r *mockDomain.MockAuthRepository, id uint64) {
				r.EXPECT().CreateSession(cc(id)).Return(models.Cookie{}, domain.ErrInternal)
			},
			expectedErr: "server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockDomain.NewMockAuthRepository(ctrl)
			test.mockBehaviorCreateSession(repo, test.input)

			s := NewWithCookieCreator(
				repo,
				cc,
			)

			c := cc(test.input)

			id, err := s.CreateSession(context.Background(), &protobuf.Session{
				SessionId: c.Value,
				UserId:    c.UserID,
				Expires:   timestamppb.New(c.Expires),
			})

			if err != nil {
				require.Equal(t, test.expectedErr, err.Error())
			} else {
				require.Equal(t, test.expected, id.GetSessionId())
			}
		})
	}
}

func TestAuth_DeleteBySessionID(t *testing.T) {
	type mockBehaviorDeleteSession func(r *mockDomain.MockAuthRepository, id string)

	tests := []struct {
		name                      string
		input                     string
		mockBehaviorDeleteSession mockBehaviorDeleteSession
		expectedErr               string
	}{
		{
			name:  "OK",
			input: "sessionID",
			mockBehaviorDeleteSession: func(r *mockDomain.MockAuthRepository, id string) {
				r.EXPECT().DeleteBySessionID(id).Return(nil)
			},
		},
		{
			name:  "Error",
			input: "sessionID",
			mockBehaviorDeleteSession: func(r *mockDomain.MockAuthRepository, id string) {
				r.EXPECT().DeleteBySessionID(id).Return(domain.ErrInternal)
			},
			expectedErr: "server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockDomain.NewMockAuthRepository(ctrl)
			test.mockBehaviorDeleteSession(repo, test.input)

			s := New(repo)

			_, err := s.DeleteBySessionID(context.Background(), &protobuf.SessionID{
				SessionId: test.input,
			})

			if err != nil {
				require.Equal(t, test.expectedErr, err.Error())
			}
		})
	}
}

func TestAuth_GetBySessionID(t *testing.T) {
	type mockBehaviorGetBySessionID func(r *mockDomain.MockAuthRepository, id string)

	tests := []struct {
		name                       string
		input                      string
		mockBehaviorGetBySessionID mockBehaviorGetBySessionID
		expected                   *protobuf.Session
		expectedErr                string
	}{
		{
			name:  "OK",
			input: "sessionID",
			mockBehaviorGetBySessionID: func(r *mockDomain.MockAuthRepository, id string) {
				r.EXPECT().GetBySessionID(id).Return(models.Cookie{
					Value: "sessionID",
				}, nil)
			},
			expected: &protobuf.Session{
				SessionId: "sessionID",
				Expires:   timestamppb.New(time.Time{}),
			},
		},
		{
			name:  "Error",
			input: "sessionID",
			mockBehaviorGetBySessionID: func(r *mockDomain.MockAuthRepository, id string) {
				r.EXPECT().GetBySessionID(id).Return(models.Cookie{}, domain.ErrInternal)
			},
			expectedErr: "server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockDomain.NewMockAuthRepository(ctrl)
			test.mockBehaviorGetBySessionID(repo, test.input)

			s := New(repo)

			session, err := s.GetBySessionID(context.Background(), &protobuf.SessionID{
				SessionId: test.input,
			})

			if err != nil {
				require.Equal(t, test.expectedErr, err.Error())
			} else {
				require.Equal(t, test.expected, session)
			}
		})
	}
}
