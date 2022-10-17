package httpPosts

import (
	"bytes"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	mock_domain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHangler_GetPosts(t *testing.T) {
	type mockBehaviorGet func(s mock_domain.MockPostsUseCase, userID uint64)
	type mockBehaviorUser func(s mock_domain.MockUsersUseCase, cookie string)

	tests := []struct {
		name                 string
		method               string
		userID               int
		cookie               string
		mockBehaviorGet      mockBehaviorGet
		mockBehaviorUser     mockBehaviorUser
		expectedRequestBody  string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			userID: 123,
			mockBehaviorUser: func(s mock_domain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(&models.User{
					ID: 123,
				}, nil)
			},
			mockBehaviorGet: func(s mock_domain.MockPostsUseCase, userID uint64) {
				s.EXPECT().GetPostsByUserID(userID).Return([]*models.Post{
					{
						UserID: userID,
						Img:    "path/to/img",
						Title:  "Look at this!!!",
						Text:   "Some text about my work",
					},
				}, nil)
			},
			expectedRequestBody: `[{"id":0,"user_id":123,"img":"path/to/img","title":"Look at this!!!","text":"Some text about my work"}]`,
		},
		{
			name:   "ServerError",
			userID: 123,
			mockBehaviorUser: func(s mock_domain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(&models.User{
					ID: 123,
				}, nil)
			},
			mockBehaviorGet: func(s mock_domain.MockPostsUseCase, userID uint64) {
				s.EXPECT().GetPostsByUserID(userID).Return([]*models.Post{}, domain.ErrNotFound)
			},
			expectedErrorMessage: "code=404, message=failed to find item, internal=failed to find item",
		},
		{
			name:   "BadId",
			userID: -1,
			mockBehaviorUser: func(s mock_domain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(nil, domain.ErrBadSession)
			},
			mockBehaviorGet:      func(s mock_domain.MockPostsUseCase, userID uint64) {},
			expectedErrorMessage: "code=500, message=bad session",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mock_domain.NewMockPostsUseCase(ctrl)
			users := mock_domain.NewMockUsersUseCase(ctrl)

			test.mockBehaviorUser(*users, test.cookie)
			test.mockBehaviorGet(*post, uint64(test.userID))

			handler := NewHandler(post, users)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/posts", nil)
			req.Header.Add("Cookie", "session_id="+test.cookie)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts")

			err := handler.GetPosts(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := ioutil.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedRequestBody, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHangler_GetPost(t *testing.T) {
	type mockBehaviorGet func(s mock_domain.MockPostsUseCase, userID uint64)

	tests := []struct {
		name                 string
		method               string
		postID               int
		mockBehaviorGet      mockBehaviorGet
		expectedRequestBody  string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			postID: 123,
			mockBehaviorGet: func(s mock_domain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetPostByID(postID).Return(&models.Post{
					Img:   "path/to/img",
					Title: "Look at this!!!",
					Text:  "Some text about my work",
				}, nil)
			},
			expectedRequestBody: `{"id":0,"user_id":0,"img":"path/to/img","title":"Look at this!!!","text":"Some text about my work"}`,
		},
		{
			name:   "NotFound",
			postID: 123,
			mockBehaviorGet: func(s mock_domain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetPostByID(postID).Return(nil, domain.ErrNotFound)
			},
			expectedErrorMessage: "code=404, message=failed to find item, internal=failed to find item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mock_domain.NewMockPostsUseCase(ctrl)
			users := mock_domain.NewMockUsersUseCase(ctrl)

			test.mockBehaviorGet(*post, uint64(test.postID))

			handler := NewHandler(post, users)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.postID), 10))

			err := handler.GetPost(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := ioutil.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedRequestBody, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_CreatePosts(t *testing.T) {
	type mockBehaviorUsers func(s *mock_domain.MockUsersUseCase, cookie string)
	type mockBehaviorCreate func(s *mock_domain.MockPostsUseCase, postID models.Post)

	tests := []struct {
		name                 string
		userID               int
		cookie               string
		inputBody            string
		inputPost            models.Post
		mockBehaviorUsers    mockBehaviorUsers
		mockBehaviorCreate   mockBehaviorCreate
		expectedStatusCode   int
		expectedResponseBody string
		expectedErrorMessage string
	}{
		{
			name:      "OK",
			cookie:    "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputBody: `{"title":"Title","text":"Text"}`,
			inputPost: models.Post{
				UserID: 100,
				Title:  "Title",
				Text:   "Text",
			},
			mockBehaviorUsers: func(s *mock_domain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(&models.User{
					ID: 100,
				}, nil)
			},
			mockBehaviorCreate: func(s *mock_domain.MockPostsUseCase, post models.Post) {
				s.EXPECT().Create(post).Return(&models.Post{
					ID:     0,
					UserID: 100,
					Title:  post.Title,
					Text:   post.Text,
				}, nil)
			},
			expectedResponseBody: `{"id":0,"user_id":100,"img":"","title":"Title","text":"Text"}`,
		},
		{
			name:                 "NoSession",
			userID:               -1,
			cookie:               "",
			inputBody:            `{"title":"Title","text":"Text"}`,
			inputPost:            models.Post{},
			mockBehaviorUsers:    func(s *mock_domain.MockUsersUseCase, cookie string) {},
			mockBehaviorCreate:   func(s *mock_domain.MockPostsUseCase, post models.Post) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=http: named cookie not present",
		},
		{
			name:      "NoSessionForUser",
			userID:    100,
			cookie:    "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputPost: models.Post{},
			mockBehaviorUsers: func(s *mock_domain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(nil, domain.ErrNoSession)
			},
			mockBehaviorCreate:   func(s *mock_domain.MockPostsUseCase, post models.Post) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=no existing session",
		},
		{
			name:      "ErrBind",
			userID:    100,
			cookie:    "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputBody: `NotJSON`,
			inputPost: models.Post{
				UserID: 100,
				Title:  "Title",
				Text:   "Text",
			},
			mockBehaviorUsers: func(s *mock_domain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(&models.User{
					ID: 100,
				}, nil)
			},
			mockBehaviorCreate: func(s *mock_domain.MockPostsUseCase, post models.Post) {},
			expectedErrorMessage: "code=400," +
				" message=bad request," +
				" internal=code=400," +
				" message=Syntax error: offset=1, error=invalid character 'N' looking for beginning of value, " +
				"internal=invalid character 'N' looking for beginning of value",
		},
		{
			name:      "ErrCreate",
			userID:    100,
			cookie:    "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputBody: `{"title":"Title","text":"Text"}`,
			inputPost: models.Post{
				UserID: 100,
				Title:  "Title",
				Text:   "Text",
			},
			mockBehaviorUsers: func(s *mock_domain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(&models.User{
					ID: 100,
				}, nil)
			},
			mockBehaviorCreate: func(s *mock_domain.MockPostsUseCase, post models.Post) {
				s.EXPECT().Create(post).Return(nil, domain.ErrCreate)
			},
			expectedErrorMessage: "code=500, message=failed to create item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mock_domain.NewMockPostsUseCase(ctrl)
			user := mock_domain.NewMockUsersUseCase(ctrl)

			test.mockBehaviorUsers(user, test.cookie)
			test.mockBehaviorCreate(post, test.inputPost)
			handler := NewHandler(post, user)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/posts/", bytes.NewBufferString(test.inputBody))
			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			if len(test.cookie) != 0 {
				req.Header.Add("Cookie", "session_id="+test.cookie)
			}

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.userID), 10))

			err := handler.CreatePosts(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := ioutil.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
		})
	}
}
