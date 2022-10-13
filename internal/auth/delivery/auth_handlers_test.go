package httpAuth

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/users/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_SignUp(t *testing.T) {
	type mockBehavior func(r *mock_auth.MockUseCase, user models.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "SignUp-Ok",
			inputBody: `{"username":"username","password":"qwerty","email":"ex@example.com"}`,
			inputUser: models.User{
				Username: "username",
				Password: "qwerty",
				Email:    "ex@example.com",
			},
			mockBehavior: func(r *mock_auth.MockUseCase, user models.User) {
				r.EXPECT().SignUp(&user).Return("dsapfapspasf", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":0,"username":"username","email":"ex@example.com","is_author":false}`,
		},
		{
			name:      "SignUp-IncorrectInput",
			inputBody: `{}`,
			inputUser: models.User{},
			mockBehavior: func(r *mock_auth.MockUseCase, user models.User) {
				r.EXPECT().SignUp(&user).Return("", errors.New("Empty user"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"bad request"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cntx := gomock.NewController(t)
			defer cntx.Finish()

			repo := mock_auth.NewMockUseCase(cntx)
			user := mock_users.NewMockUseCase(cntx)
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

func TestHandler_login(t *testing.T) {
	type mockBehaviorLogin func(r *mock_auth.MockUseCase, user models.AuthUser)
	type mockBehaviorUser func(r *mock_users.MockUseCase, sessionId string)

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
			name:      "Login-Ok",
			inputBody: `{"username":"username","password":"qwerty"}`,
			inputUser: models.AuthUser{
				Username: "username",
				Password: "qwerty",
			},
			cookie: "jadbdsap324na4sa-4523sdnasodnoasdsna",
			mockBehaviorLogin: func(r *mock_auth.MockUseCase, user models.AuthUser) {
				r.EXPECT().Login(user.Username, user.Password).Return("jadbdsap324na4sa-4523sdnasodnoasdsna", nil)
			},
			mockBehaviorUser: func(r *mock_users.MockUseCase, sessionId string) {
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cntx := gomock.NewController(t)
			defer cntx.Finish()

			repo := mock_auth.NewMockUseCase(cntx)
			user := mock_users.NewMockUseCase(cntx)

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

func TestHandler_auth(t *testing.T) {
	type mockBehaviorAuth func(r *mock_auth.MockUseCase, sessionId string)
	type mockBehaviorUser func(r *mock_users.MockUseCase, sessionId string)

	tests := []struct {
		name                 string
		cookie               string
		mockBehaviorAuth     mockBehaviorAuth
		mockBehaviorUser     mockBehaviorUser
		expectedResponseBody string
	}{
		{
			name:      "Auth-Ok",
			cookie:    "nadojads-dasasondfno312nnsandjo12",
			mockBehaviorAuth: func(r *mock_auth.MockUseCase, session_id string) {
				r.EXPECT().Auth(session_id).Return(true, nil)
			},
			mockBehaviorUser: func(r *mock_users.MockUseCase, sessionId string) {
				r.EXPECT().GetBySessionID(sessionId).Return(&models.User{
					ID:       1,
					Email:    "ex@example.com",
					Username: "username",
				}, nil)
			},
			expectedResponseBody: `{"id":1,"username":"username","email":"ex@example.com","is_author":false}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cntx := gomock.NewController(t)
			defer cntx.Finish()

			repo := mock_auth.NewMockUseCase(cntx)
			user := mock_users.NewMockUseCase(cntx)

			test.mockBehaviorAuth(repo, test.cookie)
			test.mockBehaviorUser(user, test.cookie)

			handler := NewHandler(repo, user)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/auth")
			c.Request().Header.Add("Cookie", "session_id="+test.cookie)

			err := handler.Auth(c)
			require.NoError(t, err)
			body, _ := ioutil.ReadAll(rec.Body)

			assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_logout(t *testing.T) {
	type mockBehaviorAuth func(r *mock_auth.MockUseCase, sessionId string)
	type mockBehaviorUser func(r *mock_users.MockUseCase, sessionId string)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            models.AuthUser
		cookie               string
		mockBehaviorAuth     mockBehaviorAuth
		expectedResponseBody string
	}{
		{
			name:      "Logout-Ok",
			inputBody: `{"username":"username","password":"qwerty","email":"ex@example.com"}`,
			cookie:    "nadojads-dasasondfno312nnsandjo12",
			mockBehaviorAuth: func(r *mock_auth.MockUseCase, session_id string) {
				r.EXPECT().Logout(session_id).Return(true, nil)
			},
			expectedResponseBody: `{}`,
		},
		{
			name:      "Logout-NoSession",
			inputBody: `{"username":"username","password":"qwerty","email":"ex@example.com"}`,
			cookie:    "nadojads-dasasondfno312nnsandjo12",
			mockBehaviorAuth: func(r *mock_auth.MockUseCase, session_id string) {
				r.EXPECT().Logout(session_id).Return(false, nil)
			},
			expectedResponseBody: `{"message":"bad session"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cntx := gomock.NewController(t)
			defer cntx.Finish()

			repo := mock_auth.NewMockUseCase(cntx)
			user := mock_users.NewMockUseCase(cntx)

			test.mockBehaviorAuth(repo, test.cookie)

			handler := NewHandler(repo, user)

			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "https://127.0.0.1/api/v1/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/logout")
			c.Request().Header.Add("Cookie", "session_id="+test.cookie)

			err := handler.Logout(c)
			require.NoError(t, err)

			body, _ := ioutil.ReadAll(rec.Body)

			assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
		})
	}
}
