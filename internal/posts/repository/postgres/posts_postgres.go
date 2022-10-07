package postsPostgres

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sqlx.DB
}

func New(URL string) (postsRepository.API, error) {
	db, err := sqlx.Open("postgres", URL)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Repository{DB: db}, nil
}

func (r *Repository) Close() error {
	if err := r.DB.Close(); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllByUserID(userID uint) ([]*models.PostDB, error) {
	var posts []*models.PostDB
	if err := r.DB.Select(&posts, "SELECT * FROM posts WHERE user_id=$1;", userID); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *Repository) GetPostByUserID(userID, postID uint) (*models.PostDB, error) {
	var post models.PostDB
	if err := r.DB.Get(&post, "SELECT * FROM posts WHERE user_id=$1 AND post_id=$2;", userID, postID); err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *Repository) CreateInUserByID(post models.PostDB) error {
	return r.DB.QueryRowx(
		`
		INSERT INTO posts (user_id, title, text) 
		VALUES ($1, $2, $3) RETURNING *;`,
		post.UserID,
		post.Title,
		post.Text,
	).Scan(&post)
}

func (r *Repository) DeleteInUserByID(userID, postID uint) error {
	_, err := r.DB.Query("DELETE FROM posts WHERE user_id=? AND post_id=$1;", userID, postID)
	if err != nil {
		return err
	}
	return nil
}
