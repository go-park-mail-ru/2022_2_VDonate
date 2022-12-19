package postsRepository

import (
	"database/sql"
	"errors"

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
	posts := make([]models.Post, 0)
	if err := r.DB.Select(
		&posts,
		"SELECT * FROM posts WHERE user_id=$1;",
		authorID,
	); err != nil && !errors.Is(err, sql.ErrNoRows) {
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

func (r Postgres) Create(post models.Post) (models.Post, error) {
	err := r.DB.QueryRowx(
		`
		INSERT INTO posts (user_id, content, tier) 
		VALUES ($1, $2, $3) RETURNING post_id, date_created;`,
		post.UserID,
		post.Content,
		post.Tier,
	).Scan(&post.ID, &post.DateCreated)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (r Postgres) Update(post models.Post) error {
	_, err := r.DB.NamedExec(
		`
                UPDATE posts
                SET user_id=:user_id,
                    content=:content,
                    tier=:tier
                WHERE post_id = :post_id`, &post)

	return err
}

func (r Postgres) GetPostsBySubscriptions(userID uint64) ([]models.Post, error) {
	posts := make([]models.Post, 0)
	if err := r.DB.Select(&posts, `
		SELECT p.post_id, p.user_id, p.content, p.tier, p.date_created
		FROM subscriptions s
		JOIN posts p on s.author_id = p.user_id
		JOIN author_subscriptions "as" on "as".id = s.subscription_id
		WHERE s.subscriber_id=$1 AND "as".tier >= p.tier;
	`, userID); err != nil && !errors.Is(err, sql.ErrNoRows) {
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
	likes := make([]models.Like, 0)
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
	tagDeps := make([]models.TagDep, 0)
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

func (r Postgres) CreateComment(comment models.Comment) (models.Comment, error) {
	err := r.DB.QueryRowx(
		`
		INSERT INTO comments (user_id, post_id, content)
		VALUES ($1, $2, $3) RETURNING id, date_created;`,
		comment.UserID,
		comment.PostID,
		comment.Content,
	).Scan(&comment.ID, &comment.DateCreated)
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (r Postgres) GetCommentsByPostId(postID uint64) ([]models.Comment, error) {
	comments := make([]models.Comment, 0)
	if err := r.DB.Select(&comments, `SELECT * FROM comments WHERE post_id=$1;`, postID); err != nil {
		return nil, err
	}
	return comments, nil
}

func (r Postgres) GetCommentByID(commentID uint64) (models.Comment, error) {
	var comment models.Comment
	if err := r.DB.Get(&comment, "SELECT * FROM comments WHERE id=$1;", commentID); err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (r Postgres) UpdateComment(comment models.Comment) error {
	_, err := r.DB.NamedExec(
		`
		UPDATE comments
		SET content=:content
		WHERE id = :id`, &comment)
	return err
}

func (r Postgres) DeleteCommentByID(commentID uint64) error {
	_, err := r.DB.Exec("DELETE FROM comments WHERE id=$1;", commentID)
	return err
}
