package httpPosts

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetPosts(t *testing.T) {
	type mockBehaviorGet func(s mockDomain.MockPostsUseCase, userID, authorID uint64)
	type mockBehaviorCookie func(s mockDomain.MockUsersUseCase, sessionID string)

	tests := []struct {
		name                 string
		cookie               string
		userID               int
		postID               uint64
		authorID             int64
		mockBehaviorGet      mockBehaviorGet
		mockBehaviorCookie   mockBehaviorCookie
		expectedRequestBody  string
		expectedErrorMessage string
	}{
		{
			name:     "OK",
			cookie:   "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			userID:   123,
			postID:   0,
			authorID: 123,
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID, authorID uint64) {
				s.EXPECT().GetPostsByFilter(userID, authorID).Return([]models.Post{
					{
						ID:       0,
						UserID:   123,
						LikesNum: 123,
						Content:  "content",
						Tier:     1000,
					},
				}, nil)
			},
			mockBehaviorCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 123,
				}, nil)
			},
			expectedRequestBody: `[{"postID":0,"userID":123,"contentTemplate":"","content":"content","tier":1000,"isAllowed":false,"dateCreated":"0001-01-01T00:00:00Z","tags":null,"author":{"userID":0,"username":"","imgPath":""},"likesNum":123,"isLiked":false}]`,
		},
		{
			name:   "OK-Empty",
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			userID: 123,
			postID: 0,
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID, authorID uint64) {
				s.EXPECT().GetPostsByFilter(userID, authorID).Return([]models.Post{}, nil)
			},
			mockBehaviorCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 123,
				}, nil)
			},
			expectedRequestBody: `[]`,
		},
		{
			name:   "ServerError",
			userID: 123,
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID, authorID uint64) {
				s.EXPECT().GetPostsByFilter(userID, authorID).Return([]models.Post{}, domain.ErrNotFound)
			},
			mockBehaviorCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 123,
				}, nil)
			},
			expectedErrorMessage: "code=404, message=failed to find item, internal=failed to find item",
		},
		{
			name:            "BadId",
			authorID:        -1,
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID, authorID uint64) {},
			mockBehaviorCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 123,
				}, nil)
			},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mockDomain.NewMockPostsUseCase(ctrl)
			users := mockDomain.NewMockUsersUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehaviorGet(*post, uint64(test.userID), uint64(test.authorID))
			test.mockBehaviorCookie(*users, test.cookie)

			handler := NewHandler(post, users, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/posts", nil)
			req.Header.Add("Cookie", "session_id="+test.cookie)
			v := req.URL.Query()
			v.Add("filter", fmt.Sprint(test.authorID))
			req.URL.RawQuery = v.Encode()
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts")

			err := handler.GetPosts(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedRequestBody, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHangler_GetPost(t *testing.T) {
	type mockBehaviorGet func(s mockDomain.MockPostsUseCase, postID, userID uint64)
	type mockBehaviorCookie func(s mockDomain.MockUsersUseCase, sessionID string)

	tests := []struct {
		name                 string
		method               string
		postID               int
		userID               uint64
		cookie               string
		mockBehaviorGet      mockBehaviorGet
		mockBehaviorCookie   mockBehaviorCookie
		expectedRequestBody  string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			userID: 123,
			postID: 123,
			mockBehaviorCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:       123,
					Username: "username",
				}, nil)
			},
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, postID, userID uint64) {
				s.EXPECT().GetPostByID(postID, userID).Return(models.Post{
					ID:       123,
					UserID:   123,
					LikesNum: 123,
					Content:  "content",
					Tier:     1000,
				}, nil)
			},
			expectedRequestBody: `{"postID":123,"userID":123,"contentTemplate":"","content":"content","tier":1000,"isAllowed":false,"dateCreated":"0001-01-01T00:00:00Z","tags":null,"author":{"userID":0,"username":"","imgPath":""},"likesNum":123,"isLiked":false}`,
		},
		{
			name:   "NotFound",
			postID: 123,
			userID: 123,
			mockBehaviorCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:       123,
					Username: "username",
				}, nil)
			},
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, postID, userID uint64) {
				s.EXPECT().GetPostByID(postID, userID).Return(models.Post{}, domain.ErrNotFound)
			},
			expectedErrorMessage: "code=404, message=failed to find item, internal=failed to find item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mockDomain.NewMockPostsUseCase(ctrl)
			users := mockDomain.NewMockUsersUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehaviorGet(*post, uint64(test.postID), test.userID)
			test.mockBehaviorCookie(*users, test.cookie)

			handler := NewHandler(post, users, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			req.Header.Add("Cookie", "session_id="+test.cookie)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/:postID")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.postID), 10))

			err := handler.GetPost(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedRequestBody, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_DeletePost(t *testing.T) {
	type mockDelete func(u *mockDomain.MockPostsUseCase, postID uint64)

	tests := []struct {
		name                 string
		postID               int
		mockDelete           mockDelete
		expectedRequestBody  string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			postID: 10,
			mockDelete: func(u *mockDomain.MockPostsUseCase, postID uint64) {
				u.EXPECT().DeleteByID(postID).Return(nil)
			},
			expectedRequestBody: "{}",
		},
		{
			name:                 "ErrBadRequest",
			postID:               -1,
			mockDelete:           func(u *mockDomain.MockPostsUseCase, postID uint64) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:   "ErrDelete",
			postID: 10,
			mockDelete: func(u *mockDomain.MockPostsUseCase, postID uint64) {
				u.EXPECT().DeleteByID(postID).Return(domain.ErrDelete)
			},
			expectedErrorMessage: "code=500, message=server error, internal=failed to delete item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mockDomain.NewMockPostsUseCase(ctrl)
			users := mockDomain.NewMockUsersUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockDelete(post, uint64(test.postID))

			handler := NewHandler(post, users, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/:postID")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.postID), 10))

			err := handler.DeletePost(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			} else {
				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				assert.Equal(t, test.expectedRequestBody, strings.Trim(string(body), "\n"))
			}
		})
	}
}

func TestHandler_CreatePost(t *testing.T) {
	type mockBehaviorCreate func(s *mockDomain.MockPostsUseCase, post models.Post, userID uint64)
	type mockBehaviorCookie func(s *mockDomain.MockUsersUseCase, sessionID string)

	tests := []struct {
		name                 string
		postID               int
		sessionID            string
		inputPost            models.Post
		mockBehaviorCreate   mockBehaviorCreate
		mockBehaviorCookie   mockBehaviorCookie
		expectedErrorMessage string
	}{
		{
			name:      "OK",
			postID:    10,
			sessionID: "session_id",
			inputPost: models.Post{},
			mockBehaviorCreate: func(s *mockDomain.MockPostsUseCase, post models.Post, userID uint64) {
				s.EXPECT().Create(post, userID).Return(models.Post{
					ID: 10,
				}, nil)
			},
			mockBehaviorCookie: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{}, nil)
			},
		},
		{
			name:               "BadRequest-Bind",
			postID:             10,
			sessionID:          "session_id",
			inputPost:          models.Post{},
			mockBehaviorCreate: func(s *mockDomain.MockPostsUseCase, post models.Post, userID uint64) {},
			mockBehaviorCookie: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{}, nil)
			},
			expectedErrorMessage: "code=400, message=bad request, internal=code=400, message=strconv.ParseUint: parsing \"�\": invalid syntax, internal=strconv.ParseUint: parsing \"�\": invalid syntax",
		},
		{
			name:      "ErrCreate",
			postID:    10,
			inputPost: models.Post{},
			mockBehaviorCreate: func(s *mockDomain.MockPostsUseCase, post models.Post, userID uint64) {
				s.EXPECT().Create(post, userID).Return(models.Post{}, domain.ErrCreate)
			},
			mockBehaviorCookie: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{}, nil)
			},
			expectedErrorMessage: "code=500, message=failed to create item, internal=failed to create item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mockDomain.NewMockPostsUseCase(ctrl)
			user := mockDomain.NewMockUsersUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehaviorCreate(post, test.inputPost, test.inputPost.UserID)
			test.mockBehaviorCookie(user, test.sessionID)

			handler := NewHandler(post, user, image)

			e := echo.New()
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)

			if test.name == "BadRequest-Bind" {
				err := writer.WriteField("userID", string(rune(-1)))
				assert.NoError(t, err)
			} else {
				err := writer.WriteField("userID", strconv.FormatUint(test.inputPost.UserID, 10))
				assert.NoError(t, err)
			}

			var formFile io.Writer

			formFile, err := writer.CreateFormFile("file", "../../../../test/test.txt")
			assert.NoError(t, err)

			sample, err := os.Open("../../../../test/test.txt")
			assert.NoError(t, err)

			_, err = io.Copy(formFile, sample)
			assert.NoError(t, err)

			err = writer.Close()
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/posts/", body)
			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Add("Cookie", "session_id="+test.sessionID)

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts/:post_id")

			if err = handler.CreatePost(c); err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}
		})
	}
}

