package imagesMicroservice

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/images/protobuf"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestImagesClient_Create(t *testing.T) {
	type mockCreate func(r *mockDomain.MockImagesClient, c context.Context, in *protobuf.Image, opts ...grpc.CallOption)

	tests := []struct {
		name     string
		filename string
		file     []byte
		size     int64
		oldName  string
		mock     mockCreate
		want     string
		wantErr  error
	}{
		{
			name:     "OK",
			filename: "test",
			file:     []byte("test"),
			size:     4,
			oldName:  "test",
			mock: func(r *mockDomain.MockImagesClient, c context.Context, in *protobuf.Image, opts ...grpc.CallOption) {
				r.EXPECT().Create(c, in).Return(&protobuf.Filename{Filename: "test"}, nil)
			},
			want:    "test",
			wantErr: nil,
		},
		{
			name:     "Error",
			filename: "test",
			file:     []byte("test"),
			size:     4,
			oldName:  "test",
			mock: func(r *mockDomain.MockImagesClient, c context.Context, in *protobuf.Image, opts ...grpc.CallOption) {
				r.EXPECT().Create(c, in).Return(&protobuf.Filename{Filename: "test"}, status.Error(codes.Canceled, "canceled"))
			},
			want:    "",
			wantErr: status.Error(codes.Canceled, "canceled"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mockDomain.NewMockImagesClient(ctrl)
			test.mock(mock, context.Background(), &protobuf.Image{
				Filename:    test.filename,
				Content:     test.file,
				Size:        test.size,
				OldFilename: test.oldName,
			})

			m := ImagesMicroservice{
				client: mock,
			}

			got, err := m.Create(test.filename, test.file, test.size, test.oldName)

			require.Equal(t, test.want, got)
			require.Equal(t, test.wantErr, err)
		})
	}
}

func TestImagesClient_Get(t *testing.T) {
	type mockGet func(r *mockDomain.MockImagesClient, c context.Context, in *protobuf.Filename, opts ...grpc.CallOption)

	tests := []struct {
		name     string
		filename string
		mock     mockGet
		want     string
		wantErr  error
	}{
		{
			name:     "OK",
			filename: "test",
			mock: func(r *mockDomain.MockImagesClient, c context.Context, in *protobuf.Filename, opts ...grpc.CallOption) {
				r.EXPECT().Get(c, in).Return(&protobuf.URL{Url: "test"}, nil)
			},
			want: "test",
		},
		{
			name:     "Error",
			filename: "test",
			mock: func(r *mockDomain.MockImagesClient, c context.Context, in *protobuf.Filename, opts ...grpc.CallOption) {
				r.EXPECT().Get(c, in).Return(&protobuf.URL{Url: "test"}, status.Error(codes.Canceled, "canceled"))
			},
			want:    "",
			wantErr: status.Error(codes.Canceled, "canceled"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mockDomain.NewMockImagesClient(ctrl)
			test.mock(mock, context.Background(), &protobuf.Filename{Filename: test.filename})

			m := ImagesMicroservice{
				client: mock,
			}

			got, err := m.Get(test.filename)

			require.Equal(t, test.want, got)
			require.Equal(t, test.wantErr, err)
		})
	}
}
