package userRepository

// import (
// 	"testing"

// 	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	sqlmock "github.com/zhashkevych/go-sqlxmock"
// )

// type getID struct {
// 	id     int64
// 	efRows int64
// }

// func (g getID) LastInsertId() (int64, error) {
// 	return g.id, nil
// }

// func (g getID) RowsAffected() (int64, error) {
// 	return g.efRows, nil
// }

// func TestPostPostgres_Create(t *testing.T) {
// 	db, mock, err := sqlmock.Newx()
// 	defer db.Close()
// 	assert.NoError(t, err)

// 	r := &Postgres{DB: db}
// 	require.NoError(t, err)

// 	type MockBehaviour func(user models.User, id uint64)

// 	tests := []struct {
// 		name          string
// 		mockBehaviour MockBehaviour
// 		user          models.User
// 		id            uint64
// 		responseError bool
// 	}{
// 		// FIXME this case need to fix
// 		{
// 			name: "Ok",
// 			user: models.User{
// 				Username: "user",
// 				Email:    "user@ex.com",
// 				Password: "Qwerty",
// 				IsAuthor: false,
// 			},
// 			id: 2,
// 			mockBehaviour: func(user models.User, id uint64) {
// 				mock.ExpectBegin()

// 				mock.ExpectExec("INSERT INTO users").
// 					WithArgs(user.Username, user.Email).
// 					WillReturnResult(
// 						getID{
// 							id:     2,
// 							efRows: 1,
// 						},
// 					)
// 				mock.ExpectExec("INSERT INTO user_info").
// 					WithArgs(
// 						2,
// 						user.Avatar,
// 						user.Password,
// 						user.IsAuthor,
// 						user.About,
// 					).
// 					WillReturnResult(
// 						getID{
// 							id:     0,
// 							efRows: 1,
// 						},
// 					)
// 				mock.ExpectCommit()
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.mockBehaviour(test.user, test.id)

// 			_, err := r.Create(test.user)
// 			if test.responseError {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }
