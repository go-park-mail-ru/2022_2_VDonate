package grpcUsers

import (
	"context"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	authProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	mock_domain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestConvertToProto(t *testing.T) {
	input := models.User{
		ID:       1,
		Username: "username",
		Email:    "email",
		Password: "password",
		Avatar:   "avatar",
		IsAuthor: true,
		About:    "about",
	}

	expected := &protobuf.User{
		Id:       1,
		Username: "username",
		Email:    "email",
		Password: "password",
		Avatar:   "avatar",
		IsAuthor: true,
		About:    "about",
	}

	actual := ConvertToProto(input)

	assert.Equal(t, expected, actual)
}

func TestConvertToModel(t *testing.T) {
	input := &protobuf.User{
		Id:       1,
		Username: "username",
		Email:    "email",
		Password: "password",
		Avatar:   "avatar",
		IsAuthor: true,
		About:    "about",
	}

	expected := models.User{
		ID:       1,
		Username: "username",
		Email:    "email",
		Password: "password",
		Avatar:   "avatar",
		IsAuthor: true,
		About:    "about",
	}

	actual := ConvertToModel(input)

	assert.Equal(t, expected, actual)
}

func TestUserService_GetAllAuthors(t *testing.T) {
	type mockBehaviorGetAllAuthors func(r *mock_domain.MockUsersRepository)

	tests := []struct {
		name                      string
		input                     *emptypb.Empty
		mockBehaviorGetAllAuthors mockBehaviorGetAllAuthors
		expected                  *protobuf.UsersArray
		expectedError             string
	}{
		{
			name:  "OK",
			input: &emptypb.Empty{},
			mockBehaviorGetAllAuthors: func(r *mock_domain.MockUsersRepository) {
				r.EXPECT().GetAllAuthors().Return([]models.User{
					{
						ID:       1,
						Username: "username",
						Email:    "email",
						Password: "password",
						Avatar:   "avatar",
						IsAuthor: true,
						About:    "about",
					},
				}, nil)
			},
			expected: &protobuf.UsersArray{
				Users: []*protobuf.User{
					{
						Id:       1,
						Username: "username",
						Email:    "email",
						Password: "password",
						Avatar:   "avatar",
						IsAuthor: true,
						About:    "about",
					},
				},
			},
		},
		{
			name:  "Error",
			input: &emptypb.Empty{},
			mockBehaviorGetAllAuthors: func(r *mock_domain.MockUsersRepository) {
				r.EXPECT().GetAllAuthors().Return(nil, domain.ErrInternal)
			},
			expectedError: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockUsersRepository(ctrl)
			test.mockBehaviorGetAllAuthors(repo)

			s := New(repo)

			users, err := s.GetAllAuthors(context.Background(), test.input)

			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			} else {
				assert.Equal(t, test.expected, users)
			}
		})
	}
}

func TestUserService_Create(t *testing.T) {
	type mockBehaviorCreate func(r *mock_domain.MockUsersRepository, user models.User)

	tests := []struct {
		name               string
		input              *protobuf.User
		mockBehaviorCreate mockBehaviorCreate
		expected           *protobuf.UserID
		expectedError      string
	}{
		{
			name: "OK",
			input: &protobuf.User{
				Id: 1,
			},
			mockBehaviorCreate: func(r *mock_domain.MockUsersRepository, user models.User) {
				r.EXPECT().Create(user).Return(uint64(1), nil)
			},
			expected: &protobuf.UserID{
				UserId: 1,
			},
		},
		{
			name: "Error",
			input: &protobuf.User{
				Id: 1,
			},
			mockBehaviorCreate: func(r *mock_domain.MockUsersRepository, user models.User) {
				r.EXPECT().Create(user).Return(uint64(0), domain.ErrInternal)
			},
			expectedError: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockUsersRepository(ctrl)
			test.mockBehaviorCreate(repo, ConvertToModel(test.input))

			s := New(repo)

			id, err := s.Create(context.Background(), test.input)

			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			} else {
				assert.Equal(t, test.expected, id)
			}
		})
	}
}

