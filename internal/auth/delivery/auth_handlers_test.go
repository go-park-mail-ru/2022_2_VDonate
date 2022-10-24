package httpAuth

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_SignUp(t *testing.T) {
	type mockBehavior func(r *mockDomain.MockAuthUseCase, user models.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"username","password":"qwerty","email":"ex@example.com"}`,
			inputUser: models.User{
				Username: "username",
				Password: "qwerty",
				Email:    "ex@example.com",
			},
			mockBehavior: func(r *mockDomain.MockAuthUseCase, user models.User) {
				r.EXPECT().SignUp(&user).Return("dsapfapspasf", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":0,"username":"username","email":"ex@example.com","is_author":false}`,
		},
		{
			name:      "IncorrectUser",
			inputBody: `{}`,
			inputUser: models.User{},
			mockBehavior: func(r *mockDomain.MockAuthUseCase, user models.User) {
				r.EXPECT().SignUp(&user).Return("", errors.New("empty user"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"no existing session"}`,
		},
		{
			name:                 "IncorrectBody",
			inputBody:            `mdaosmdop[23eomqwd`,
			inputUser:            models.User{},
			mockBehavior:         func(r *mockDomain.MockAuthUseCase, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Syntax error: offset=1, error=invalid character 'm' looking for beginning of value"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cntx := gomock.NewController(t)
			defer cntx.Finish()

			repo := mockDomain.NewMockAuthUseCase(cntx)
			user := mockDomain.NewMockUsersUseCase(cntx)
			test.mockBehavior(repo, test.inputUser)

			handler := NewHandler(repo, user)

			e := echo.New()
			e.Group("/api/v1")
			e.POST("/users", handler.SignUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(test.inputBody))
			req.Header.Set("Content-Type", "application/json")

			e.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, strings.Trim(w.Body.String(), "\n"))
		})
	}
}

func TestHandler_Login(t *testing.T) {
	type mockBehaviorLogin func(r *mockDomain.MockAuthUseCase, user models.AuthUser)

	type mockBehaviorUser func(r *mockDomain.MockUsersUseCase, sessionId string)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            models.AuthUser
		cookie               string
		mockBehaviorLogin    mockBehaviorLogin
		mockBehaviorUser     mockBehaviorUser
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"username","password":"qwerty"}`,
			inputUser: models.AuthUser{
				Username: "username",
				Password: "qwerty",
			},
			cookie: "jadbdsap324na4sa-4523sdnasodnoasdsna",
			mockBehaviorLogin: func(r *mockDomain.MockAuthUseCase, user models.AuthUser) {
				r.EXPECT().Login(user.Username, user.Password).Return("jadbdsap324na4sa-4523sdnasodnoasdsna", nil)
			},
			mockBehaviorUser: func(r *mockDomain.MockUsersUseCase, sessionId string) {
				r.EXPECT().GetBySessionID(sessionId).Return(&models.User{
					ID:        10,
					FirstName: "Jane",
					LastName:  "Doe",
					Email:     "john@email.com",
					Username:  "username",
					Password:  "qwerty",
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":10,"username":"username","first_name":"Jane","last_name":"Doe","email":"john@email.com","is_author":false}`,
		},
		{
			name:      "NoExistingSession-1",
			inputBody: `{"username":"username","password":"qwerty"}`,
			inputUser: models.AuthUser{
				Username: "username",
				Password: "qwerty",
			},
			cookie: "jadbdsap324na4sa-4523sdnasodnoasdsna",
			mockBehaviorLogin: func(r *mockDomain.MockAuthUseCase, user models.AuthUser) {
				r.EXPECT().Login(user.Username, user.Password).Return("", errors.New(""))
			},
			mockBehaviorUser:     func(r *mockDomain.MockUsersUseCase, sessionId string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"no existing session"}`,
		},
		{
			name:      "NoExistingSession-2",
			inputBody: `{"username":"username","password":"qwerty"}`,
			inputUser: models.AuthUser{
				Username: "username",
				Password: "qwerty",
			},
			cookie: "jadbdsap324na4sa-4523sdnasodnoasdsna",
			mockBehaviorLogin: func(r *mockDomain.MockAuthUseCase, user models.AuthUser) {
				r.EXPECT().Login(user.Username, user.Password).Return("jadbdsap324na4sa-4523sdnasodnoasdsna", nil)
			},
			mockBehaviorUser: func(r *mockDomain.MockUsersUseCase, sessionId string) {
				r.EXPECT().GetBySessionID(sessionId).Return(&models.User{}, errors.New("no user"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"failed to find item"}`,
		},
		{
			name:                 "BindError",
			inputBody:            `ksda[k[askd[aksd[a`,
			inputUser:            models.AuthUser{},
			mockBehaviorLogin:    func(r *mockDomain.MockAuthUseCase, user models.AuthUser) {},
			mockBehaviorUser:     func(r *mockDomain.MockUsersUseCase, sessionId string) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Syntax error: offset=1, error=invalid character 'k' looking for beginning of value"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cntx := gomock.NewController(t)
			defer cntx.Finish()

			repo := mockDomain.NewMockAuthUseCase(cntx)
			user := mockDomain.NewMockUsersUseCase(cntx)

			test.mockBehaviorLogin(repo, test.inputUser)
			test.mockBehaviorUser(user, test.cookie)

			handler := NewHandler(repo, user)

			e := echo.New()
			e.Group("/api/v1")
			e.POST("/login", handler.Login)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(test.inputBody))
			req.Header.Set("Content-Type", "application/json")

			e.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, strings.Trim(w.Body.String(), "\n"))
		})
	}
}

func TestHandler_Auth(t *testing.T) {
	type mockBehaviorAuth func(r *mockDomain.MockAuthUseCase, sessionId string)

	type mockBehaviorUser func(r *mockDomain.MockUsersUseCase, sessionId string)

	tests := []struct {
		name                 string
		cookie               string
		mockBehaviorAuth     mockBehaviorAuth
		mockBehaviorUser     mockBehaviorUser
		expectedErrorMessage string
		expectedResponseBody string
	}{
		{
			name:   "OK",
			cookie: "nadojads-dasasondfno312nnsandjo12",
			mockBehaviorAuth: func(r *mockDomain.MockAuthUseCase, sessionID string) {
				r.EXPECT().Auth(sessionID).Return(true, nil)
			},
			mockBehaviorUser: func(r *mockDomain.MockUsersUseCase, sessionID string) {
				r.EXPECT().GetBySessionID(sessionID).Return(&models.User{
					ID:       1,
					Email:    "ex@example.com",
					Username: "username",
				}, nil)
			},
			expectedResponseBody: `{"id":1,"username":"username","email":"ex@example.com","is_author":false}`,
		},
		{
			name:                 "NoExistingSession-1",
			cookie:               "",
			mockBehaviorAuth:     func(r *mockDomain.MockAuthUseCase, sessionID string) {},
			mockBehaviorUser:     func(r *mockDomain.MockUsersUseCase, sessionID string) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=http: named cookie not present",
		},
		{
			name:   "NoExistingSession-2",
			cookie: "nadojads-dasasondfno312nnsandjo12",
			mockBehaviorAuth: func(r *mockDomain.MockAuthUseCase, sessionID string) {
				r.EXPECT().Auth(sessionID).Return(false, nil)
			},
			mockBehaviorUser:     func(r *mockDomain.MockUsersUseCase, sessionID string) {},
			expectedErrorMessage: "code=401, message=failed to authenticate",
		},
		{
			name:   "NoExistingSession-3",
			cookie: "nadojads-dasasondfno312nnsandjo12",
			mockBehaviorAuth: func(r *mockDomain.MockAuthUseCase, sessionID string) {
				r.EXPECT().Auth(sessionID).Return(true, nil)
			},
			mockBehaviorUser: func(r *mockDomain.MockUsersUseCase, sessionID string) {
				r.EXPECT().GetBySessionID(sessionID).Return(&models.User{}, errors.New("no existing session"))
			},
			expectedErrorMessage: "code=404, message=failed to find item, internal=no existing session",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cntx := gomock.NewController(t)
			defer cntx.Finish()

			repo := mockDomain.NewMockAuthUseCase(cntx)
			user := mockDomain.NewMockUsersUseCase(cntx)

			test.mockBehaviorAuth(repo, test.cookie)
			test.mockBehaviorUser(user, test.cookie)

			handler := NewHandler(repo, user)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/auth")
			if len(test.cookie) > 0 {
				c.Request().Header.Add("Cookie", "session_id="+test.cookie)
			}

			err := handler.Auth(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_Logout(t *testing.T) {
	type mockBehaviorAuth func(r *mockDomain.MockAuthUseCase, sessionId string)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            models.AuthUser
		cookie               string
		mockBehaviorAuth     mockBehaviorAuth
		expectedResponseBody string
		expectedErrorMessage string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"username","password":"qwerty","email":"ex@example.com"}`,
			cookie:    "nadojads-dasasondfno312nnsandjo12",
			mockBehaviorAuth: func(r *mockDomain.MockAuthUseCase, sessionID string) {
				r.EXPECT().Logout(sessionID).Return(true, nil)
			},
			expectedResponseBody: `{}`,
		},
		{
			name:      "BadSession",
			inputBody: `{"username":"username","password":"qwerty","email":"ex@example.com"}`,
			cookie:    "nadojads-dasasondfno312nnsandjo12",
			mockBehaviorAuth: func(r *mockDomain.MockAuthUseCase, sessionID string) {
				r.EXPECT().Logout(sessionID).Return(false, nil)
			},
			expectedErrorMessage: "code=500, message=bad session",
		},
		{
			name:                 "NoSession",
			inputBody:            `{"username":"username","password":"qwerty","email":"ex@example.com"}`,
			cookie:               "",
			mockBehaviorAuth:     func(r *mockDomain.MockAuthUseCase, sessionID string) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=http: named cookie not present",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cntx := gomock.NewController(t)
			defer cntx.Finish()

			repo := mockDomain.NewMockAuthUseCase(cntx)
			user := mockDomain.NewMockUsersUseCase(cntx)

			test.mockBehaviorAuth(repo, test.cookie)

			handler := NewHandler(repo, user)

			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "https://127.0.0.1/api/v1/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/logout")
			if len(test.cookie) > 0 {
				c.Request().Header.Add("Cookie", "session_id="+test.cookie)
			}

			err := handler.Logout(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
		})
	}
}
