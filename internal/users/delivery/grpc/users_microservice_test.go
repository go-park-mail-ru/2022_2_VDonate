package usersMicroservice

import (
	"context"
	"testing"

	authProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestUsersMicroservice_Update(t *testing.T) {
	type mockUpdate func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.User, opts ...grpc.CallOption)

	tests := []struct {
		name  string
		users models.User
		mock  mockUpdate
		err   error
	}{
		{
			name: "OK",
			users: models.User{
				ID: 1,
			},
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.User, opts ...grpc.CallOption) {
				s.EXPECT().Update(c, in).Return(&emptypb.Empty{}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.User, opts ...grpc.CallOption) {
				s.EXPECT().Update(c, in).Return(&emptypb.Empty{}, grpc.ErrClientConnClosing)
			},
			err: grpc.ErrClientConnClosing,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockUsersClient(ctrl)

			test.mock(mock, context.Background(), &userProto.User{
				Id: test.users.ID,
			})

			client := New(mock)
			err := client.Update(test.users)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestUsersMicroservice_Create(t *testing.T) {
	type mockCreate func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.User, opts ...grpc.CallOption)

	tests := []struct {
		name  string
		users models.User
		mock  mockCreate
		err   error
	}{
		{
			name: "OK",
			users: models.User{
				ID: 1,
			},
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.User, opts ...grpc.CallOption) {
				s.EXPECT().Create(c, in).Return(&userProto.UserID{
					UserId: 1,
				}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.User, opts ...grpc.CallOption) {
				s.EXPECT().Create(c, in).Return(&userProto.UserID{}, grpc.ErrClientConnClosing)
			},
			err: grpc.ErrClientConnClosing,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockUsersClient(ctrl)

			test.mock(mock, context.Background(), &userProto.User{
				Id: test.users.ID,
			})

			client := New(mock)
			id, err := client.Create(test.users)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.users.ID, id)
		})
	}
}

