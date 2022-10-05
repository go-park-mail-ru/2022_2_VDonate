package postsRepository

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	storage "github.com/go-park-mail-ru/2022_2_VDonate/internal/storages"
)

type Repo struct {
	Storage *storage.Storage
}

func New(s *storage.Storage) *Repo {
	return &Repo{Storage: s}
}

func (r *Repo) GetPostsByID(id uint) (*models.PostDB, error) {
	p := &models.PostDB{}
	if err := r.Storage.DB.QueryRow(
		"SELECT post_id, user_id, title FROM posts WHERE user_id = $1",
		id,
	).Scan(
		&p.ID,
		&p.UserID,
		&p.Title,
	); err != nil {
		return nil, err
	}

	return p, nil
}
