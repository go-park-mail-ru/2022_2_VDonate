package userRepository

import (
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
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

func (r *Postgres) Create(user *model.User) error {
	return r.DB.QueryRowx(
		`
		INSERT INTO users (username, avatar, email, password, is_author, about) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id;`,
		user.Username,
		user.Avatar,
		user.Email,
		user.Password,
		user.IsAuthor,
		user.About,
	).Err()
}

func (r *Postgres) GetByUsername(username string) (*model.User, error) {
	var u model.User
	if err := r.DB.Get(
		&u,
		`
		SELECT id, username, avatar, email, password, is_author, about 
		FROM users 
		WHERE username = $1;`,
		username,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Postgres) GetByID(id uint64) (*model.User, error) {
	var u model.User
	if err := r.DB.Get(
		&u,
		`
		SELECT id, username, avatar, email, password, is_author, about 
		FROM users 
		WHERE id = $1;`,
		id,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Postgres) GetByEmail(email string) (*model.User, error) {
	var u model.User
	if err := r.DB.Get(
		&u,
		`
		SELECT id, username, avatar, email, password, is_author, about 
		FROM users 
		WHERE email = $1;`,
		email,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Postgres) GetBySessionID(sessionID string) (*model.User, error) {
	var u model.User
	if err := r.DB.Get(
		&u,
		`
		SELECT id, username, avatar, email, password, is_author, about 
		FROM users JOIN sessions ON sessions.user_id = id
    	WHERE sessions.value = $1;`,
		sessionID,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Postgres) GetUserByPostID(postID uint64) (*model.User, error) {
	var user model.User
	if err := r.DB.Get(&user, `
		SELECT id, username, avatar, email, password, is_author, about 
		FROM posts 
		JOIN users on users.id = posts.user_id 
		WHERE posts.post_id = $1`, postID,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Postgres) Update(user *model.User) error {
	_, err := r.DB.NamedExec(
		`
		UPDATE users 
		SET username=:username,
		    avatar=:avatar,
		    email=:email,
		    password=:password,
		    is_author=:is_author,
		    about=:about 
		WHERE id = :id`, user)

	return err
}

func (r *Postgres) DeleteByID(id uint64) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id=$1;", id)
	if err != nil {
		return err
	}

	return nil
}
