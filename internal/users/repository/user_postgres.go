package userRepository

import (
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgres(URL string) (*Postgres, error) {
	db, err := sqlx.Open("postgres", URL)
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

func (r *Postgres) Create(u *model.User) (*model.User, error) {
	err := r.DB.QueryRowx(
		`
		INSERT INTO users (username, first_name, last_name, avatar, email, password, is_author, about) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`,
		u.Username,
		u.FirstName,
		u.LastName,
		u.Avatar,
		u.Email,
		u.Password,
		u.IsAuthor,
		u.About,
	).Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *Postgres) GetByUsername(username string) (*model.User, error) {
	var u model.User
	if err := r.DB.Get(
		&u,
		"SELECT id, username, first_name, last_name, avatar, email, password, is_author, about FROM users WHERE username = $1;",
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
		"SELECT id, username, first_name, last_name, avatar, email, password, is_author, about FROM users WHERE id = $1;",
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
		"SELECT id, username, first_name, last_name, avatar, email, password, is_author, about FROM users WHERE email = $1;",
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
		SELECT id, username, first_name, last_name, avatar, email, password, is_author, about 
		FROM users JOIN sessions ON sessions.user_id = id
    	WHERE sessions.value = $1;`,
		sessionID,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Postgres) Update(user *model.User) (*model.User, error) {
	_, err := r.DB.NamedExec(
		`
		UPDATE users 
		SET username=:username,
		    first_name=:first_name,
		    last_name=:last_name,
		    avatar=:avatar,
		    email=:email,
		    password=:password,
		    phone=:phone,
		    is_author=:is_author,
		    about=:about 
		WHERE id = :id`, &user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *Postgres) DeleteByID(id uint64) error {
	_, err := r.DB.Query("DELETE FROM users WHERE id=$1;", id)
	if err != nil {
		return err
	}
	return nil
}