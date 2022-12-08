package grpcImages

import (
	"context"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/images/protobuf"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestImagesService_Create(t *testing.T) {
	type mockBehaviorCreateOrUpdateImage func(r *mockDomain.MockImagesRepository, filename string, content []byte, size int64, oldFilename string)

	tests := []struct {
		name                            string
		inputFilename                   string
		inputContent                    []byte
		inputSize                       int64
		inputOldFilename                string
		mockBehaviorCreateOrUpdateImage mockBehaviorCreateOrUpdateImage
		expected                        *protobuf.Filename
		expectedErr                     string
	}{
		{
			name:             "OK",
			inputFilename:    "filename",
			inputContent:     []byte("content"),
			inputSize:        1,
			inputOldFilename: "oldFilename",
			mockBehaviorCreateOrUpdateImage: func(r *mockDomain.MockImagesRepository, filename string, content []byte, size int64, oldFilename string) {
				r.EXPECT().CreateOrUpdateImage(filename, content, size, oldFilename).Return("filename", nil)
			},
			expected: &protobuf.Filename{
				Filename: "filename",
			},
		},
		{
			name:             "Error",
			inputFilename:    "filename",
			inputContent:     []byte("content"),
			inputSize:        1,
			inputOldFilename: "oldFilename",
			mockBehaviorCreateOrUpdateImage: func(r *mockDomain.MockImagesRepository, filename string, content []byte, size int64, oldFilename string) {
				r.EXPECT().CreateOrUpdateImage(filename, content, size, oldFilename).Return("", domain.ErrInternal)
			},
			expectedErr: "server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockDomain.NewMockImagesRepository(ctrl)
			test.mockBehaviorCreateOrUpdateImage(repo, test.inputFilename, test.inputContent, test.inputSize, test.inputOldFilename)

			s := New(repo)
			result, err := s.Create(context.Background(), &protobuf.Image{
				Filename:    test.inputFilename,
				Content:     test.inputContent,
				Size:        test.inputSize,
				OldFilename: test.inputOldFilename,
			})

			if err != nil {
				require.Equal(t, test.expectedErr, err.Error())
			} else {
				require.Equal(t, test.expected, result)
			}
		})
	}
}

func TestImagesService_Get(t *testing.T) {
	type mockBehaviorGetPermanentImage func(r *mockDomain.MockImagesRepository, filename string)

	tests := []struct {
		name                          string
		inputFilename                 string
		mockBehaviorGetPermanentImage mockBehaviorGetPermanentImage
		expected                      *protobuf.URL
		expectedErr                   string
	}{
		{
			name:          "OK",
			inputFilename: "filename",
			mockBehaviorGetPermanentImage: func(r *mockDomain.MockImagesRepository, filename string) {
				r.EXPECT().GetPermanentImage(filename).Return("url", nil)
			},
			expected: &protobuf.URL{
				Url: "url",
			},
		},
		{
			name:          "Error",
			inputFilename: "filename",
			mockBehaviorGetPermanentImage: func(r *mockDomain.MockImagesRepository, filename string) {
				r.EXPECT().GetPermanentImage(filename).Return("", domain.ErrInternal)
			},
			expectedErr: "server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockDomain.NewMockImagesRepository(ctrl)
			test.mockBehaviorGetPermanentImage(repo, test.inputFilename)

			s := New(repo)
			result, err := s.Get(context.Background(), &protobuf.Filename{
				Filename: test.inputFilename,
			})

			if err != nil {
				require.Equal(t, test.expectedErr, err.Error())
			} else {
				require.Equal(t, test.expected, result)
			}
		})
	}
}
