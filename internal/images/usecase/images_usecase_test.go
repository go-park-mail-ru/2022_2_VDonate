package images

import (
	"net/url"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_GetImage(t *testing.T) {
	type mockBehaviourGetImage func(r *mockDomain.MockImagesRepository, bucket, filename string)
	type mockBehaviourGetPermImage func(r *mockDomain.MockImagesRepository, bucket, filename string)

	tests := []struct {
		name                      string
		bucket                    string
		filename                  string
		mockBehaviourGetImage     mockBehaviourGetImage
		mockBehaviourGetPermImage mockBehaviourGetPermImage
		responseError             string
		expectedResult            string
	}{
		{
			name:     "OK-image",
			filename: "filename.jpg",
			bucket:   "image",
			mockBehaviourGetImage: func(r *mockDomain.MockImagesRepository, bucket, filename string) {
				r.EXPECT().GetImage(bucket, filename).Return(&url.URL{
					Path: "newURL",
				}, nil)
			},
			mockBehaviourGetPermImage: func(r *mockDomain.MockImagesRepository, bucket, filename string) {},
			expectedResult:            "https://wsrv.nl/?url=newURL",
		},
		{
			name:     "Bad-image",
			filename: "filename.jpg",
			bucket:   "image",
			mockBehaviourGetImage: func(r *mockDomain.MockImagesRepository, bucket, filename string) {
				r.EXPECT().GetImage(bucket, filename).Return(nil, domain.ErrNotFound)
			},
			mockBehaviourGetPermImage: func(r *mockDomain.MockImagesRepository, bucket, filename string) {},
			responseError:             "failed to find item",
		},
		{
			name:                  "OK-avatar",
			filename:              "filename.jpg",
			bucket:                "avatar",
			mockBehaviourGetImage: func(r *mockDomain.MockImagesRepository, bucket, filename string) {},
			mockBehaviourGetPermImage: func(r *mockDomain.MockImagesRepository, bucket, filename string) {
				r.EXPECT().GetPermanentImage(bucket, filename).Return("newURL", nil)
			},
			expectedResult: "https://wsrv.nl/?url=newURL",
		},
		{
			name:                  "Bad-avatar",
			filename:              "filename.jpg",
			bucket:                "avatar",
			mockBehaviourGetImage: func(r *mockDomain.MockImagesRepository, bucket, filename string) {},
			mockBehaviourGetPermImage: func(r *mockDomain.MockImagesRepository, bucket, filename string) {
				r.EXPECT().GetPermanentImage(bucket, filename).Return("", domain.ErrNotFound)
			},
			responseError: "failed to find item",
		},
		{
			name:                      "BadRequest-Bucket",
			filename:                  "filename.jpg",
			bucket:                    "",
			mockBehaviourGetImage:     func(r *mockDomain.MockImagesRepository, bucket, filename string) {},
			mockBehaviourGetPermImage: func(r *mockDomain.MockImagesRepository, bucket, filename string) {},
			responseError:             "bad url",
		},
		{
			name:                      "BadRequest-Filename",
			filename:                  "",
			bucket:                    "",
			mockBehaviourGetImage:     func(r *mockDomain.MockImagesRepository, bucket, filename string) {},
			mockBehaviourGetPermImage: func(r *mockDomain.MockImagesRepository, bucket, filename string) {},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			image := mockDomain.NewMockImagesRepository(ctrl)

			test.mockBehaviourGetImage(image, test.bucket, test.filename)
			test.mockBehaviourGetPermImage(image, test.bucket, test.filename)

			usecase := New(image)
			result, err := usecase.GetImage(test.bucket, test.filename)
			if err != nil {
				assert.Equal(t, err.Error(), test.responseError)
			} else {
				assert.Equal(t, result, test.expectedResult)
			}
		})
	}
}
