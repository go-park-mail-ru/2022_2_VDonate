package posts

import (
	"testing"

	mock_posts "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/posts/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_GetPostsByUserId(t *testing.T) {
	type mockBehaviour func(s *mock_posts.MockRepository, userId uint64)

	tests := []struct {
		name          string
		userId        uint64
		mockBehaviour mockBehaviour
		err           string
		response      []*models.PostDB
	}{
		{
			name:   "OK",
			userId: 200,
			mockBehaviour: func(s *mock_posts.MockRepository, userId uint64) {
				s.EXPECT().GetAllByUserID(userId).Return([]*models.PostDB{
					{
						ID:     1,
						UserID: 200,
						Title:  "title",
						Text:   "text",
					},
				}, nil)
			},
			err: "",
			response: []*models.PostDB{
				{
					ID:     1,
					UserID: 200,
					Title:  "title",
					Text:   "text",
				},
			},
		},
		{
			name:   "Not-OK",
			userId: 200,
			mockBehaviour: func(s *mock_posts.MockRepository, userId uint64) {
				s.EXPECT().GetAllByUserID(userId).Return([]*models.PostDB{}, nil)
			},
			err: "no posts",
			response: []*models.PostDB(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			postMock := mock_posts.NewMockRepository(ctrl)
			test.mockBehaviour(postMock, test.userId)

			usecase := New(postMock)
			post, err := usecase.GetPostsByUserID(test.userId)
			if err != nil {
				require.EqualError(t, err, test.err)
			}

			require.Equal(t, test.response, post)
		})
	}
}
