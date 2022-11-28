package auth

import (
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/interface"

	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_SignUp(t *testing.T) {
	type mockBehaviourUsername func(r *mockDomain.MockUsersRepository, username, password string)
	type mockBehaviourEmail func(r *mockDomain.MockUsersRepository, email, password string)
	type mockBehaviourCreate func(r *mockDomain.MockUsersRepository, user models.User)
	type mockBehaviourAuthRepo func(r *mockDomain.MockAuthRepository, cookie models.Cookie)

	tests := []struct {
		name                  string
		cookie                models.Cookie
		user                  models.User
		mockBehaviourUsername mockBehaviourUsername
		mockBehaviourEmail    mockBehaviourEmail
		mockBehaviourCreate   mockBehaviourCreate
		mockBehaviourAuthRepo mockBehaviourAuthRepo
		responseError         error
	}{
		{
			name: "OK",
			cookie: models.Cookie{
				Value: "cookie",
			},
			user: models.User{
				ID:       1,
				Username: "leo",
				Password: "abc",
			},
			mockBehaviourUsername: func(r *mockDomain.MockUsersRepository, username, password string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourEmail: func(r *mockDomain.MockUsersRepository, email, password string) {
				r.EXPECT().GetByEmail(email).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourCreate: func(r *mockDomain.MockUsersRepository, user models.User) {
				r.EXPECT().Create(user).Return(user.ID, nil)
			},
			mockBehaviourAuthRepo: func(r *mockDomain.MockAuthRepository, cookie models.Cookie) {
				r.EXPECT().CreateSession(cookie).Return(models.Cookie{
					Value: "cookie",
				}, nil)
			},
			responseError: nil,
		},
		{
			name: "ErrUsernameExists",
			cookie: models.Cookie{
				Value: "cookie",
			},
			user: models.User{
				Username: "leo",
				Password: "",
			},
			mockBehaviourUsername: func(r *mockDomain.MockUsersRepository, username, password string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, nil)
			},
			mockBehaviourEmail:    func(r *mockDomain.MockUsersRepository, email, password string) {},
			mockBehaviourCreate:   func(r *mockDomain.MockUsersRepository, user models.User) {},
			mockBehaviourAuthRepo: func(r *mockDomain.MockAuthRepository, cookie models.Cookie) {},
			responseError:         domain.ErrUsernameExist,
		},
		{
			name: "ErrEmailExists",
			cookie: models.Cookie{
				Value: "cookie",
			},
			user: models.User{
				Username: "leo",
				Password: "",
			},
			mockBehaviourUsername: func(r *mockDomain.MockUsersRepository, username, password string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourEmail: func(r *mockDomain.MockUsersRepository, email, password string) {
				r.EXPECT().GetByEmail(email).Return(models.User{}, nil)
			},
			mockBehaviourCreate:   func(r *mockDomain.MockUsersRepository, user models.User) {},
			mockBehaviourAuthRepo: func(r *mockDomain.MockAuthRepository, cookie models.Cookie) {},
			responseError:         domain.ErrEmailExist,
		},
		{
			name: "ErrInternal-CantHashPassword",
			cookie: models.Cookie{
				Value: "cookie",
			},
			user: models.User{
				Username: "leo",
				Password: "",
			},
			mockBehaviourUsername: func(r *mockDomain.MockUsersRepository, username, password string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourEmail: func(r *mockDomain.MockUsersRepository, email, password string) {
				r.EXPECT().GetByEmail(email).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourCreate:   func(r *mockDomain.MockUsersRepository, user models.User) {},
			mockBehaviourAuthRepo: func(r *mockDomain.MockAuthRepository, cookie models.Cookie) {},
			responseError:         domain.ErrInternal,
		},
		{
			name: "ErrInternal-Create",
			cookie: models.Cookie{
				Value: "cookie",
			},
			user: models.User{
				Username: "leo",
				Password: "abc",
			},
			mockBehaviourUsername: func(r *mockDomain.MockUsersRepository, username, password string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourEmail: func(r *mockDomain.MockUsersRepository, email, password string) {
				r.EXPECT().GetByEmail(email).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourCreate: func(r *mockDomain.MockUsersRepository, user models.User) {
				r.EXPECT().Create(user).Return(user.ID, domain.ErrCreate)
			},
			mockBehaviourAuthRepo: func(r *mockDomain.MockAuthRepository, cookie models.Cookie) {},
			responseError:         domain.ErrCreate,
		},
		{
			name: "ErrInternal-ErrorCreateSession",
			cookie: models.Cookie{
				Value: "cookie",
			},
			user: models.User{
				Username: "leo",
				Password: "abc",
			},
			mockBehaviourUsername: func(r *mockDomain.MockUsersRepository, username, password string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourEmail: func(r *mockDomain.MockUsersRepository, email, password string) {
				r.EXPECT().GetByEmail(email).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourCreate: func(r *mockDomain.MockUsersRepository, user models.User) {
				r.EXPECT().Create(user).Return(user.ID, nil)
			},
			mockBehaviourAuthRepo: func(r *mockDomain.MockAuthRepository, cookie models.Cookie) {
				r.EXPECT().CreateSession(cookie).Return(models.Cookie{}, domain.ErrBadSession)
			},
			responseError: domain.ErrBadSession,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersRepository(ctrl)
			authMock := mockDomain.NewMockAuthRepository(ctrl)

			test.mockBehaviourUsername(userMock, test.user.Username, test.user.Password)
			test.mockBehaviourEmail(userMock, test.user.Email, test.user.Password)
			test.mockBehaviourCreate(userMock, test.user)
			test.mockBehaviourAuthRepo(authMock, test.cookie)

			usecase := WithCreators(
				authMock,
				userMock,
				func(id uint64) models.Cookie {
					return test.cookie
				},
				func(password string) (string, error) {
					if len(password) == 0 {
						return "", domain.ErrInternal
					}
					return password, nil
				},
			)
			sessionID, err := usecase.SignUp(test.user)
			if err != nil {
				require.EqualError(t, err, test.responseError.Error())
			} else {
				require.Equal(t, sessionID, test.cookie.Value)
			}
		})
	}
}

func TestUsecase_Login(t *testing.T) {
	type mockBehaviourUsername func(r *mockDomain.MockUsersRepository, username, password string)
	type mockBehaviourEmail func(r *mockDomain.MockUsersRepository, email, password string)
	type mockBehaviourAuthRepo func(r *mockDomain.MockAuthRepository, cookie models.Cookie)

	tests := []struct {
		name                  string
		cookie                models.Cookie
		username              string
		password              string
		mockBehaviourUsername mockBehaviourUsername
		mockBehaviourEmail    mockBehaviourEmail
		mockBehaviourAuthRepo mockBehaviourAuthRepo
		responseError         error
	}{
		{
			name: "OK",
			cookie: models.Cookie{
				Value: "cookie",
			},
			username: "leo",
			password: "*****",
			mockBehaviourUsername: func(r *mockDomain.MockUsersRepository, username, password string) {
				p, err := utils.HashPassword(password)
				assert.NoError(t, err)
				r.EXPECT().GetByUsername(username).Return(models.User{ID: 12, Password: p}, nil)
			},
			mockBehaviourEmail: func(r *mockDomain.MockUsersRepository, email, password string) {},
			mockBehaviourAuthRepo: func(r *mockDomain.MockAuthRepository, cookie models.Cookie) {
				r.EXPECT().CreateSession(cookie).Return(models.Cookie{
					Value: "cookie",
				}, nil)
			},
			responseError: nil,
		},
		{
			name: "OK-WithEmail",
			cookie: models.Cookie{
				Value: "cookie",
			},
			username: "leo@gmail.com",
			password: "*****",
			mockBehaviourUsername: func(r *mockDomain.MockUsersRepository, username, password string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourEmail: func(r *mockDomain.MockUsersRepository, email, password string) {
				p, err := utils.HashPassword(password)
				assert.NoError(t, err)
				r.EXPECT().GetByEmail(email).Return(models.User{
					ID:       12,
					Password: p,
				}, nil)
			},
			mockBehaviourAuthRepo: func(r *mockDomain.MockAuthRepository, cookie models.Cookie) {
				r.EXPECT().CreateSession(cookie).Return(models.Cookie{
					Value: "cookie",
				}, nil)
			},
			responseError: nil,
		},
		{
			name: "NoUsernameOrEmail",
			cookie: models.Cookie{
				Value: "cookie",
			},
			username: "leo@gmail.com",
			password: "*****",
			mockBehaviourUsername: func(r *mockDomain.MockUsersRepository, username, password string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourEmail: func(r *mockDomain.MockUsersRepository, email, password string) {
				r.EXPECT().GetByEmail(email).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourAuthRepo: func(r *mockDomain.MockAuthRepository, cookie models.Cookie) {},
			responseError:         domain.ErrUsernameOrEmailNotExist,
		},
		{
			name: "NotEqualPasswords",
			cookie: models.Cookie{
				Value: "cookie",
			},
			username: "leo@gmail.com",
			password: "*****",
			mockBehaviourUsername: func(r *mockDomain.MockUsersRepository, username, password string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourEmail: func(r *mockDomain.MockUsersRepository, email, password string) {
				r.EXPECT().GetByEmail(email).Return(models.User{
					ID: 12, Password: password,
				}, nil)
			},
			mockBehaviourAuthRepo: func(r *mockDomain.MockAuthRepository, cookie models.Cookie) {},
			responseError:         domain.ErrPasswordsNotEqual,
		},
		{
			name: "InternalError-CreateSession",
			cookie: models.Cookie{
				Value: "cookie",
			},
			username: "leo@gmail.com",
			password: "*****",
			mockBehaviourUsername: func(r *mockDomain.MockUsersRepository, username, password string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, domain.ErrNotFound)
			},
			mockBehaviourEmail: func(r *mockDomain.MockUsersRepository, email, password string) {
				p, err := utils.HashPassword(password)
				assert.NoError(t, err)
				r.EXPECT().GetByEmail(email).Return(models.User{
					ID: 12, Password: p,
				}, nil)
			},
			mockBehaviourAuthRepo: func(r *mockDomain.MockAuthRepository, cookie models.Cookie) {
				r.EXPECT().CreateSession(cookie).Return(models.Cookie{}, domain.ErrInternal)
			},
			responseError: domain.ErrInternal,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersRepository(ctrl)
			authMock := mockDomain.NewMockAuthRepository(ctrl)

			test.mockBehaviourUsername(userMock, test.username, test.password)
			test.mockBehaviourEmail(userMock, test.username, test.password)
			test.mockBehaviourAuthRepo(authMock, test.cookie)

			usecase := WithCookieCreator(authMock, userMock, func(id uint64) models.Cookie {
				return test.cookie
			})
			sessionID, err := usecase.Login(test.username, test.password)
			if err != nil {
				require.EqualError(t, err, test.responseError.Error())
			} else {
				require.Equal(t, sessionID, test.cookie.Value)
			}
		})
	}
}

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
				r.EXPECT().GetBySessionID(cookie).Return(models.Cookie{
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
				r.EXPECT().GetBySessionID(cookie).Return(models.Cookie{}, errors.New("user not found"))
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
		{
			name:      "ErrDelete",
			sessionID: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			mockBehaviourDelete: func(r *mockDomain.MockAuthRepository, cookie string) {
				r.EXPECT().DeleteBySessionID(cookie).Return(domain.ErrDelete)
			},
			responseError: domain.ErrDelete,
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
				assert.EqualError(t, err, test.responseError.Error())
			} else {
				assert.Equal(t, test.response, isAuth)
			}
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
		responseError    error
	}{
		{
			name:      "OK",
			userID:    22,
			sessionID: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			mockBehaviourGet: func(r *mockDomain.MockUsersRepository, cookie string) {
				r.EXPECT().GetBySessionID(cookie).Return(models.User{
					ID: 22,
				}, nil)
			},
			response: true,
		},
		{
			name:      "UserNotFound",
			userID:    22,
			sessionID: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			mockBehaviourGet: func(r *mockDomain.MockUsersRepository, cookie string) {
				r.EXPECT().GetBySessionID(cookie).Return(models.User{}, domain.ErrNotFound)
			},
			response: false,
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
