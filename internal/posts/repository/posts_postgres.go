package postsRepository

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgres(url string) (*Postgres, error) {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}

func (r *Postgres) Close() error {
	if err := r.DB.Close(); err != nil {
		return err
	}

	return nil
}

func (r *Postgres) GetAllByUserID(userID uint64) ([]*models.Post, error) {
	var posts []*models.Post
	if err := r.DB.Select(&posts, "SELECT * FROM posts WHERE user_id=$1;", userID); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *Postgres) GetPostByID(postID uint64) (*models.Post, error) {
	var post models.Post
	if err := r.DB.Get(&post, "SELECT * FROM posts WHERE post_id=$1;", postID); err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *Postgres) Create(post models.Post) (*models.Post, error) {
	err := r.DB.QueryRowx(
		`
		INSERT INTO posts (user_id, img, title, text) 
		VALUES ($1, $2, $3, $4) RETURNING post_id;`,
		post.UserID,
		post.Img,
		post.Title,
		post.Text,
	).Scan(&post.ID)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *Postgres) Update(post models.Post) (*models.Post, error) {
	_, err := r.DB.NamedExec(
		`
                UPDATE posts
                SET user_id=:user_id,
                    title=:title,
                    text=:text,
                    img=:img
                WHERE post_id = :post_id`, &post)
	if err != nil {
		return nil, err
	}

	return &post, err
}

func (r *Postgres) DeleteByID(postID uint64) error {
	_, err := r.DB.Exec("DELETE FROM posts WHERE post_id=$1;", postID)
	if err != nil {
		return err
	}

	return nil
}
