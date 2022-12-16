package images

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
)

func TestUsecase_GetImage(t *testing.T) {
	type mockBehaviourGetImage func(r *mockDomain.MockImageMicroservice, filename string)

	tests := []struct {
		name                  string
		filename              string
		mockBehaviourGetImage mockBehaviourGetImage
		responseError         string
		expectedResult        string
	}{
		{
			name:     "OK",
			filename: "filename.jpg",
			mockBehaviourGetImage: func(r *mockDomain.MockImageMicroservice, filename string) {
				r.EXPECT().Get(filename).Return("newURL", nil)
			},
			expectedResult: "https://wsrv.nl/?url=newURL",
		},
		{
			name:     "Bad-image",
			filename: "filename.jpg",
			mockBehaviourGetImage: func(r *mockDomain.MockImageMicroservice, filename string) {
				r.EXPECT().Get(filename).Return("", domain.ErrInternal)
			},
			responseError: "server error",
		},
		{
			name:                  "BadRequest-Filename",
			filename:              "",
			mockBehaviourGetImage: func(r *mockDomain.MockImageMicroservice, filename string) {},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			image := mockDomain.NewMockImageMicroservice(ctrl)

			test.mockBehaviourGetImage(image, test.filename)

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
