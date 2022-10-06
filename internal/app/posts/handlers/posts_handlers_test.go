package postsHandlers

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/gorilla/mux"
)

type mockPostRepo struct {
	postDB []model.PostDB
}

func (p *mockPostRepo) GetPostsByID(id uint) ([]models.PostDB, error) {
	return p.postDB, nil
}

func TestGetPosts(t *testing.T) {
	test := struct {
		posts     []model.PostDB
		postsJson string
	}{
		posts: []model.PostDB{
			{
				ID:     1,
				UserID: 1,
				Title:  "test",
			},
			{
				ID:     12,
				UserID: 1,
				Title:  "TEST",
			},
		},
		postsJson: "[{\"id\":1,\"user_id\":1,\"title\":\"test\"},{\"id\":12,\"user_id\":1,\"title\":\"TEST\"}]",
	}

	postRepo := mockPostRepo{
		postDB: test.posts,
	}

	handler := NewHTPPHandler(&postRepo)
	req := httptest.NewRequest("GET", "/api/v1/users/1/post", nil)
	res := httptest.NewRecorder()
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	handler.Posts(res, req)
	bytes, _ := ioutil.ReadAll(res.Body)
	if strings.Trim(string(bytes), "\n") != test.postsJson {
		t.Errorf("expected: [%s], got: [%s]", test.postsJson, string(bytes))
	}
}