func TestUserService_Update(t *testing.T) {
	type mockBehaviorUpdate func(r *mock_domain.MockUsersRepository, user models.User)

	tests := []struct {
		name               string
		input              *protobuf.User
		mockBehaviorUpdate mockBehaviorUpdate
		expectedError      string
	}{
		{
			name: "OK",
			input: &protobuf.User{
				Id: 1,
			},
			mockBehaviorUpdate: func(r *mock_domain.MockUsersRepository, user models.User) {
				r.EXPECT().Update(user).Return(nil)
			},
		},
		{
			name: "Error",
			input: &protobuf.User{
				Id: 1,
			},
			mockBehaviorUpdate: func(r *mock_domain.MockUsersRepository, user models.User) {
				r.EXPECT().Update(user).Return(domain.ErrInternal)
			},
			expectedError: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockUsersRepository(ctrl)
			test.mockBehaviorUpdate(repo, ConvertToModel(test.input))

			s := New(repo)

			_, err := s.Update(context.Background(), test.input)
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestUserService_GetAuthorByUsername(t *testing.T) {
	type mockBehaviorGetAuthorByUsername func(r *mock_domain.MockUsersRepository, keyword string)

	tests := []struct {
		name                            string
		keyword                         *protobuf.Keyword
		mockBehaviorGetAuthorByUsername mockBehaviorGetAuthorByUsername
		expected                        *protobuf.UsersArray
		expectedError                   string
	}{
		{
			name: "OK",
			keyword: &protobuf.Keyword{
				Keyword: "username",
			},
			mockBehaviorGetAuthorByUsername: func(r *mock_domain.MockUsersRepository, keyword string) {
				r.EXPECT().GetAuthorByUsername(keyword).Return([]models.User{
					{
						ID:       1,
						Username: "username",
					},
				}, nil)
			},
			expected: &protobuf.UsersArray{
				Users: []*protobuf.User{
					{
						Id:       1,
						Username: "username",
					},
				},
			},
		},
		{
			name: "Error",
			keyword: &protobuf.Keyword{
				Keyword: "username",
			},
			mockBehaviorGetAuthorByUsername: func(r *mock_domain.MockUsersRepository, keyword string) {
				r.EXPECT().GetAuthorByUsername(keyword).Return(nil, domain.ErrInternal)
			},
			expectedError: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockUsersRepository(ctrl)
			test.mockBehaviorGetAuthorByUsername(repo, test.keyword.Keyword)

			s := New(repo)

			users, err := s.GetAuthorByUsername(context.Background(), test.keyword)

			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			} else {
				assert.Equal(t, test.expected, users)
			}
		})
	}
}

func TestUserService_GetUserByPostID(t *testing.T) {
	type mockBehaviorGetUserByPostID func(r *mock_domain.MockUsersRepository, postID uint64)

	tests := []struct {
		name                        string
		postID                      *protobuf.PostID
		mockBehaviorGetUserByPostID mockBehaviorGetUserByPostID
		expected                    *protobuf.User
		expectedError               string
	}{
		{
			name: "OK",
			postID: &protobuf.PostID{
				PostID: 1,
			},
			mockBehaviorGetUserByPostID: func(r *mock_domain.MockUsersRepository, postID uint64) {
				r.EXPECT().GetUserByPostID(postID).Return(models.User{
					ID: 1,
				}, nil)
			},
			expected: &protobuf.User{
				Id: 1,
			},
		},
		{
			name: "Error",
			postID: &protobuf.PostID{
				PostID: 1,
			},
			mockBehaviorGetUserByPostID: func(r *mock_domain.MockUsersRepository, postID uint64) {
				r.EXPECT().GetUserByPostID(postID).Return(models.User{}, domain.ErrInternal)
			},
			expectedError: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockUsersRepository(ctrl)
			test.mockBehaviorGetUserByPostID(repo, test.postID.PostID)

			s := New(repo)

			user, err := s.GetUserByPostID(context.Background(), test.postID)

			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			} else {
				assert.Equal(t, test.expected, user)
			}
		})
	}
}

func TestUserService_GetByID(t *testing.T) {
	type mockBehaviorGetByID func(r *mock_domain.MockUsersRepository, id uint64)

	tests := []struct {
		name                string
		id                  *protobuf.UserID
		mockBehaviorGetByID mockBehaviorGetByID
		expected            *protobuf.User
		expectedError       string
	}{
		{
			name: "OK",
			id: &protobuf.UserID{
				UserId: 1,
			},
			mockBehaviorGetByID: func(r *mock_domain.MockUsersRepository, id uint64) {
				r.EXPECT().GetByID(id).Return(models.User{
					ID: 1,
				}, nil)
			},
			expected: &protobuf.User{
				Id: 1,
			},
		},
		{
			name: "Error",
			id: &protobuf.UserID{
				UserId: 1,
			},
			mockBehaviorGetByID: func(r *mock_domain.MockUsersRepository, id uint64) {
				r.EXPECT().GetByID(id).Return(models.User{}, domain.ErrInternal)
			},
			expectedError: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockUsersRepository(ctrl)
			test.mockBehaviorGetByID(repo, test.id.UserId)

			s := New(repo)

			user, err := s.GetByID(context.Background(), test.id)

			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			} else {
				assert.Equal(t, test.expected, user)
			}
		})
	}
}

