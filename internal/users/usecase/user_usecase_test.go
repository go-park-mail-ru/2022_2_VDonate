package users

import (
	"errors"
	"mime/multipart"
	"testing"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_Update(t *testing.T) {
	type mockBehaviourGet func(r *mockDomain.MockUsersMicroservice, userID uint64)
	type mockBehaviourUpdate func(r *mockDomain.MockUsersMicroservice, user models.User)
	type mockImage func(r *mockDomain.MockImageUseCase, file *multipart.FileHeader, avatar string)

	tests := []struct {
		name                string
		inputUser           models.User
		mockBehaviourGet    mockBehaviourGet
		mockBehaviourUpdate mockBehaviourUpdate
		mockImage           mockImage
		responseError       string
	}{
		{
			name: "OK",
			inputUser: models.User{
				ID:       200,
				Username: "user",
				Password: "abc",
			},
			mockBehaviourGet: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetByID(userID).Return(models.User{
					ID:       200,
					Username: "user",
					Password: "abc",
				}, nil)
			},
			mockBehaviourUpdate: func(r *mockDomain.MockUsersMicroservice, user models.User) {
				r.EXPECT().Update(user).Return(nil)
			},
			mockImage: func(r *mockDomain.MockImageUseCase, file *multipart.FileHeader, avatar string) {
				r.EXPECT().CreateOrUpdateImage(file, avatar).Return("", nil)
			},
		},
		{
			name: "NotFound",
			inputUser: models.User{
				ID:       200,
				Username: "user",
				Password: "abc",
			},
			mockBehaviourGet: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetByID(userID).Return(models.User{}, errors.New("not found"))
			},
			mockBehaviourUpdate: func(r *mockDomain.MockUsersMicroservice, user models.User) {},
			mockImage:           func(r *mockDomain.MockImageUseCase, file *multipart.FileHeader, avatar string) {},
			responseError:       "not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersMicroservice(ctrl)
			imgMock := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehaviourGet(userMock, test.inputUser.ID)
			test.mockBehaviourUpdate(userMock, test.inputUser)

			test.mockImage(imgMock, &multipart.FileHeader{}, "")

			usecase := WithHashCreator(
				userMock,
				imgMock,
				func(password string) (string, error) {
					if len(password) == 0 {
						return "", errors.New("empty password")
					}
					return password, nil
				},
			)

			_, err := usecase.Update(test.inputUser, &multipart.FileHeader{}, test.inputUser.ID)
			if err != nil {
				require.EqualError(t, err, test.responseError)
			}
		})
	}
}

func TestUsecase_CheckIDAndPassword(t *testing.T) {
	type mockBehaviourGet func(r *mockDomain.MockUsersMicroservice, userID uint64)

	tests := []struct {
		name             string
		userID           uint64
		password         string
		mockBehaviourGet mockBehaviourGet
		response         bool
	}{
		{
			name:     "OK",
			userID:   200,
			password: "Qwerty",
			mockBehaviourGet: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				p, err := utils.HashPassword("Qwerty")
				if err != nil {
					t.Error(err)
				}
				r.EXPECT().GetByID(userID).Return(models.User{
					ID:       userID,
					Password: p,
				}, nil)
			},
			response: true,
		},
		{
			name:     "UserNotExist",
			userID:   200,
			password: "Qwerty",
			mockBehaviourGet: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetByID(userID).Return(models.User{}, errors.New("user not found"))
			},
			response: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersMicroservice(ctrl)
			imgMock := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehaviourGet(userMock, test.userID)

			usecase := New(userMock, imgMock)

			isRight := usecase.CheckIDAndPassword(test.userID, test.password)

			require.Equal(t, isRight, test.response)
		})
	}
}

func TestUsecase_IsExistUsernameAndEmail(t *testing.T) {
	type mockBehaviourGetByUsername func(r *mockDomain.MockUsersMicroservice, username string)

	type mockBehaviourGetByEmail func(r *mockDomain.MockUsersMicroservice, email string)

	tests := []struct {
		name                       string
		username                   string
		email                      string
		mockBehaviourGetByUsername mockBehaviourGetByUsername
		mockBehaviourGetByEmail    mockBehaviourGetByEmail
		response                   bool
	}{
		{
			name:     "OK",
			username: "user",
			email:    "a@d.com",
			mockBehaviourGetByUsername: func(r *mockDomain.MockUsersMicroservice, username string) {
				r.EXPECT().GetByUsername(username).Return(models.User{
					Username: username,
					Email:    "a@d.com",
				}, nil)
			},
			mockBehaviourGetByEmail: func(r *mockDomain.MockUsersMicroservice, email string) {
				r.EXPECT().GetByEmail(email).Return(models.User{
					Username: "user",
					Email:    email,
				}, nil)
			},
			response: true,
		},
		{
			name:     "IncorrectUsername",
			username: "user",
			email:    "a@d.com",
			mockBehaviourGetByUsername: func(r *mockDomain.MockUsersMicroservice, username string) {
				r.EXPECT().GetByUsername(username).Return(models.User{}, errors.New("not exist"))
			},
			mockBehaviourGetByEmail: func(r *mockDomain.MockUsersMicroservice, email string) {},
			response:                false,
		},
		{
			name:     "IncorrectUsername",
			username: "user",
			email:    "a@d.com",
			mockBehaviourGetByUsername: func(r *mockDomain.MockUsersMicroservice, username string) {
				r.EXPECT().GetByUsername(username).Return(models.User{
					Username: username,
					Email:    "adnsonjo@d.com",
				}, nil)
			},
			mockBehaviourGetByEmail: func(r *mockDomain.MockUsersMicroservice, email string) {
				r.EXPECT().GetByEmail(email).Return(models.User{}, errors.New("not exist"))
			},
			response: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMock := mockDomain.NewMockUsersMicroservice(ctrl)
			imgMock := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehaviourGetByUsername(userMock, test.username)
			test.mockBehaviourGetByEmail(userMock, test.email)

			usecase := New(userMock, imgMock)

			isExist := usecase.IsExistUsernameAndEmail(test.username, test.email)

			require.Equal(t, isExist, test.response)
		})
	}
}

