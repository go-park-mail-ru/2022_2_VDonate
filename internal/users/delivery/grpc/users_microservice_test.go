package usersMicroservice

import (
	"context"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
				s.EXPECT().Update(c, in).Return(&emptypb.Empty{}, status.Error(codes.Canceled, "canceled"))
			},
			err: status.Error(codes.Canceled, "canceled"),
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
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
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
				s.EXPECT().Create(c, in).Return(&userProto.UserID{}, status.Error(codes.Canceled, "canceled"))
			},
			err: status.Error(codes.Canceled, "canceled"),
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
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
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
				s.EXPECT().GetAuthorByUsername(c, in).Return(&userProto.UsersArray{}, status.Error(codes.Canceled, "canceled"))
			},
			err: status.Error(codes.Canceled, "canceled"),
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
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
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
				s.EXPECT().GetAllAuthors(c, in).Return(&userProto.UsersArray{}, status.Error(codes.Canceled, "canceled"))
			},
			err: status.Error(codes.Canceled, "canceled"),
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
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
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
				s.EXPECT().GetByID(c, in).Return(&userProto.User{}, status.Error(codes.Canceled, "canceled"))
			},
			err: status.Error(codes.Canceled, "canceled"),
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
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
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
				s.EXPECT().GetBySessionID(c, in).Return(&userProto.User{}, status.Error(codes.Canceled, "canceled"))
			},
			err: status.Error(codes.Canceled, "canceled"),
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
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
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
				s.EXPECT().GetByEmail(c, in).Return(&userProto.User{}, status.Error(codes.Canceled, "canceled"))
			},
			err: status.Error(codes.Canceled, "canceled"),
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
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
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
				s.EXPECT().GetByUsername(c, in).Return(&userProto.User{}, domain.ErrInternal)
			},
			err: domain.ErrInternal,
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
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
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
				s.EXPECT().GetUserByPostID(c, in).Return(&userProto.User{}, domain.ErrInternal)
			},
			err: domain.ErrInternal,
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
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
		})
	}
}

func TestUsersMicroservice_GetPostsNum(t *testing.T) {
	type mockGetPostsNum func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption)

	tests := []struct {
		name string
		id   uint64
		mock mockGetPostsNum
		err  error
	}{
		{
			name: "OK",
			id:   1,
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().GetPostsNum(c, in).Return(&userProto.PostsNum{
					CountPosts: 1,
				}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().GetPostsNum(c, in).Return(&userProto.PostsNum{}, domain.ErrInternal)
			},
			err: domain.ErrInternal,
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
			_, err := client.GetPostsNum(test.id)
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
		})
	}
}

func TestUsersMicroservice_GetSubscribersNum(t *testing.T) {
	type mockGetSubscribersNum func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption)

	tests := []struct {
		name string
		id   uint64
		mock mockGetSubscribersNum
		err  error
	}{
		{
			name: "OK",
			id:   1,
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().GetSubscribersNumForMounth(c, in).Return(&userProto.SubscribersNum{
					CountSubscribers: 1,
				}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().GetSubscribersNumForMounth(c, in).Return(&userProto.SubscribersNum{}, domain.ErrInternal)
			},
			err: domain.ErrInternal,
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
			_, err := client.GetSubscribersNum(test.id)
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
		})
	}
}

func TestUsersMicroservice_GetProfitForMounth(t *testing.T) {
	type mockGetProfitForMounth func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption)

	tests := []struct {
		name string
		id   uint64
		mock mockGetProfitForMounth
		err  error
	}{
		{
			name: "OK",
			id:   1,
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().GetProfitForMounth(c, in).Return(&userProto.Profit{
					CountMounthProfit: 1,
				}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().GetProfitForMounth(c, in).Return(&userProto.Profit{}, domain.ErrInternal)
			},
			err: domain.ErrInternal,
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
			_, err := client.GetProfitForMounth(test.id)
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
		})
	}
}

func TestUsersMicroservice_DropBalance(t *testing.T) {
	type mockDropBalance func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption)

	tests := []struct {
		name string
		id   uint64
		mock mockDropBalance
		err  error
	}{
		{
			name: "OK",
			id:   1,
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().DropBalance(c, in).Return(&emptypb.Empty{}, nil)
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockUsersClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().DropBalance(c, in).Return(&emptypb.Empty{}, domain.ErrInternal)
			},
			err: domain.ErrInternal,
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
			err := client.DropBalance(test.id)
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}
		})
	}
}
