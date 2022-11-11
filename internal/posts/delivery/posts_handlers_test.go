package httpPosts

import (
	"bytes"
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
	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetPosts(t *testing.T) {
	type mockBehaviorGet func(s mockDomain.MockPostsUseCase, userID uint64)

	type mockBehaviorImage func(s mockDomain.MockImageUseCase, bucket, filename string)

	type mockBehaviourPost func(s mockDomain.MockPostsUseCase, postID uint64)

	tests := []struct {
		name                 string
		method               string
		userID               int
		postID               uint64
		cookie               string
		mockBehaviorGet      mockBehaviorGet
		mockBehaviorImage    mockBehaviorImage
		mockBehaviourPost    mockBehaviourPost
		expectedRequestBody  string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			userID: 123,
			postID: 0,
			mockBehaviorImage: func(s mockDomain.MockImageUseCase, bucket, filename string) {
				s.EXPECT().GetImage(bucket, filename).Return("", nil)
			},
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID uint64) {
				s.EXPECT().GetPostsByUserID(userID).Return([]models.Post{
					{
						UserID: userID,
						Img:    "path/to/img",
						Title:  "Look at this!!!",
						Text:   "Some text about my work",
					},
				}, nil)
			},
			mockBehaviourPost: func(s mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesNum(postID).Return(uint64(0), nil)
			},
			expectedRequestBody: `[{"postID":0,"userID":123,"img":"","title":"Look at this!!!","text":"Some text about my work","likesNum":0}]`,
		},
		{
			name:              "OK-Empty",
			cookie:            "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			userID:            123,
			postID:            0,
			mockBehaviorImage: func(s mockDomain.MockImageUseCase, bucket, filename string) {},
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID uint64) {
				s.EXPECT().GetPostsByUserID(userID).Return([]models.Post{}, nil)
			},
			mockBehaviourPost:   func(s mockDomain.MockPostsUseCase, postID uint64) {},
			expectedRequestBody: `{}`,
		},
		{
			name:   "ServerError",
			userID: 123,
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID uint64) {
				s.EXPECT().GetPostsByUserID(userID).Return([]models.Post{}, domain.ErrNotFound)
			},
			mockBehaviorImage:    func(s mockDomain.MockImageUseCase, bucket, filename string) {},
			mockBehaviourPost:    func(s mockDomain.MockPostsUseCase, postID uint64) {},
			expectedErrorMessage: "code=404, message=failed to find item, internal=failed to find item",
		},
		{
			name:                 "BadId",
			userID:               -1,
			mockBehaviorGet:      func(s mockDomain.MockPostsUseCase, userID uint64) {},
			mockBehaviorImage:    func(s mockDomain.MockImageUseCase, bucket, filename string) {},
			mockBehaviourPost:    func(s mockDomain.MockPostsUseCase, postID uint64) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:   "ErrInternal-Likes",
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			userID: 123,
			postID: 0,
			mockBehaviorImage: func(s mockDomain.MockImageUseCase, bucket, filename string) {
				s.EXPECT().GetImage(bucket, filename).Return("", nil)
			},
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID uint64) {
				s.EXPECT().GetPostsByUserID(userID).Return([]models.Post{
					{
						UserID: userID,
						Img:    "path/to/img",
						Title:  "Look at this!!!",
						Text:   "Some text about my work",
					},
				}, nil)
			},
			mockBehaviourPost: func(s mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesNum(postID).Return(uint64(0), domain.ErrInternal)
			},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
		},
		{
			name:   "ErrInternal-Likes",
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			userID: 123,
			postID: 0,
			mockBehaviorImage: func(s mockDomain.MockImageUseCase, bucket, filename string) {
				s.EXPECT().GetImage(bucket, filename).Return("", domain.ErrInternal)
			},
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID uint64) {
				s.EXPECT().GetPostsByUserID(userID).Return([]models.Post{
					{
						UserID: userID,
						Img:    "path/to/img",
						Title:  "Look at this!!!",
						Text:   "Some text about my work",
					},
				}, nil)
			},
			mockBehaviourPost:    func(s mockDomain.MockPostsUseCase, postID uint64) {},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			post := mockDomain.NewMockPostsUseCase(ctrl)
			users := mockDomain.NewMockUsersUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehaviorGet(*post, uint64(test.userID))
			test.mockBehaviorImage(*image, "image", "path/to/img")
			test.mockBehaviourPost(*post, test.postID)

			handler := NewHandler(post, users, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/posts", nil)
			req.Header.Add("Cookie", "session_id="+test.cookie)
			v := req.URL.Query()
			v.Add("user_id", fmt.Sprint(test.userID))
			req.URL.RawQuery = v.Encode()
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts")
			c.Set("bucket", "image")

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

func TestHangler_GetPost(t *testing.T) {
	type mockBehaviorGet func(s mockDomain.MockPostsUseCase, userID uint64)

	type mockBehaviorImage func(s mockDomain.MockImageUseCase, bucket, filename string)

	type mockBehaviourPost func(s mockDomain.MockPostsUseCase, postID uint64)

	tests := []struct {
		name                 string
		method               string
		postID               int
		mockBehaviorGet      mockBehaviorGet
		mockBehaviorImage    mockBehaviorImage
		mockBehaviourPost    mockBehaviourPost
		expectedRequestBody  string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			postID: 123,
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetPostByID(postID).Return(models.Post{
					Img:   "path/to/img",
					Title: "Look at this!!!",
					Text:  "Some text about my work",
				}, nil)
			},
			mockBehaviorImage: func(s mockDomain.MockImageUseCase, bucket, filename string) {
				s.EXPECT().GetImage(bucket, filename).Return("", nil)
			},
			mockBehaviourPost: func(s mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesNum(postID).Return(uint64(0), nil)
			},
			expectedRequestBody: `{"postID":0,"userID":0,"img":"","title":"Look at this!!!","text":"Some text about my work","likesNum":0}`,
		},
		{
			name:   "NotFound",
			postID: 123,
			mockBehaviorGet: func(s mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetPostByID(postID).Return(models.Post{}, domain.ErrNotFound)
			},
			mockBehaviorImage:    func(s mockDomain.MockImageUseCase, bucket, filename string) {},
			mockBehaviourPost:    func(s mockDomain.MockPostsUseCase, postID uint64) {},
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

			test.mockBehaviorGet(*post, uint64(test.postID))
			test.mockBehaviorImage(*image, "image", "path/to/img")
			test.mockBehaviourPost(*post, uint64(test.postID))

			handler := NewHandler(post, users, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/:postID")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.postID), 10))
			c.Set("bucket", "image")

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

func TestHandler_CreatePosts(t *testing.T) {
	type mockBehaviorPut func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64)

	type mockBehaviorImage func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context)

	type mockBehaviorGetImage func(s *mockDomain.MockImageUseCase, name, bucket string)

	tests := []struct {
		name                 string
		postID               int
		inputPost            models.Post
		mockBehaviorImage    mockBehaviorImage
		mockBehaviorPut      mockBehaviorPut
		mockBehaviorGetImage mockBehaviorGetImage
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			postID: 10,
			inputPost: models.Post{
				Title: "Title",
				Img:   "../../../test/test.txt",
				Text:  "Text",
			},
			mockBehaviorImage: func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {
				file, err := images.GetFileFromContext(c)
				assert.NoError(t, err)
				s.EXPECT().CreateImage(file, bucket).Return("../../../test/test.txt", nil)
			},
			mockBehaviorPut: func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {
				s.EXPECT().Update(post, postID).Return(nil)
			},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {
				s.EXPECT().GetImage(bucket, name).Return("path/to/img", nil)
			},
		},
		{
			name:   "BadRequest-ID",
			postID: -1,
			inputPost: models.Post{
				Title: "Title",
				Img:   "../../../test/test.txt",
				Text:  "Text",
			},
			mockBehaviorImage:    func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorPut:      func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:   "BadRequest-Bind",
			postID: 10,
			inputPost: models.Post{
				Title: "Title",
				Img:   "../../../test/test.txt",
				Text:  "Text",
			},
			mockBehaviorImage:    func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorPut:      func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {},
			expectedErrorMessage: "code=400, message=bad request, internal=code=400, message=strconv.ParseUint: parsing \"�\": invalid syntax, internal=strconv.ParseUint: parsing \"�\": invalid syntax",
		},
		{
			name:   "ErrCreate",
			postID: 10,
			inputPost: models.Post{
				Title: "Title",
				Img:   "../../../test/test.txt",
				Text:  "Text",
			},
			mockBehaviorImage: func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {
				file, err := images.GetFileFromContext(c)
				assert.NoError(t, err)
				s.EXPECT().CreateImage(file, bucket).Return("", domain.ErrCreate)
			},
			mockBehaviorPut:      func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {},
			expectedErrorMessage: "code=500, message=failed to create item, internal=failed to create item",
		},
		{
			name:   "Update",
			postID: 10,
			inputPost: models.Post{
				Title: "Title",
				Img:   "../../../test/test.txt",
				Text:  "Text",
			},
			mockBehaviorImage: func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {
				file, err := images.GetFileFromContext(c)
				assert.NoError(t, err)
				s.EXPECT().CreateImage(file, bucket).Return("../../../test/test.txt", nil)
			},
			mockBehaviorPut: func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {
				s.EXPECT().Update(post, postID).Return(domain.ErrUpdate)
			},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {},
			expectedErrorMessage: "code=500, message=failed to update item, internal=failed to update item",
		},
		{
			name:   "ErrInternal",
			postID: 10,
			inputPost: models.Post{
				Title: "Title",
				Img:   "../../../test/test.txt",
				Text:  "Text",
			},
			mockBehaviorImage: func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {
				file, err := images.GetFileFromContext(c)
				assert.NoError(t, err)
				s.EXPECT().CreateImage(file, bucket).Return("../../../test/test.txt", nil)
			},
			mockBehaviorPut: func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {
				s.EXPECT().Update(post, postID).Return(nil)
			},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {
				s.EXPECT().GetImage(bucket, name).Return("", domain.ErrInternal)
			},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
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
			test.mockBehaviorGetImage(image, "../../../test/test.txt", "image")

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

			formFile, err := writer.CreateFormFile("file", "../../../test/test.txt")
			assert.NoError(t, err)

			sample, err := os.Open("../../../test/test.txt")
			assert.NoError(t, err)

			_, err = io.Copy(formFile, sample)
			assert.NoError(t, err)

			err = writer.WriteField("img", test.inputPost.Img)
			assert.NoError(t, err)

			err = writer.WriteField("text", test.inputPost.Text)
			assert.NoError(t, err)

			err = writer.WriteField("title", test.inputPost.Title)
			assert.NoError(t, err)

			err = writer.Close()
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/posts/", body)
			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", writer.FormDataContentType())

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts/:post_id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.postID), 10))
			c.Set("bucket", "image")

			test.mockBehaviorImage(image, "image", c)

			if err = handler.PutPost(c); err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}
		})
	}
}

