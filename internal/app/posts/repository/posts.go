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

func (r *Repo) GetPostsByID(id uint) ([]models.PostDB, error) {
	rows, err := r.Storage.DB.Query("SELECT post_id, user_id, title FROM posts WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostDB

	for rows.Next() {
		var post models.PostDB
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return posts, err
	}
	return posts, nil
}
