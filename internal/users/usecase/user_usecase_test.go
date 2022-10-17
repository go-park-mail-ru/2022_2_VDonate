package users

import (
	"errors"
	"testing"

	mock_users "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/users/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_Update(t *testing.T) {
	type mockBehaviourGet func(r *mock_users.MockRepository, userId uint64)
	type mockBehaviourUpdate func(r *mock_users.MockRepository, user *models.User)

	tests := []struct {
		name                string
		inputUser           *models.User
		mockBehaviourGet    mockBehaviourGet
		mockBehaviourUpdate mockBehaviourUpdate
		response            *models.User
		responseError       string
	}{
		{
			name: "OK",
			inputUser: &models.User{
				ID:       200,
				Username: "user",
			},
			mockBehaviourGet: func(r *mock_users.MockRepository, userId uint64) {
				r.EXPECT().GetByID(userId).Return(&models.User{
					ID:       200,
					Username: "user",
				}, nil)
			},
			mockBehaviourUpdate: func(r *mock_users.MockRepository, user *models.User) {
				r.EXPECT().Update(user).Return(&models.User{
					ID:       200,
					Username: "username",
					Email:    "user@ex.org",
				}, nil)
			},
			response: &models.User{
				ID:       200,
				Username: "username",
				Email:    "user@ex.org",
			},
		},
		{
			name: "NotFound",
			inputUser: &models.User{
				ID:       200,
				Username: "user",
			},
			mockBehaviourGet: func(r *mock_users.MockRepository, userId uint64) {
				r.EXPECT().GetByID(userId).Return(&models.User{}, errors.New("not found"))
			},
			mockBehaviourUpdate: func(r *mock_users.MockRepository, user *models.User) {},
			response:            nil,
			responseError:       "not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mock_users.NewMockRepository(ctrl)
			test.mockBehaviourGet(userMock, test.inputUser.ID)
			test.mockBehaviourUpdate(userMock, test.inputUser)

			usecase := New(userMock)

			user, err := usecase.Update(test.inputUser.ID, test.inputUser)
			if err != nil {
				require.EqualError(t, err, test.responseError)
			}
			require.Equal(t, user, test.response)
		})
	}
}

func TestUsecase_DeleteByUsername(t *testing.T) {
	type mockBehaviourGet func(r *mock_users.MockRepository, username string)
	type mockBehaviourDelete func(r *mock_users.MockRepository, userId uint64)

	tests := []struct {
		name                string
		userId              uint64
		usernmae            string
		mockBehaviourGet    mockBehaviourGet
		mockBehaviourDelete mockBehaviourDelete
		responseError       error
	}{
		{
			name:     "OK",
			userId:   123,
			usernmae: "user",
			mockBehaviourGet: func(r *mock_users.MockRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(&models.User{
					ID:       123,
					Username: username,
				}, nil)
			},
			mockBehaviourDelete: func(r *mock_users.MockRepository, userId uint64) {
				r.EXPECT().DeleteByID(userId).Return(nil)
			},
			responseError: nil,
		},
		{
			name:     "UserNotFound",
			userId:   123,
			usernmae: "user",
			mockBehaviourGet: func(r *mock_users.MockRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(&models.User{}, errors.New("user not found"))
			},
			mockBehaviourDelete: func(r *mock_users.MockRepository, userId uint64) {},
			responseError:       errors.New("user not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mock_users.NewMockRepository(ctrl)
			test.mockBehaviourGet(userMock, test.usernmae)
			test.mockBehaviourDelete(userMock, test.userId)

			usecase := New(userMock)

			err := usecase.DeleteByUsername(test.usernmae)

			require.Equal(t, err, test.responseError)
		})
	}
}

func TestUsecase_DeleteByEmail(t *testing.T) {
	type mockBehaviourGet func(r *mock_users.MockRepository, email string)
	type mockBehaviourDelete func(r *mock_users.MockRepository, userId uint64)

	tests := []struct {
		name                string
		userId              uint64
		email               string
		mockBehaviourGet    mockBehaviourGet
		mockBehaviourDelete mockBehaviourDelete
		responseError       error
	}{
		{
			name:   "OK",
			userId: 200,
			email:  "ex@mail.ru",
			mockBehaviourGet: func(r *mock_users.MockRepository, email string) {
				r.EXPECT().GetByEmail(email).Return(&models.User{
					ID:       200,
					Username: "user",
					Email:    email,
				}, nil)
			},
			mockBehaviourDelete: func(r *mock_users.MockRepository, userId uint64) {
				r.EXPECT().DeleteByID(userId).Return(nil)
			},
			responseError: nil,
		},
		{
			name:   "UserNotExist",
			userId: 200,
			email:  "ex@mail.ru",
			mockBehaviourGet: func(r *mock_users.MockRepository, email string) {
				r.EXPECT().GetByEmail(email).Return(&models.User{}, errors.New("user not found"))
			},
			mockBehaviourDelete: func(r *mock_users.MockRepository, userId uint64) {},
			responseError:       errors.New("user not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mock_users.NewMockRepository(ctrl)
			test.mockBehaviourGet(userMock, test.email)
			test.mockBehaviourDelete(userMock, test.userId)

			usecase := New(userMock)

			err := usecase.DeleteByEmail(test.email)

			require.Equal(t, err, test.responseError)
		})
	}
}

func TestUsecase_CheckIDAndPassword(t *testing.T) {
	type mockBehaviourGet func(r *mock_users.MockRepository, userId uint64)

	tests := []struct {
		name             string
		userId           uint64
		password         string
		mockBehaviourGet mockBehaviourGet
		response         bool
	}{
		{
			name:     "OK",
			userId:   200,
			password: "Qwerty",
			mockBehaviourGet: func(r *mock_users.MockRepository, userId uint64) {
				r.EXPECT().GetByID(userId).Return(&models.User{
					ID:       userId,
					Password: "Qwerty",
				}, nil)
			},
			response: true,
		},
		{
			name:     "UserNotExist",
			userId:   200,
			password: "Qwerty",
			mockBehaviourGet: func(r *mock_users.MockRepository, userId uint64) {
				r.EXPECT().GetByID(userId).Return(&models.User{}, errors.New("user not found"))
			},
			response: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mock_users.NewMockRepository(ctrl)
			test.mockBehaviourGet(userMock, test.userId)

			usecase := New(userMock)

			isRight := usecase.CheckIDAndPassword(test.userId, test.password)

			require.Equal(t, isRight, test.response)
		})
	}
}

func TestUsecase_IsExistUsernameAndEmail(t *testing.T) {
	type mockBehaviourGetByUsername func(r *mock_users.MockRepository, username string)
	type mockBehaviourGetByEmail func(r *mock_users.MockRepository, email string)

	tests := []struct {
		name                       string
		username                   string
		email                      string
		mockBehaviourGetByUsername mockBehaviourGetByUsername
		mockBehaviourGetByEmail    mockBehaviourGetByEmail
		response                   bool
	}{
		{
			name:     "OK",
			username: "user",
			email:    "a@d.com",
			mockBehaviourGetByUsername: func(r *mock_users.MockRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(&models.User{
					Username: username,
					Email:    "a@d.com",
				}, nil)
			},
			mockBehaviourGetByEmail: func(r *mock_users.MockRepository, email string) {
				r.EXPECT().GetByEmail(email).Return(&models.User{
					Username: "user",
					Email:    email,
				}, nil)
			},
			response: true,
		},
		{
			name:     "IncorrectUsername",
			username: "user",
			email:    "a@d.com",
			mockBehaviourGetByUsername: func(r *mock_users.MockRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(&models.User{}, errors.New("not exist"))
			},
			mockBehaviourGetByEmail: func(r *mock_users.MockRepository, email string) {},
			response: false,
		},
		{
			name:     "IncorrectUsername",
			username: "user",
			email:    "a@d.com",
			mockBehaviourGetByUsername: func(r *mock_users.MockRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(&models.User{
					Username: username,
					Email:    "adnsonjo@d.com",
				}, nil)
			},
			mockBehaviourGetByEmail: func(r *mock_users.MockRepository, email string) {
				r.EXPECT().GetByEmail(email).Return(&models.User{}, errors.New("not exist"))
			},
			response: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mock_users.NewMockRepository(ctrl)
			test.mockBehaviourGetByUsername(userMock, test.username)
			test.mockBehaviourGetByEmail(userMock, test.email)

			usecase := New(userMock)

			isExist := usecase.IsExistUsernameAndEmail(test.username, test.email)

			require.Equal(t, isExist, test.response)
		})
	}
}
