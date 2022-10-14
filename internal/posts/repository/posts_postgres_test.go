package postsRepository

import (
	"log"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zhashkevych/go-sqlxmock"
)

func TestPostPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := &Postgres{DB: db}
	require.NoError(t, err)
	db = r.DB

	type Args struct {
		userId uint64
		post   models.PostDB
	}

	type MockBehaviour func(args Args, id uint64)

	tests := []struct {
		name          string
		mockBehaviour MockBehaviour
		args          Args
		id            uint64
		responseError bool
	}{
		{
			name: "Ok",
			args: Args{
				userId: 0,
				post: models.PostDB{
					Title: "Hey",
					Text: "text",
				},
			},
			id: 2,
			mockBehaviour: func(args Args, id uint64) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO posts").
					WithArgs(args.userId, args.post.Img, args.post.Title, args.post.Text).
						WillReturnRows(rows)
				mock.ExpectCommit()
			},
		}, 
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args, test.id)

			got, err := r.Create(&test.args.post)
			if test.responseError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.id, got.ID)
			}
		})
	}
}