func TestHandler_PutPost(t *testing.T) {
	type mockBehaviorPut func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64)

	tests := []struct {
		name                 string
		postID               int64
		inputPost            models.Post
		mockBehaviorPut      mockBehaviorPut
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			postID: 12,
			mockBehaviorPut: func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {
				s.EXPECT().Update(post, postID).Return(models.Post{
					ID: 12,
				}, nil)
			},
			inputPost: models.Post{},
		},
		{
			name:            "ErrBind",
			inputPost:       models.Post{},
			mockBehaviorPut: func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {},
			expectedErrorMessage: "code=400," +
				" message=bad request," +
				" internal=code=400," +
				" message=strconv.ParseUint: parsing \"�\": invalid syntax, internal=strconv.ParseUint: parsing \"�\": invalid syntax",
		},
		{
			name:      "ErrPut",
			inputPost: models.Post{},
			mockBehaviorPut: func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {
				s.EXPECT().Update(post, postID).Return(models.Post{}, domain.ErrUpdate)
			},
			expectedErrorMessage: "code=500, message=failed to update item, internal=failed to update item",
		},
		{
			name:                 "BadId",
			postID:               -1,
			mockBehaviorPut:      func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {},
			inputPost:            models.Post{},
			expectedErrorMessage: `code=400, message=bad request, internal=strconv.ParseUint: parsing "-1": invalid syntax`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mockDomain.NewMockPostsUseCase(ctrl)
			user := mockDomain.NewMockUsersUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehaviorPut(post, test.inputPost, uint64(test.postID))

			handler := NewHandler(post, user, image)

			e := echo.New()
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)

			if test.name == "ErrBind" {
				err := writer.WriteField("userID", string(rune(-1)))
				assert.NoError(t, err)
			} else {
				err := writer.WriteField("userID", strconv.FormatUint(test.inputPost.UserID, 10))
				assert.NoError(t, err)
			}

			formFile, err := writer.CreateFormFile("file", "../../../../test/test.txt")
			assert.NoError(t, err)

			sample, err := os.Open("../../../../test/test.txt")
			assert.NoError(t, err)

			_, err = io.Copy(formFile, sample)
			assert.NoError(t, err)

			err = writer.Close()
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/posts/", body)
			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", writer.FormDataContentType())

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts/:postID")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(test.postID, 10))

			if err = handler.PutPost(c); err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}
		})
	}
}