func TestHandler_PutPost(t *testing.T) {
	type mockBehaviorUsers func(s *mockDomain.MockUsersUseCase, cookie string)

	type mockBehaviorCreate func(s *mockDomain.MockPostsUseCase, postID models.Post)

	type mockBehaviorImage func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context)

	type mockBehaviorGetImage func(s *mockDomain.MockImageUseCase, bucket string)

	tests := []struct {
		name                 string
		userID               int
		cookie               string
		inputPost            models.Post
		mockBehaviorUsers    mockBehaviorUsers
		mockBehaviorCreate   mockBehaviorCreate
		mockBehaviorImage    mockBehaviorImage
		mockBehaviorGetImage mockBehaviorGetImage
		expectedStatusCode   int
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputPost: models.Post{
				UserID: 100,
				Title:  "Title",
				Img:    "../../../test/test.txt",
				Text:   "Text",
			},
			mockBehaviorUsers: func(s *mockDomain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(models.User{
					ID: 100,
				}, nil)
			},
			mockBehaviorImage: func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {
				file, err := images.GetFileFromContext(c)
				assert.NoError(t, err)
				s.EXPECT().CreateImage(file, bucket).Return("../../../test/test.txt", nil)
			},
			mockBehaviorCreate: func(s *mockDomain.MockPostsUseCase, post models.Post) {
				s.EXPECT().Create(post, post.UserID).Return(uint64(1), nil)
			},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, bucket string) {
				s.EXPECT().GetImage(bucket, "../../../test/test.txt").Return("", nil)
			},
		},
		{
			name:                 "NoSession",
			userID:               -1,
			cookie:               "",
			inputPost:            models.Post{},
			mockBehaviorUsers:    func(s *mockDomain.MockUsersUseCase, cookie string) {},
			mockBehaviorCreate:   func(s *mockDomain.MockPostsUseCase, post models.Post) {},
			mockBehaviorImage:    func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, bucket string) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=http: named cookie not present",
		},
		{
			name:      "NoSessionForUser",
			userID:    100,
			cookie:    "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputPost: models.Post{},
			mockBehaviorUsers: func(s *mockDomain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(models.User{}, domain.ErrNoSession)
			},
			mockBehaviorCreate:   func(s *mockDomain.MockPostsUseCase, post models.Post) {},
			mockBehaviorImage:    func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, bucket string) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=no existing session",
		},
		{
			name:      "ErrBind",
			userID:    100,
			cookie:    "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputPost: models.Post{},
			mockBehaviorUsers: func(s *mockDomain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(models.User{
					ID: 100,
				}, nil)
			},
			mockBehaviorImage:    func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorCreate:   func(s *mockDomain.MockPostsUseCase, post models.Post) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, bucket string) {},
			expectedErrorMessage: "code=400," +
				" message=bad request," +
				" internal=code=400," +
				" message=strconv.ParseUint: parsing \"�\": invalid syntax, internal=strconv.ParseUint: parsing \"�\": invalid syntax",
		},
		{
			name:   "ErrCreate",
			userID: 100,
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputPost: models.Post{
				UserID: 100,
				Title:  "Title",
				Img:    "../../../test/test.txt",
				Text:   "Text",
			},
			mockBehaviorUsers: func(s *mockDomain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(models.User{
					ID: 100,
				}, nil)
			},
			mockBehaviorImage: func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {
				file, err := images.GetFileFromContext(c)
				assert.NoError(t, err)
				s.EXPECT().CreateImage(file, bucket).Return("../../../test/test.txt", nil)
			},
			mockBehaviorCreate: func(s *mockDomain.MockPostsUseCase, post models.Post) {
				s.EXPECT().Create(post, post.UserID).Return(uint64(0), domain.ErrCreate)
			},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, bucket string) {},
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

			test.mockBehaviorUsers(user, test.cookie)
			test.mockBehaviorCreate(post, test.inputPost)

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

			formFile, err := writer.CreateFormFile("file", "../../../test/test.txt")
			assert.NoError(t, err)

			sample, err := os.Open("../../../test/test.txt")
			assert.NoError(t, err)

			_, err = io.Copy(formFile, sample)
			assert.NoError(t, err)

			err = writer.WriteField("img", test.inputPost.Img)
			assert.NoError(t, err)

			err = writer.WriteField("text", test.inputPost.Text)
			assert.NoError(t, err)

			err = writer.WriteField("title", test.inputPost.Title)
			assert.NoError(t, err)

			err = writer.Close()
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/posts/", body)
			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", writer.FormDataContentType())
			if len(test.cookie) != 0 {
				req.Header.Add("Cookie", "session_id="+test.cookie)
			}

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts/:postID")
			c.SetParamNames("postID")
			c.SetParamValues(strconv.FormatInt(int64(test.userID), 10))
			c.Set("bucket", "image")

			test.mockBehaviorImage(image, "image", c)
			test.mockBehaviorGetImage(image, "image")

			if err = handler.CreatePost(c); err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}
		})
	}
}
