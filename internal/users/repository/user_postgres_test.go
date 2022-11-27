package userRepository

import (
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestPostPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	assert.NoError(t, err)

	r := &Postgres{DB: db}

	require.NoError(t, err)

	type MockBehaviour func(user models.User, id uint64)

	tests := []struct {
		name          string
		mockBehaviour MockBehaviour
		user          models.User
		id            uint64
		responseError bool
	}{
		{
			name: "Ok",
			user: models.User{
				Username: "user",
				Email:    "user@ex.com",
				Password: "Qwerty",
				IsAuthor: false,
			},
			id: 2,
			mockBehaviour: func(user models.User, id uint64) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO users").
					WithArgs(user.Username, user.Avatar, user.Email, user.Password, user.IsAuthor, user.About).
					WillReturnRows(rows)
				mock.ExpectCommit()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.user, test.id)

			_, err := r.Create(test.user)
			if test.responseError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
