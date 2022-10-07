package userPostgres

import (
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sqlx.DB
}

func New(URL string) (userRepository.API, error) {
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

func (r *Repository) Create(u *model.UserDB) (*model.UserDB, error) {
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

func (r *Repository) FindByUsername(username string) (*model.UserDB, error) {
	var u model.UserDB
	if err := r.DB.Get(
		&u,
		"SELECT id, username, first_name, last_name, avatar, email, password, is_author, about FROM users WHERE username = $1;",
		username,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) FindByID(id uint) (*model.UserDB, error) {
	var u model.UserDB
	if err := r.DB.Get(
		&u,
		"SELECT id, username, first_name, last_name, avatar, email, password, is_author, about FROM users WHERE id = $1;",
		id,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) FindByEmail(email string) (*model.UserDB, error) {
	var u model.UserDB
	if err := r.DB.Get(
		&u,
		"SELECT id, username, first_name, last_name, avatar, email, password, is_author, about FROM users WHERE email = $1;",
		email,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) Update(user *model.UserDB) (*model.UserDB, error) {
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

func (r *Repository) DeleteByID(id uint) error {
	_, err := r.DB.Query("DELETE FROM users WHERE id=$1;", id)
	if err != nil {
		return err
	}
	return nil
}
