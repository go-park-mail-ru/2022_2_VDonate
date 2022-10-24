package auth

import (
	"errors"
	"testing"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_Auth(t *testing.T) {
	type mockBehaviourDelete func(r *mockDomain.MockAuthRepository, cookie string)

	tests := []struct {
		name                string
		sessionID           string
		mockBehaviourDelete mockBehaviourDelete
		response            bool
		responseError       error
	}{
		{
			name:      "OK",
			sessionID: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			mockBehaviourDelete: func(r *mockDomain.MockAuthRepository, cookie string) {
				r.EXPECT().GetBySessionID(cookie).Return(&models.Cookie{
					Value:  "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
					UserID: 22,
				}, nil)
			},
			response: true,
		},
		{
			name:      "NotFound",
			sessionID: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			mockBehaviourDelete: func(r *mockDomain.MockAuthRepository, cookie string) {
				r.EXPECT().GetBySessionID(cookie).Return(&models.Cookie{}, errors.New("user not found"))
			},
			response:      false,
			responseError: errors.New("user not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersRepository(ctrl)
			authMock := mockDomain.NewMockAuthRepository(ctrl)

			test.mockBehaviourDelete(authMock, test.sessionID)

			usecase := New(authMock, userMock)
			isAuth, err := usecase.Auth(test.sessionID)
			if err != nil {
				require.EqualError(t, err, test.responseError.Error())
			}

			require.Equal(t, test.response, isAuth)
		})
	}
}

func TestUsecase_Logout(t *testing.T) {
	type mockBehaviourDelete func(r *mockDomain.MockAuthRepository, cookie string)

	tests := []struct {
		name                string
		sessionID           string
		mockBehaviourDelete mockBehaviourDelete
		response            bool
		responseError       error
	}{
		{
			name:      "OK",
			sessionID: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			mockBehaviourDelete: func(r *mockDomain.MockAuthRepository, cookie string) {
				r.EXPECT().DeleteBySessionID(cookie).Return(nil)
			},
			response: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersRepository(ctrl)
			authMock := mockDomain.NewMockAuthRepository(ctrl)

			test.mockBehaviourDelete(authMock, test.sessionID)

			usecase := New(authMock, userMock)
			isAuth, err := usecase.Logout(test.sessionID)
			if err != nil {
				require.EqualError(t, err, test.responseError.Error())
			}

			require.Equal(t, test.response, isAuth)
		})
	}
}

func TestUsecase_IsSameSession(t *testing.T) {
	type mockBehaviourGet func(r *mockDomain.MockUsersRepository, cookie string)

	tests := []struct {
		name             string
		userID           uint64
		sessionID        string
		mockBehaviourGet mockBehaviourGet
		response         bool
	}{
		{
			name:      "OK",
			userID:    22,
			sessionID: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			mockBehaviourGet: func(r *mockDomain.MockUsersRepository, cookie string) {
				r.EXPECT().GetBySessionID(cookie).Return(&models.User{
					ID: 22,
				}, nil)
			},
			response: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersRepository(ctrl)
			authMock := mockDomain.NewMockAuthRepository(ctrl)

			test.mockBehaviourGet(userMock, test.sessionID)

			usecase := New(authMock, userMock)
			isSame := usecase.IsSameSession(test.sessionID, test.userID)

			require.Equal(t, test.response, isSame)
		})
	}
}
