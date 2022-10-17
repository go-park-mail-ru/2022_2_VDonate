package auth

import (
	"errors"
	"testing"

	mock_auth "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/auth/usecase"
	mock_users "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/users/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_Auth(t *testing.T) {
	type mockBehaviourDelete func(r *mock_auth.MockRepository, cookie string)

	tests := []struct {
		name                string
		sessionId           string
		mockBehaviourDelete mockBehaviourDelete
		response            bool
		responseError       error
	}{
		{
			name:      "OK",
			sessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			mockBehaviourDelete: func(r *mock_auth.MockRepository, cookie string) {
				r.EXPECT().GetBySessionID(cookie).Return(&models.Cookie{
					Value:  "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
					UserID: 22,
				}, nil)
			},
			response: true,
		},
		{
			name:      "NotFound",
			sessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			mockBehaviourDelete: func(r *mock_auth.MockRepository, cookie string) {
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

			userMock := mock_users.NewMockRepository(ctrl)
			authMock := mock_auth.NewMockRepository(ctrl)

			test.mockBehaviourDelete(authMock, test.sessionId)

			usecase := New(authMock, userMock)
			isAuth, err := usecase.Auth(test.sessionId)
			if err != nil {
				require.EqualError(t, err, test.responseError.Error())
			}

			require.Equal(t, test.response, isAuth)
		})
	}
}

func TestUsecase_Logout(t *testing.T) {
	type mockBehaviourDelete func(r *mock_auth.MockRepository, cookie string)

	tests := []struct {
		name                string
		sessionId           string
		mockBehaviourDelete mockBehaviourDelete
		response            bool
		responseError       error
	}{
		{
			name:      "OK",
			sessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			mockBehaviourDelete: func(r *mock_auth.MockRepository, cookie string) {
				r.EXPECT().DeleteBySessionID(cookie).Return(nil)
			},
			response: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mock_users.NewMockRepository(ctrl)
			authMock := mock_auth.NewMockRepository(ctrl)

			test.mockBehaviourDelete(authMock, test.sessionId)

			usecase := New(authMock, userMock)
			isAuth, err := usecase.Logout(test.sessionId)
			if err != nil {
				require.EqualError(t, err, test.responseError.Error())
			}

			require.Equal(t, test.response, isAuth)
		})
	}
}

func TestUsecase_IsSameSession(t *testing.T) {
	type mockBehaviourGet func(r *mock_users.MockRepository, cookie string)

	tests := []struct {
		name                string
		userId              uint64
		sessionId           string
		mockBehaviourGet mockBehaviourGet
		response            bool
	}{
		{
			name:      "OK",
			userId:    22,
			sessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			mockBehaviourGet: func(r *mock_users.MockRepository, cookie string) {
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

			userMock := mock_users.NewMockRepository(ctrl)
			authMock := mock_auth.NewMockRepository(ctrl)

			test.mockBehaviourGet(userMock, test.sessionId)

			usecase := New(authMock, userMock)
			isSame := usecase.IsSameSession(test.sessionId, test.userId)

			require.Equal(t, test.response, isSame)
		})
	}
}
