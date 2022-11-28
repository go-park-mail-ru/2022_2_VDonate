package images

import (
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_GetImage(t *testing.T) {
	type mockBehaviourGetPermImage func(r *mockDomain.MockImagesRepository, filename string)

	tests := []struct {
		name                      string
		filename                  string
		mockBehaviourGetPermImage mockBehaviourGetPermImage
		responseError             string
		expectedResult            string
	}{
		{
			name:     "OK",
			filename: "filename.jpg",
			mockBehaviourGetPermImage: func(r *mockDomain.MockImagesRepository, filename string) {
				r.EXPECT().GetPermanentImage(filename).Return("newURL", nil)
			},
			expectedResult: "https://wsrv.nl/?url=newURL",
		},
		{
			name:     "Bad-image",
			filename: "filename.jpg",
			mockBehaviourGetPermImage: func(r *mockDomain.MockImagesRepository, filename string) {
				r.EXPECT().GetPermanentImage(filename).Return("", domain.ErrInternal)
			},
			responseError: "server error",
		},
		{
			name:                      "BadRequest-Filename",
			filename:                  "",
			mockBehaviourGetPermImage: func(r *mockDomain.MockImagesRepository, filename string) {},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			image := mockDomain.NewMockImagesRepository(ctrl)

			test.mockBehaviourGetPermImage(image, test.filename)

			usecase := New(image)
			result, err := usecase.GetImage(test.filename)
			if err != nil {
				assert.Equal(t, err.Error(), test.responseError)
			} else {
				assert.Equal(t, result, test.expectedResult)
			}
		})
	}
}
