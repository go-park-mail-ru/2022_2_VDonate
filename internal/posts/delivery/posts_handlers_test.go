package httpPosts

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/posts/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHangler_GetPosts(t *testing.T) {
	type mockBehavior func(s *mock_posts.MockUseCase, postId uint64)

	tests := []struct {
		name                string
		method				string
		userId              uint64
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:   "OK",
			userId: 123,
			mockBehavior: func(s *mock_posts.MockUseCase, postId uint64) {
				s.EXPECT().GetPostsByUserID(postId).Return([]*models.PostDB{
					{
						Img:   "path/to/img",
						Title: "Look at this!!!",
						Text:  "Some text about my work",
					},
				}, nil)
			},
			expectedRequestBody: `[{"id":0,"user_id":0,"img":"path/to/img","title":"Look at this!!!","text":"Some text about my work"}]`,
		},
		{
			name:   "NotOK-BadId",
			userId: 123,
			mockBehavior: func(s *mock_posts.MockUseCase, postId uint64) {
				s.EXPECT().GetPostsByUserID(postId).Return([]*models.PostDB{}, errors.New("Bad request"))
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"message":"server error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mock_posts.NewMockUseCase(ctrl)
			test.mockBehavior(post, test.userId)

			handler := NewHandler(post)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.userId), 10))

			err := handler.GetPosts(c)
			require.NoError(t, err)

			body, _ := ioutil.ReadAll(rec.Body)

			assert.Equal(t, test.expectedRequestBody, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_CreatePost(t *testing.T) {
	type mockBehavior func(s *mock_posts.MockUseCase, postId models.PostDB)

	tests := []struct {
		name                string
		inputBody 			string
		inputPost 			models.PostDB
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedResponseBody string
	}{
		{
			name: "OK",
			inputBody: `{"title":"Title","text":"Text"}`,
			inputPost: models.PostDB{
				UserID: 100,
				Title: "Title",
				Text: "Text",
			},
			mockBehavior: func(s *mock_posts.MockUseCase, post models.PostDB) {
				s.EXPECT().Create(&post).Return(&models.PostDB{
					ID: 666,
					UserID: 100,
					Title: post.Title,
					Text: post.Text,
				}, nil)
			},
			expectedResponseBody: `{"id":666,"user_id":100,"img":"","title":"Title","text":"Text"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mock_posts.NewMockUseCase(ctrl)
			test.mockBehavior(post, test.inputPost)
			handler := NewHandler(post)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/posts/", bytes.NewBufferString(test.inputBody))
			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.inputPost.UserID), 10))

			err := handler.CreatePosts(c)
			require.NoError(t, err)

			body, _ := ioutil.ReadAll(rec.Body)

			assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
		})
	}
}