func TestUsecase_FindAuthors(t *testing.T) {
	type mockImages func(r *mockDomain.MockImageUseCase, avatar string)
	type mockAuthors func(r *mockDomain.MockUsersMicroservice, word string)
	type mockAllAuthors func(r *mockDomain.MockUsersMicroservice)

	tests := []struct {
		name           string
		word           string
		img            string
		mockImages     mockImages
		mockAuthors    mockAuthors
		mockAllAuthors mockAllAuthors
		response       []models.User
		expectedError  string
	}{
		{
			name: "OK",
			word: "user",
			img:  "",
			mockAuthors: func(r *mockDomain.MockUsersMicroservice, word string) {
				r.EXPECT().GetAuthorByUsername(word).Return([]models.User{
					{
						Username: "user",
					},
				}, nil)
			},
			mockImages: func(r *mockDomain.MockImageUseCase, avatar string) {
				r.EXPECT().GetImage(avatar).Return("image", nil)
			},
			mockAllAuthors: func(r *mockDomain.MockUsersMicroservice) {},
			response: []models.User{
				{
					Username: "user",
					Avatar:   "image",
				},
			},
		},
		{
			name: "ErrEmptyWord",
			word: "",
			mockAllAuthors: func(r *mockDomain.MockUsersMicroservice) {
				r.EXPECT().GetAllAuthors().Return([]models.User{}, errors.New("empty word"))
			},
			mockAuthors:   func(r *mockDomain.MockUsersMicroservice, word string) {},
			mockImages:    func(r *mockDomain.MockImageUseCase, avatar string) {},
			expectedError: "empty word",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			imgMock := mockDomain.NewMockImageUseCase(ctrl)
			userMicro := mockDomain.NewMockUsersMicroservice(ctrl)

			test.mockImages(imgMock, "")
			test.mockAuthors(userMicro, test.word)
			test.mockAllAuthors(userMicro)

			usecase := New(userMicro, imgMock)

			authors, err := usecase.FindAuthors(test.word)
			if err != nil {
				require.Equal(t, err.Error(), "empty word")
			}

			require.Equal(t, authors, test.response)
		})
	}
}

func TestUsersUsecase_GetPostsNum(t *testing.T) {
	type mockBehaviour func(r *mockDomain.MockUsersMicroservice, userID uint64)

	tests := []struct {
		name          string
		userID        uint64
		mockBehaviour mockBehaviour
		response      uint64
		expectedError string
	}{
		{
			name:   "OK",
			userID: 1,
			mockBehaviour: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetPostsNum(userID).Return(uint64(1), nil)
			},
			response: 1,
		},
		{
			name:   "Err",
			userID: 0,
			mockBehaviour: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetPostsNum(userID).Return(uint64(0), errors.New("empty username"))
			},
			expectedError: "empty username",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMicro := mockDomain.NewMockUsersMicroservice(ctrl)

			test.mockBehaviour(userMicro, test.userID)

			usecase := New(userMicro, nil)

			postsNum, err := usecase.GetPostsNum(test.userID)
			if err != nil {
				require.Equal(t, err.Error(), test.expectedError)
			}

			require.Equal(t, postsNum, test.response)
		})
	}
}

func TestUsersUsecase_GetSubscribersNum(t *testing.T) {
	type mockBehaviour func(r *mockDomain.MockUsersMicroservice, userID uint64)

	tests := []struct {
		name          string
		userID        uint64
		mockBehaviour mockBehaviour
		response      uint64
		expectedError string
	}{
		{
			name:   "OK",
			userID: 1,
			mockBehaviour: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetSubscribersNum(userID).Return(uint64(1), nil)
			},
			response: 1,
		},
		{
			name:   "Err",
			userID: 0,
			mockBehaviour: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetSubscribersNum(userID).Return(uint64(0), errors.New("empty username"))
			},
			expectedError: "empty username",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMicro := mockDomain.NewMockUsersMicroservice(ctrl)

			test.mockBehaviour(userMicro, test.userID)

			usecase := New(userMicro, nil)

			subscribersNum, err := usecase.GetSubscribersNum(test.userID)
			if err != nil {
				require.Equal(t, err.Error(), test.expectedError)
			}

			require.Equal(t, subscribersNum, test.response)
		})
	}
}

func TestUsersUsecase_GetProfitForMounth(t *testing.T) {
	type mockBehaviour func(r *mockDomain.MockUsersMicroservice, userID uint64)

	tests := []struct {
		name          string
		userID        uint64
		mockBehaviour mockBehaviour
		response      uint64
		expectedError string
	}{
		{
			name:   "OK",
			userID: 1,
			mockBehaviour: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetProfitForMounth(userID).Return(uint64(1), nil)
			},
			response: 1,
		},
		{
			name:   "Err",
			userID: 0,
			mockBehaviour: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetProfitForMounth(userID).Return(uint64(0), errors.New("empty username"))
			},
			expectedError: "empty username",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userMicro := mockDomain.NewMockUsersMicroservice(ctrl)

			test.mockBehaviour(userMicro, test.userID)

			usecase := New(userMicro, nil)

			profit, err := usecase.GetProfitForMounth(test.userID)
			if err != nil {
				require.Equal(t, err.Error(), test.expectedError)
			}

			require.Equal(t, profit, test.response)
		})
	}
}