func TestUsersMicroservice_GetAuthorByUsername(t *testing.T) {
	type mockGetAuthorByUsername func(s *mockDomain.MockUsersClient, c context.Context, in *protobuf.Keyword, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		keyword string
		mock    mockGetAuthorByUsername
		err     error
	}{
		{
			name:    "OK",
			keyword: "test",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.Keyword, opts ...grpc.CallOption) {
				s.EXPECT().GetAuthorByUsername(c, in).Return(&userProto.UsersArray{
					Users: []*userProto.User{
						{
							Id: 1,
						},
					},
				}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.Keyword, opts ...grpc.CallOption) {
				s.EXPECT().GetAuthorByUsername(c, in).Return(&userProto.UsersArray{}, grpc.ErrClientConnClosing)
			},
			err: grpc.ErrClientConnClosing,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockUsersClient(ctrl)

			test.mock(mock, context.Background(), &userProto.Keyword{
				Keyword: test.keyword,
			})

			client := New(mock)
			_, err := client.GetAuthorByUsername(test.keyword)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestUsersMicroservice_GetAllAuthors(t *testing.T) {
	type mockGetAllAuthors func(s *mockDomain.MockUsersClient, c context.Context, in *emptypb.Empty, opts ...grpc.CallOption)

	tests := []struct {
		name string
		mock mockGetAllAuthors
		err  error
	}{
		{
			name: "OK",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *emptypb.Empty, opts ...grpc.CallOption) {
				s.EXPECT().GetAllAuthors(c, in).Return(&userProto.UsersArray{
					Users: []*userProto.User{
						{
							Id: 1,
						},
					},
				}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *emptypb.Empty, opts ...grpc.CallOption) {
				s.EXPECT().GetAllAuthors(c, in).Return(&userProto.UsersArray{}, grpc.ErrClientConnClosing)
			},
			err: grpc.ErrClientConnClosing,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockUsersClient(ctrl)

			test.mock(mock, context.Background(), &emptypb.Empty{})

			client := New(mock)
			_, err := client.GetAllAuthors()
			assert.Equal(t, test.err, err)
		})
	}
}

func TestUsersMicroservice_GetByID(t *testing.T) {
	type mockGetByID func(s *mockDomain.MockUsersClient, c context.Context, in *protobuf.UserID, opts ...grpc.CallOption)

	tests := []struct {
		name string
		id   uint64
		mock mockGetByID
		err  error
	}{
		{
			name: "OK",
			id:   1,
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().GetByID(c, in).Return(&userProto.User{
					Id: 1,
				}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().GetByID(c, in).Return(&userProto.User{}, grpc.ErrClientConnClosing)
			},
			err: grpc.ErrClientConnClosing,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockUsersClient(ctrl)

			test.mock(mock, context.Background(), &userProto.UserID{
				UserId: test.id,
			})

			client := New(mock)
			_, err := client.GetByID(test.id)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestUsersMicroservice_GetBySessionID(t *testing.T) {
	type mockGetBySessionID func(s *mockDomain.MockUsersClient, c context.Context, in *authProto.SessionID, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		session string
		mock    mockGetBySessionID
		err     error
	}{
		{
			name:    "OK",
			session: "session",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *authProto.SessionID, opts ...grpc.CallOption) {
				s.EXPECT().GetBySessionID(c, in).Return(&userProto.User{
					Id: 1,
				}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *authProto.SessionID, opts ...grpc.CallOption) {
				s.EXPECT().GetBySessionID(c, in).Return(&userProto.User{}, grpc.ErrClientConnClosing)
			},
			err: grpc.ErrClientConnClosing,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockUsersClient(ctrl)

			test.mock(mock, context.Background(), &authProto.SessionID{
				SessionId: test.session,
			})

			client := New(mock)
			_, err := client.GetBySessionID(test.session)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestUsersMicroservice_GetByEmail(t *testing.T) {
	type mockGetByEmail func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.Email, opts ...grpc.CallOption)

	tests := []struct {
		name  string
		email string
		mock  mockGetByEmail
		err   error
	}{
		{
			name:  "OK",
			email: "email",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.Email, opts ...grpc.CallOption) {
				s.EXPECT().GetByEmail(c, in).Return(&userProto.User{
					Id: 1,
				}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.Email, opts ...grpc.CallOption) {
				s.EXPECT().GetByEmail(c, in).Return(&userProto.User{}, grpc.ErrClientConnClosing)
			},
			err: grpc.ErrClientConnClosing,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockUsersClient(ctrl)

			test.mock(mock, context.Background(), &userProto.Email{
				Email: test.email,
			})

			client := New(mock)
			_, err := client.GetByEmail(test.email)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestUsersMicroservice_GetByUsername(t *testing.T) {
	type mockGetByUsername func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.Username, opts ...grpc.CallOption)

	tests := []struct {
		name     string
		username string
		mock     mockGetByUsername
		err      error
	}{
		{
			name:     "OK",
			username: "username",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.Username, opts ...grpc.CallOption) {
				s.EXPECT().GetByUsername(c, in).Return(&userProto.User{
					Id: 1,
				}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.Username, opts ...grpc.CallOption) {
				s.EXPECT().GetByUsername(c, in).Return(&userProto.User{}, grpc.ErrClientConnClosing)
			},
			err: grpc.ErrClientConnClosing,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockUsersClient(ctrl)

			test.mock(mock, context.Background(), &userProto.Username{
				Username: test.username,
			})

			client := New(mock)
			_, err := client.GetByUsername(test.username)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestUsersMicroservice_GetUserByPostID(t *testing.T) {
	type mockGetUserByPostID func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.PostID, opts ...grpc.CallOption)

	tests := []struct {
		name   string
		postID uint64
		mock   mockGetUserByPostID
		err    error
	}{
		{
			name:   "OK",
			postID: 1,
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.PostID, opts ...grpc.CallOption) {
				s.EXPECT().GetUserByPostID(c, in).Return(&userProto.User{
					Id: 1,
				}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.PostID, opts ...grpc.CallOption) {
				s.EXPECT().GetUserByPostID(c, in).Return(&userProto.User{}, grpc.ErrClientConnClosing)
			},
			err: grpc.ErrClientConnClosing,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockUsersClient(ctrl)

			test.mock(mock, context.Background(), &userProto.PostID{
				PostID: test.postID,
			})

			client := New(mock)
			_, err := client.GetUserByPostID(test.postID)
			assert.Equal(t, test.err, err)
		})
	}
}
