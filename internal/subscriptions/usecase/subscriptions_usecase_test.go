package subscriptions

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionsUsecase_GetSubscriptionsByUserID(t *testing.T) {
	type mockGetSub func(u *mockDomain.MockSubscriptionMicroservice, userID uint64)
	type mockGetImg func(u *mockDomain.MockImageUseCase, subImg, authorImg string)
	type mockGetUser func(u *mockDomain.MockUsersMicroservice, authorID uint64)

	tests := []struct {
		name        string
		userID      uint64
		subImg      string
		authorImg   string
		mockGetImg  mockGetImg
		mockGetSub  mockGetSub
		mockGetUser mockGetUser
		want        []models.AuthorSubscription
		wantErr     error
	}{
		{
			name:      "OK",
			userID:    1,
			subImg:    "subImg",
			authorImg: "authorImg",
			mockGetImg: func(u *mockDomain.MockImageUseCase, subImg, authorImg string) {
				u.EXPECT().GetImage(subImg).Return(subImg, nil)
				u.EXPECT().GetImage(authorImg).Return(authorImg, nil)
			},
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{
					{
						AuthorID:     1,
						Img:          "subImg",
						AuthorAvatar: "authorImg",
					},
				}, nil)
			},
			mockGetUser: func(u *mockDomain.MockUsersMicroservice, authorID uint64) {
				u.EXPECT().GetByID(authorID).Return(models.User{
					ID:       1,
					Username: "authorName",
					Avatar:   "authorImg",
				}, nil)
			},
			want: []models.AuthorSubscription{
				{
					AuthorID:     1,
					Img:          "subImg",
					AuthorName:   "authorName",
					AuthorAvatar: "authorImg",
				},
			},
		},
		{
			name:       "OKEmpty",
			userID:     1,
			mockGetImg: func(u *mockDomain.MockImageUseCase, subImg, authorImg string) {},
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{}, nil)
			},
			mockGetUser: func(u *mockDomain.MockUsersMicroservice, authorID uint64) {},
			want:        []models.AuthorSubscription{},
		},
		{
			name:       "ErrGetSub",
			userID:     1,
			mockGetImg: func(u *mockDomain.MockImageUseCase, subImg, authorImg string) {},
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{}, errors.New("error"))
			},
			mockGetUser: func(u *mockDomain.MockUsersMicroservice, authorID uint64) {},
			wantErr:     errors.New("error"),
		},
		{
			name:   "ErrGetImg1",
			userID: 1,
			subImg: "subImg",
			mockGetImg: func(u *mockDomain.MockImageUseCase, subImg, authorImg string) {
				u.EXPECT().GetImage(subImg).Return("", echo.NewHTTPError(http.StatusInternalServerError, "error"))
			},
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{
					{
						AuthorID:     1,
						Img:          "subImg",
						AuthorAvatar: "authorImg",
					},
				}, nil)
			},
			mockGetUser: func(u *mockDomain.MockUsersMicroservice, authorID uint64) {},
			wantErr:     errorHandling.WrapEcho(domain.ErrInternal, echo.NewHTTPError(http.StatusInternalServerError, "error")),
		},
		{
			name:      "ErrGetImg2",
			userID:    1,
			subImg:    "subImg",
			authorImg: "authorImg",
			mockGetImg: func(u *mockDomain.MockImageUseCase, subImg, authorImg string) {
				u.EXPECT().GetImage(subImg).Return(subImg, nil)
				u.EXPECT().GetImage(authorImg).Return("", echo.NewHTTPError(http.StatusInternalServerError, "error"))
			},
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{
					{
						AuthorID:     1,
						Img:          "subImg",
						AuthorAvatar: "authorImg",
					},
				}, nil)
			},
			mockGetUser: func(u *mockDomain.MockUsersMicroservice, authorID uint64) {
				u.EXPECT().GetByID(authorID).Return(models.User{
					ID:       1,
					Username: "authorName",
					Avatar:   "authorImg",
				}, nil)
			},
			wantErr: errorHandling.WrapEcho(domain.ErrInternal, echo.NewHTTPError(http.StatusInternalServerError, "error")),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockSub := mockDomain.NewMockSubscriptionMicroservice(ctrl)
			mockUser := mockDomain.NewMockUsersMicroservice(ctrl)
			mockImg := mockDomain.NewMockImageUseCase(ctrl)

			test.mockGetSub(mockSub, test.userID)
			test.mockGetUser(mockUser, test.userID)
			test.mockGetImg(mockImg, test.subImg, test.authorImg)

			u := New(mockSub, mockUser, mockImg)

			got, err := u.GetSubscriptionsByUserID(test.userID)
			if err != nil {
				assert.Equal(t, test.wantErr, err)
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestSubscriptionsUsecase_GetAuthorSubscriptionsByAuthorID(t *testing.T) {
	tests := []struct {
		name       string
		authorID   uint64
		authorImg  string
		mockGetSub func(u *mockDomain.MockSubscriptionMicroservice, authorID uint64)
		mockGetImg func(u *mockDomain.MockImageUseCase, subImg string)
		want       []models.AuthorSubscription
		wantErr    error
	}{
		{
			name:      "OK",
			authorID:  1,
			authorImg: "authorImg",
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, authorID uint64) {
				u.EXPECT().GetSubscriptionsByAuthorID(authorID).Return([]models.AuthorSubscription{
					{
						ID:       1,
						AuthorID: 1,
						Img:      "authorImg",
					},
				}, nil)
			},
			mockGetImg: func(u *mockDomain.MockImageUseCase, subImg string) {
				u.EXPECT().GetImage(subImg).Return(subImg, nil)
			},
			want: []models.AuthorSubscription{
				{
					ID:       1,
					AuthorID: 1,
					Img:      "authorImg",
				},
			},
		},
		{
			name:      "OKEmpty",
			authorID:  1,
			authorImg: "authorImg",
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, authorID uint64) {
				u.EXPECT().GetSubscriptionsByAuthorID(authorID).Return([]models.AuthorSubscription{}, nil)
			},
			mockGetImg: func(u *mockDomain.MockImageUseCase, subImg string) {},
			want:       []models.AuthorSubscription{},
		},
		{
			name:      "ErrGetSub",
			authorID:  1,
			authorImg: "authorImg",
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, authorID uint64) {
				u.EXPECT().GetSubscriptionsByAuthorID(authorID).Return([]models.AuthorSubscription{}, errors.New("error"))
			},
			mockGetImg: func(u *mockDomain.MockImageUseCase, subImg string) {},
			wantErr:    errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockSub := mockDomain.NewMockSubscriptionMicroservice(ctrl)
			mockImg := mockDomain.NewMockImageUseCase(ctrl)

			test.mockGetSub(mockSub, test.authorID)
			test.mockGetImg(mockImg, test.authorImg)

			u := New(mockSub, nil, mockImg)

			got, err := u.GetAuthorSubscriptionsByAuthorID(test.authorID)
			if err != nil {
				assert.Equal(t, test.wantErr, err)
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestSubscriptionsUsecase_GetSubscriptionsByID(t *testing.T) {
	tests := []struct {
		name       string
		subID      uint64
		authorID   uint64
		authorImg  string
		subImg     string
		mockGetSub func(u *mockDomain.MockSubscriptionMicroservice, subID uint64)
		mockGetImg func(u *mockDomain.MockImageUseCase, subImg string, authorImg string)
		want       models.AuthorSubscription
		wantErr    string
	}{
		{
			name:     "OK",
			subID:    2,
			authorID: 1,
			subImg:   "authorImg",
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, subID uint64) {
				u.EXPECT().GetSubscriptionByID(subID).Return(models.AuthorSubscription{
					AuthorID: 1,
					Img:      "authorImg",
					ID:       2,
				}, nil)
			},
			mockGetImg: func(u *mockDomain.MockImageUseCase, subImg string, authorImg string) {
				u.EXPECT().GetImage(subImg).Return(subImg, nil)
			},
			want: models.AuthorSubscription{
				AuthorID: 1,
				Img:      "authorImg",
				ID:       2,
			},
		},
		{
			name:     "ErrGetSub",
			subID:    2,
			authorID: 1,
			subImg:   "authorImg",
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, subID uint64) {
				u.EXPECT().GetSubscriptionByID(subID).Return(models.AuthorSubscription{}, errors.New("error"))
			},
			mockGetImg: func(u *mockDomain.MockImageUseCase, subImg string, authorImg string) {},
			wantErr:    "error",
		},
		{
			name:     "ErrGetImg",
			subID:    2,
			authorID: 1,
			subImg:   "authorImg",
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, subID uint64) {
				u.EXPECT().GetSubscriptionByID(subID).Return(models.AuthorSubscription{
					AuthorID: 1,
					Img:      "authorImg",
					ID:       2,
				}, nil)
			},
			mockGetImg: func(u *mockDomain.MockImageUseCase, subImg string, authorImg string) {
				u.EXPECT().GetImage(subImg).Return("", errors.New("error"))
			},
			wantErr: "code=500, message=server error, internal=error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockSub := mockDomain.NewMockSubscriptionMicroservice(ctrl)
			mockImg := mockDomain.NewMockImageUseCase(ctrl)

			test.mockGetSub(mockSub, test.subID)
			test.mockGetImg(mockImg, test.subImg, test.authorImg)

			u := New(mockSub, nil, mockImg)

			got, err := u.GetAuthorSubscriptionByID(test.subID)
			if err != nil {
				assert.Equal(t, test.wantErr, err.Error())
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestSubscriptionsUsecase_UpdateAuthorSubscription(t *testing.T) {
	tests := []struct {
		name       string
		sub        models.AuthorSubscription
		subID      uint64
		authorID   uint64
		authorImg  string
		subImg     string
		mockGetSub func(u *mockDomain.MockSubscriptionMicroservice, subID uint64)
		mockGetImg func(u *mockDomain.MockSubscriptionMicroservice, sub models.AuthorSubscription)
		want       models.AuthorSubscription
		wantErr    string
	}{
		{
			name: "OK",
			sub: models.AuthorSubscription{
				AuthorID: 1,
				Img:      "authorImg",
				ID:       2,
			},
			subID:    2,
			authorID: 1,
			subImg:   "authorImg",
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, subID uint64) {
				u.EXPECT().GetSubscriptionByID(subID).Return(models.AuthorSubscription{
					AuthorID: 1,
					Img:      "authorImg",
					ID:       2,
				}, nil)
			},
			mockGetImg: func(u *mockDomain.MockSubscriptionMicroservice, sub models.AuthorSubscription) {
				u.EXPECT().UpdateSubscription(sub).Return(nil)
			},
			want: models.AuthorSubscription{
				AuthorID: 1,
				Img:      "authorImg",
				ID:       2,
			},
		},
		{
			name: "ErrGetSub",
			sub: models.AuthorSubscription{
				AuthorID: 1,
				Img:      "authorImg",
				ID:       2,
			},
			subID:    2,
			authorID: 1,
			subImg:   "authorImg",
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, subID uint64) {
				u.EXPECT().GetSubscriptionByID(subID).Return(models.AuthorSubscription{}, errors.New("error"))
			},
			mockGetImg: func(u *mockDomain.MockSubscriptionMicroservice, sub models.AuthorSubscription) {},
			wantErr:    "error",
		},
		{
			name: "ErrGetImg",
			sub: models.AuthorSubscription{
				AuthorID: 1,
				Img:      "authorImg",
				ID:       2,
			},
			subID:    2,
			authorID: 1,
			subImg:   "authorImg",
			mockGetSub: func(u *mockDomain.MockSubscriptionMicroservice, subID uint64) {
				u.EXPECT().GetSubscriptionByID(subID).Return(models.AuthorSubscription{
					AuthorID: 1,
					Img:      "authorImg",
					ID:       2,
				}, nil)
			},
			mockGetImg: func(u *mockDomain.MockSubscriptionMicroservice, sub models.AuthorSubscription) {
				u.EXPECT().UpdateSubscription(sub).Return(errors.New("error"))
			},
			wantErr: "error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockSub := mockDomain.NewMockSubscriptionMicroservice(ctrl)

			test.mockGetSub(mockSub, test.subID)
			test.mockGetImg(mockSub, test.sub)

			u := New(mockSub, nil, nil)

			err := u.UpdateAuthorSubscription(test.sub, test.subID)
			if err != nil {
				assert.Equal(t, test.wantErr, err.Error())
			}
		})
	}
}
