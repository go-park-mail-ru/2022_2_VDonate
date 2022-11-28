package postsRepository

import (
	"time"

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

func (r Postgres) Close() error {
	if err := r.DB.Close(); err != nil {
		return err
	}

	return nil
}

func (r Postgres) GetAllByUserID(authorID uint64) ([]models.Post, error) {
	var posts []models.Post
	if err := r.DB.Select(
		&posts,
		"SELECT * FROM posts WHERE user_id=$1;",
		authorID,
	); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r Postgres) GetPostByID(postID uint64) (models.Post, error) {
	var post models.Post
	if err := r.DB.Get(
		&post,
		"SELECT * FROM posts WHERE post_id=$1;",
		postID,
	); err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (r Postgres) Create(post models.Post) (uint64, error) {
	var postID uint64
	err := r.DB.QueryRowx(
		`
		INSERT INTO posts (user_id, img, title, text, tier, date_created) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING post_id;`,
		post.UserID,
		post.Img,
		post.Title,
		post.Text,
		post.Tier,
		time.Now(),
	).Scan(&postID)
	if err != nil {
		return 0, err
	}

	return postID, nil
}

func (r Postgres) Update(post models.Post) error {
	_, err := r.DB.NamedExec(
		`
                UPDATE posts
                SET user_id=:user_id,
                    title=:title,
                    text=:text,
                    img=:img,
                    tier=:tier
                WHERE post_id = :post_id`, &post)

	return err
}

func (r Postgres) GetPostsBySubscriptions(userID uint64) ([]models.Post, error) {
	var posts []models.Post
	if err := r.DB.Select(&posts, `
		SELECT p.post_id, p.user_id, p.img, p.title, p.text, p.tier, p.date_created
		FROM subscriptions s
		JOIN posts p on s.author_id = p.user_id
		JOIN author_subscriptions "as" on "as".id = s.subscription_id
		WHERE s.subscriber_id=$1 AND "as".tier >= p.tier;
	`, userID); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r Postgres) DeleteByID(postID uint64) error {
	_, err := r.DB.Exec("DELETE FROM posts WHERE post_id=$1;", postID)
	if err != nil {
		return err
	}

	return nil
}

func (r Postgres) GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error) {
	var like models.Like
	if err := r.DB.Get(&like, "SELECT * FROM likes WHERE user_id=$1 AND post_id=$2;", userID, postID); err != nil {
		return models.Like{}, err
	}
	return like, nil
}

func (r Postgres) GetAllLikesByPostID(postID uint64) ([]models.Like, error) {
	var likes []models.Like
	if err := r.DB.Select(&likes, "SELECT * FROM likes WHERE post_id=$1;", postID); err != nil {
		return nil, err
	}
	return likes, nil
}

func (r Postgres) CreateLike(userID, postID uint64) error {
	return r.DB.QueryRowx(
		`
		INSERT INTO likes (user_id, post_id)
		VALUES ($1, $2);`,
		userID,
		postID,
	).Err()
}

func (r Postgres) DeleteLikeByID(userID, postID uint64) error {
	_, err := r.DB.Exec("DELETE FROM likes WHERE user_id=$1 AND post_id=$2;", userID, postID)
	return err
}

func (r Postgres) CreateDepTag(postID, tagID uint64) error {
	return r.DB.QueryRowx(
		`
		INSERT INTO post_tags (post_id, tag_id)
		VALUES ($1, $2);`,
		postID,
		tagID,
	).Err()
}

func (r Postgres) DeleteDepTag(tagDep models.TagDep) error {
	_, err := r.DB.Exec("DELETE FROM post_tags WHERE post_id=$1 AND tag_id=$2;", tagDep.PostID, tagDep.TagID)
	return err
}

func (r Postgres) CreateTag(tagName string) (uint64, error) {
	var tagID uint64
	err := r.DB.QueryRowx(
		`
		INSERT INTO tags (tag_name)
		VALUES ($1) RETURNING id;`,
		tagName,
	).Scan(&tagID)
	if err != nil {
		return 0, err
	}

	return tagID, nil
}

func (r Postgres) GetTagById(tagID uint64) (models.Tag, error) {
	var tag models.Tag
	if err := r.DB.Get(&tag, "SELECT * FROM tags WHERE id=$1;", tagID); err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}

func (r Postgres) GetTagDepsByPostId(postID uint64) ([]models.TagDep, error) {
	var tagDeps []models.TagDep
	if err := r.DB.Select(&tagDeps, "SELECT * FROM post_tags WHERE post_id=$1;", postID); err != nil {
		return nil, err
	}
	return tagDeps, nil
}

func (r Postgres) GetTagByName(tagName string) (models.Tag, error) {
	var tag models.Tag
	if err := r.DB.Get(&tag, "SELECT * FROM tags WHERE tag_name=$1;", tagName); err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}
