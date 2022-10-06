package userHandlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/gorilla/mux"
)

type mockUserRepo struct {
	userDB model.UserDB
}

func (r *mockUserRepo) Create(u *model.UserDB) error {
	return nil
}

func (r *mockUserRepo) FindByUsername(username string) (*model.UserDB, error) {
	return &r.userDB, nil
}

func (r *mockUserRepo) FindByID(id uint) (*model.UserDB, error) {
	return &r.userDB, nil
}

func (r *mockUserRepo) FindByEmail(email string) (*model.UserDB, error) {
	return &r.userDB, nil
}

type mockSessionRepo struct {
	cookie http.Cookie
}

func (s *mockSessionRepo) Create(id uint) *http.Cookie {
	return &s.cookie
}

func (s *mockSessionRepo) Remove(CID string) {}

func (s *mockSessionRepo) GetId(CID string) (uint, bool) {
	return 1, true
}

func TestGetUser(t *testing.T) {
	test := struct {
		user     model.UserDB
		userJson string
	}{
		user: model.UserDB{
			ID:       1,
			Username: "test",
			Password: "testing",
			Email:    "test@test.test",
			IsAuthor: false,
		},
		userJson: "{\"id\":1,\"username\":\"test\",\"email\":\"test@test.test\",\"password\":\"testing\",\"is_author\":false}",
	}

	userRepo := mockUserRepo{
		userDB: test.user,
	}
	var sessionRepo mockSessionRepo

	handler := NewHTTPHandler(&userRepo, &sessionRepo)
	req := httptest.NewRequest("GET", "/api/v1/users/1", nil)
	res := httptest.NewRecorder()
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	handler.GetUser(res, req)
	bytes, _ := ioutil.ReadAll(res.Body)
	if strings.Trim(string(bytes), "\n") != test.userJson {
		t.Errorf("expected: [%s], got: [%s]", test.userJson, string(bytes))
	}
}
