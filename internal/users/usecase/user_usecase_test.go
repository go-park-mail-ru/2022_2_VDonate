package users

import (
	"errors"
	"testing"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_Update(t *testing.T) {
	type mockBehaviourGet func(r *mockDomain.MockUsersRepository, userID uint64)

	type mockBehaviourUpdate func(r *mockDomain.MockUsersRepository, user *models.User)

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
			mockBehaviourGet: func(r *mockDomain.MockUsersRepository, userID uint64) {
				r.EXPECT().GetByID(userID).Return(&models.User{
					ID:       200,
					Username: "user",
				}, nil)
			},
			mockBehaviourUpdate: func(r *mockDomain.MockUsersRepository, user *models.User) {
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
			mockBehaviourGet: func(r *mockDomain.MockUsersRepository, userID uint64) {
				r.EXPECT().GetByID(userID).Return(&models.User{}, errors.New("not found"))
			},
			mockBehaviourUpdate: func(r *mockDomain.MockUsersRepository, user *models.User) {},
			response:            nil,
			responseError:       "not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersRepository(ctrl)

			test.mockBehaviourGet(userMock, test.inputUser.ID)
			test.mockBehaviourUpdate(userMock, test.inputUser)

			usecase := New(userMock)

			user, err := usecase.Update(*test.inputUser)
			if err != nil {
				require.EqualError(t, err, test.responseError)
			}
			require.Equal(t, user, test.response)
		})
	}
}

func TestUsecase_DeleteByUsername(t *testing.T) {
	type mockBehaviourGet func(r *mockDomain.MockUsersRepository, username string)

	type mockBehaviourDelete func(r *mockDomain.MockUsersRepository, userID uint64)

	tests := []struct {
		name                string
		userID              uint64
		usernmae            string
		mockBehaviourGet    mockBehaviourGet
		mockBehaviourDelete mockBehaviourDelete
		responseError       error
	}{
		{
			name:     "OK",
			userID:   123,
			usernmae: "user",
			mockBehaviourGet: func(r *mockDomain.MockUsersRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(&models.User{
					ID:       123,
					Username: username,
				}, nil)
			},
			mockBehaviourDelete: func(r *mockDomain.MockUsersRepository, userID uint64) {
				r.EXPECT().DeleteByID(userID).Return(nil)
			},
			responseError: nil,
		},
		{
			name:     "UserNotFound",
			userID:   123,
			usernmae: "user",
			mockBehaviourGet: func(r *mockDomain.MockUsersRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(&models.User{}, errors.New("user not found"))
			},
			mockBehaviourDelete: func(r *mockDomain.MockUsersRepository, userID uint64) {},
			responseError:       errors.New("user not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersRepository(ctrl)
			test.mockBehaviourGet(userMock, test.usernmae)
			test.mockBehaviourDelete(userMock, test.userID)

			usecase := New(userMock)

			err := usecase.DeleteByUsername(test.usernmae)

			require.Equal(t, err, test.responseError)
		})
	}
}

func TestUsecase_DeleteByEmail(t *testing.T) {
	type mockBehaviourGet func(r *mockDomain.MockUsersRepository, email string)

	type mockBehaviourDelete func(r *mockDomain.MockUsersRepository, userID uint64)

	tests := []struct {
		name                string
		userID              uint64
		email               string
		mockBehaviourGet    mockBehaviourGet
		mockBehaviourDelete mockBehaviourDelete
		responseError       error
	}{
		{
			name:   "OK",
			userID: 200,
			email:  "ex@mail.ru",
			mockBehaviourGet: func(r *mockDomain.MockUsersRepository, email string) {
				r.EXPECT().GetByEmail(email).Return(&models.User{
					ID:       200,
					Username: "user",
					Email:    email,
				}, nil)
			},
			mockBehaviourDelete: func(r *mockDomain.MockUsersRepository, userID uint64) {
				r.EXPECT().DeleteByID(userID).Return(nil)
			},
			responseError: nil,
		},
		{
			name:   "UserNotExist",
			userID: 200,
			email:  "ex@mail.ru",
			mockBehaviourGet: func(r *mockDomain.MockUsersRepository, email string) {
				r.EXPECT().GetByEmail(email).Return(&models.User{}, errors.New("user not found"))
			},
			mockBehaviourDelete: func(r *mockDomain.MockUsersRepository, userID uint64) {},
			responseError:       errors.New("user not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersRepository(ctrl)
			test.mockBehaviourGet(userMock, test.email)
			test.mockBehaviourDelete(userMock, test.userID)

			usecase := New(userMock)

			err := usecase.DeleteByEmail(test.email)

			require.Equal(t, err, test.responseError)
		})
	}
}

func TestUsecase_CheckIDAndPassword(t *testing.T) {
	type mockBehaviourGet func(r *mockDomain.MockUsersRepository, userID uint64)

	tests := []struct {
		name             string
		userID           uint64
		password         string
		mockBehaviourGet mockBehaviourGet
		response         bool
	}{
		{
			name:     "OK",
			userID:   200,
			password: "Qwerty",
			mockBehaviourGet: func(r *mockDomain.MockUsersRepository, userID uint64) {
				p, err := utils.HashPassword("Qwerty")
				if err != nil {
					t.Error(err)
				}
				r.EXPECT().GetByID(userID).Return(&models.User{
					ID:       userID,
					Password: p,
				}, nil)
			},
			response: true,
		},
		{
			name:     "UserNotExist",
			userID:   200,
			password: "Qwerty",
			mockBehaviourGet: func(r *mockDomain.MockUsersRepository, userID uint64) {
				r.EXPECT().GetByID(userID).Return(nil, errors.New("user not found"))
			},
			response: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersRepository(ctrl)
			test.mockBehaviourGet(userMock, test.userID)

			usecase := New(userMock)

			isRight := usecase.CheckIDAndPassword(test.userID, test.password)

			require.Equal(t, isRight, test.response)
		})
	}
}

func TestUsecase_IsExistUsernameAndEmail(t *testing.T) {
	type mockBehaviourGetByUsername func(r *mockDomain.MockUsersRepository, username string)

	type mockBehaviourGetByEmail func(r *mockDomain.MockUsersRepository, email string)

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
			mockBehaviourGetByUsername: func(r *mockDomain.MockUsersRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(&models.User{
					Username: username,
					Email:    "a@d.com",
				}, nil)
			},
			mockBehaviourGetByEmail: func(r *mockDomain.MockUsersRepository, email string) {
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
			mockBehaviourGetByUsername: func(r *mockDomain.MockUsersRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(&models.User{}, errors.New("not exist"))
			},
			mockBehaviourGetByEmail: func(r *mockDomain.MockUsersRepository, email string) {},
			response:                false,
		},
		{
			name:     "IncorrectUsername",
			username: "user",
			email:    "a@d.com",
			mockBehaviourGetByUsername: func(r *mockDomain.MockUsersRepository, username string) {
				r.EXPECT().GetByUsername(username).Return(&models.User{
					Username: username,
					Email:    "adnsonjo@d.com",
				}, nil)
			},
			mockBehaviourGetByEmail: func(r *mockDomain.MockUsersRepository, email string) {
				r.EXPECT().GetByEmail(email).Return(&models.User{}, errors.New("not exist"))
			},
			response: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersRepository(ctrl)
			test.mockBehaviourGetByUsername(userMock, test.username)
			test.mockBehaviourGetByEmail(userMock, test.email)

			usecase := New(userMock)

			isExist := usecase.IsExistUsernameAndEmail(test.username, test.email)

			require.Equal(t, isExist, test.response)
		})
	}
}
