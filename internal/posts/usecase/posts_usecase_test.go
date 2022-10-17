package posts

import (
	mock_domain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_GetPostsByUserID(t *testing.T) {
	type mockBehaviour func(s *mock_domain.MockPostsRepository, userID uint64)

	tests := []struct {
		name                 string
		userID               uint64
		mockBehaviour        mockBehaviour
		response             []*models.Post
		responseErrorMessage string
	}{
		{
			name:   "OK",
			userID: 200,
			mockBehaviour: func(s *mock_domain.MockPostsRepository, userID uint64) {
				s.EXPECT().GetAllByUserID(userID).Return([]*models.Post{
					{
						ID:     1,
						UserID: 200,
						Title:  "title",
						Text:   "text",
					},
				}, nil)
			},
			response: []*models.Post{
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
			mockBehaviour: func(s *mock_domain.MockPostsRepository, userID uint64) {
				s.EXPECT().GetAllByUserID(userID).Return([]*models.Post{}, nil)
			},
			responseErrorMessage: "no posts were found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			postMock := mock_domain.NewMockPostsRepository(ctrl)

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