func TestUserService_GetBySessionID(t *testing.T) {
	type mockBehaviorGetBySessionID func(r *mock_domain.MockUsersRepository, sessionID string)

	tests := []struct {
		name                       string
		sessionID                  *authProto.SessionID
		mockBehaviorGetBySessionID mockBehaviorGetBySessionID
		expected                   *protobuf.User
		expectedError              string
	}{
		{
			name: "OK",
			sessionID: &authProto.SessionID{
				SessionId: "sessionID",
			},
			mockBehaviorGetBySessionID: func(r *mock_domain.MockUsersRepository, sessionID string) {
				r.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 1,
				}, nil)
			},
			expected: &protobuf.User{
				Id: 1,
			},
		},
		{
			name: "Error",
			sessionID: &authProto.SessionID{
				SessionId: "sessionID",
			},
			mockBehaviorGetBySessionID: func(r *mock_domain.MockUsersRepository, sessionID string) {
				r.EXPECT().GetBySessionID(sessionID).Return(models.User{}, domain.ErrInternal)
			},
			expectedError: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockUsersRepository(ctrl)
			test.mockBehaviorGetBySessionID(repo, test.sessionID.SessionId)

			s := New(repo)

			user, err := s.GetBySessionID(context.Background(), test.sessionID)

			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			} else {
				assert.Equal(t, test.expected, user)
			}
		})
	}
}

func TestUserService_GetByEmail(t *testing.T) {
	type mockBehaviorGetByEmail func(r *mock_domain.MockUsersRepository, email string)

	tests := []struct {
		name                   string
		email                  *protobuf.Email
		mockBehaviorGetByEmail mockBehaviorGetByEmail
		expected               *protobuf.User
		expectedError          string
	}{
		{
			name: "OK",
			email: &protobuf.Email{
				Email: "email",
			},
			mockBehaviorGetByEmail: func(r *mock_domain.MockUsersRepository, email string) {
				r.EXPECT().GetByEmail(email).Return(models.User{
					ID: 1,
				}, nil)
			},
			expected: &protobuf.User{
				Id: 1,
			},
		},
		{
			name: "Error",
			email: &protobuf.Email{
				Email: "email",
			},
			mockBehaviorGetByEmail: func(r *mock_domain.MockUsersRepository, email string) {
				r.EXPECT().GetByEmail(email).Return(models.User{}, domain.ErrInternal)
			},
			expectedError: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockUsersRepository(ctrl)
			test.mockBehaviorGetByEmail(repo, test.email.Email)

			s := New(repo)

			user, err := s.GetByEmail(context.Background(), test.email)

			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			} else {
				assert.Equal(t, test.expected, user)
			}
		})
	}
}

func TestUserService_GetByUsername(t *testing.T) {
	type mockBehaviorGetByUsername func(r *mock_domain.MockUsersRepository, username string)

	tests := []struct {
		name                      string
		username                  *protobuf.Username
		mockBehaviorGetByUsername mockBehaviorGetByUsername
		expected                  *protobuf.User
		expectedError             string
	}{
		{
			name: "OK",
			username: &protobuf.Username{
				Username: "username",
			},
			mockBehaviorGetByUsername: func(r *mock_domain.MockUsersRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(models.User{
					ID: 1,
				}, nil)
			},
			expected: &protobuf.User{
				Id: 1,
			},
		},
		{
			name: "Error",
			username: &protobuf.Username{
				Username: "username",
			},
			mockBehaviorGetByUsername: func(r *mock_domain.MockUsersRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, domain.ErrInternal)
			},
			expectedError: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockUsersRepository(ctrl)
			test.mockBehaviorGetByUsername(repo, test.username.Username)

			s := New(repo)

			user, err := s.GetByUsername(context.Background(), test.username)

			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			} else {
				assert.Equal(t, test.expected, user)
			}
		})
	}
}