func TestHandler_GetLikes(t *testing.T) {
	type mockBehaviourLikes func(s *mockDomain.MockPostsUseCase, postID uint64)

	tests := []struct {
		name                 string
		postID               int64
		mockBehaviourLikes   mockBehaviourLikes
		expectedResponse     string
		expectedErrorMessage string
	}{
		{
			name:   "OK-1",
			postID: 23,
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesByPostID(postID).Return([]models.Like{
					{
						UserID: 12,
						PostID: 23,
					},
				}, nil)
			},
			expectedResponse: `[{"userID":12,"postID":23}]`,
		},
		{
			name:   "OK-2",
			postID: 12,
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesByPostID(postID).Return([]models.Like{}, nil)
			},
			expectedResponse: "[]",
		},
		{
			name:   "ErrNotFound",
			postID: 100,
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesByPostID(postID).Return([]models.Like{}, errors.New(""))
			},
			expectedErrorMessage: `code=404, message=failed to find item, internal=`,
		},
		{
			name:                 "BadRequest",
			postID:               -1,
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, postID uint64) {},
			expectedErrorMessage: `code=400, message=bad request, internal=strconv.ParseUint: parsing "-1": invalid syntax`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mockDomain.NewMockPostsUseCase(ctrl)
			users := mockDomain.NewMockUsersUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehaviourLikes(post, uint64(test.postID))

			handler := NewHandler(post, users, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "https://127.0.0.1/api/v1/posts/:id/likes", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts/:id/likes")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(test.postID, 10))

			err := handler.GetLikes(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponse, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_CreateLike(t *testing.T) {
	type mockBehaviourLikes func(s *mockDomain.MockPostsUseCase, userID, postID uint64)
	type mockBehaviourUsers func(s *mockDomain.MockUsersUseCase, sessionID string)

	tests := []struct {
		name                 string
		cookie               string
		postID               int64
		userID               uint64
		mockBehaviourUsers   mockBehaviourUsers
		mockBehaviourLikes   mockBehaviourLikes
		expectedResponse     string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			postID: 21,
			userID: 2,
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:       2,
					Username: "a",
					Email:    "a@a.ru",
				}, nil)
			},
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {
				s.EXPECT().LikePost(userID, postID).Return(nil)
			},
			expectedResponse: `{}`,
		},
		{
			name:   "ErrNoSession-1",
			cookie: "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			userID: 1,
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{}, errors.New("no session"))
			},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=no session",
		},
		{
			name:                 "ErrNoSession-2",
			cookie:               "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			userID:               1,
			mockBehaviourUsers:   func(s *mockDomain.MockUsersUseCase, sessionID string) {},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=http: named cookie not present",
		},
		{
			name:   "ErrLikeExist",
			cookie: "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			userID: 2,
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:       2,
					Username: "a",
					Email:    "a@a.ru",
				}, nil)
			},
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {
				s.EXPECT().LikePost(userID, postID).Return(errors.New(""))
			},
			expectedErrorMessage: "code=500, message=like alredy exist, internal=",
		},
		{
			name:                 "ErrBadId",
			postID:               -1,
			mockBehaviourUsers:   func(s *mockDomain.MockUsersUseCase, sessionID string) {},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: `code=400, message=bad request, internal=strconv.ParseUint: parsing "-1": invalid syntax`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mockDomain.NewMockPostsUseCase(ctrl)
			users := mockDomain.NewMockUsersUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehaviourUsers(users, test.cookie)
			test.mockBehaviourLikes(post, test.userID, uint64(test.postID))

			handler := NewHandler(post, users, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "https://127.0.0.1/api/v1/posts/:id/likes", nil)
			if test.name != "ErrNoSession-2" {
				req.Header.Add("Cookie", "session_id="+test.cookie)
			}
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts/:id/likes")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(test.postID, 10))

			err := handler.CreateLike(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponse, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_DeleteLike(t *testing.T) {
	type mockBehaviourUsers func(s *mockDomain.MockUsersUseCase, sessionID string)
	type mockBehaviourLikes func(s *mockDomain.MockPostsUseCase, userID, postID uint64)

	tests := []struct {
		name                 string
		postID               int64
		userID               uint64
		cookie               string
		mockBehaviourLikes   mockBehaviourLikes
		mockBehaviourUsers   mockBehaviourUsers
		expectedResponse     string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			postID: 3,
			userID: 21,
			cookie: "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:       21,
					Username: "user",
					Email:    "a@a.ru",
				}, nil)
			},
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {
				s.EXPECT().UnlikePost(userID, postID).Return(nil)
			},
			expectedResponse: `{}`,
		},
		{
			name:   "OK",
			postID: 3,
			userID: 21,
			cookie: "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:       21,
					Username: "user",
					Email:    "a@a.ru",
				}, nil)
			},
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {
				s.EXPECT().UnlikePost(userID, postID).Return(errors.New("like not found"))
			},
			expectedErrorMessage: "code=404, message=failed to find item, internal=like not found",
		},
		{
			name:   "ErrNoSession-1",
			postID: 3,
			userID: 21,
			cookie: "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{}, errors.New("no session"))
			},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=no session",
		},
		{
			name:                 "ErrNoSession-2",
			cookie:               "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			mockBehaviourUsers:   func(s *mockDomain.MockUsersUseCase, sessionID string) {},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=http: named cookie not present",
		},
		{
			name:                 "BadId",
			postID:               -1,
			mockBehaviourUsers:   func(s *mockDomain.MockUsersUseCase, sessionID string) {},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			users := mockDomain.NewMockUsersUseCase(ctrl)
			post := mockDomain.NewMockPostsUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehaviourUsers(users, test.cookie)
			test.mockBehaviourLikes(post, test.userID, uint64(test.postID))

			handler := NewHandler(post, users, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "https://127.0.0.1/api/v1/posts/:id/likes", nil)
			if test.name != "ErrNoSession-2" {
				req.Header.Add("Cookie", "session_id="+test.cookie)
			}
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts/:id/likes")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(test.postID, 10))

			err := handler.DeleteLike(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponse, strings.Trim(string(body), "\n"))
		})
	}
}
