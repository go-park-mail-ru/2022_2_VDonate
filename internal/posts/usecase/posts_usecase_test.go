package posts

import (
	"errors"
	"testing"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsecase_GetPostsByUserID(t *testing.T) {
	type mockBehaviour func(s *mockDomain.MockPostsRepository, userID uint64)

	tests := []struct {
		name                 string
		userID               uint64
		mockBehaviour        mockBehaviour
		response             []models.Post
		responseErrorMessage string
	}{
		{
			name:   "OK",
			userID: 200,
			mockBehaviour: func(s *mockDomain.MockPostsRepository, userID uint64) {
				s.EXPECT().GetAllByUserID(userID).Return([]models.Post{
					{
						ID:     1,
						UserID: 200,
						Title:  "title",
						Text:   "text",
					},
				}, nil)
			},
			response: []models.Post{
				{
					ID:     1,
					UserID: 200,
					Title:  "title",
					Text:   "text",
				},
			},
		},
		{
			name:   "NoPosts",
			userID: 200,
			mockBehaviour: func(s *mockDomain.MockPostsRepository, userID uint64) {
				s.EXPECT().GetAllByUserID(userID).Return([]models.Post{}, nil)
			},
			response:             []models.Post{},
			responseErrorMessage: "no posts were found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			postMock := mockDomain.NewMockPostsRepository(ctrl)

			test.mockBehaviour(postMock, test.userID)

			usecase := New(postMock)
			post, err := usecase.GetPostsByUserID(test.userID)
			if err != nil {
				assert.Equal(t, test.responseErrorMessage, err.Error())
			}

			require.Equal(t, test.response, post)
		})
	}
}

func TestUsecase_GetLikesNum(t *testing.T) {
	type mockBehaviourGetAllLikes func(s *mockDomain.MockPostsRepository, postID uint64)

	tests := []struct {
		name                        string
		postID                      uint64
		mockBehaviourGetAllLikes    mockBehaviourGetAllLikes
		expectedResponse            uint64
		expectedErrorMessage        string
	}{
		{
			name:   "OK",
			postID: 100,
			mockBehaviourGetAllLikes: func(s *mockDomain.MockPostsRepository, postID uint64) {
				s.EXPECT().GetAllLikesByPostID(postID).Return([]models.Like{
					{
						UserID: 12,
						PostID: 100,
					},
				}, nil)
			},
			expectedResponse: 1,
		},
		{
			name:   "ErrNotFound",
			postID: 100,
			mockBehaviourGetAllLikes: func(s *mockDomain.MockPostsRepository, postID uint64) {
				s.EXPECT().GetAllLikesByPostID(postID).Return([]models.Like{}, errors.New("likes not found"))
			},
			expectedErrorMessage: "likes not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			postMockRepo := mockDomain.NewMockPostsRepository(ctrl)

			test.mockBehaviourGetAllLikes(postMockRepo, test.postID)

			usecase := New(postMockRepo)
			post, err := usecase.GetLikesNum(test.postID)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			require.Equal(t, test.expectedResponse, post)
		})
	}
}
