package userRepository

import (
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	storage "github.com/go-park-mail-ru/2022_2_VDonate/internal/storages"
)

type RepoI interface {
	Create(u *model.UserDB) error
	FindByUsername(username string) (*model.UserDB, error)
	FindByID(id uint) (*model.UserDB, error)
	FindByEmail(email string) (*model.UserDB, error)
}

type Repo struct {
	Storage *storage.Storage
}

func New(s *storage.Storage) *Repo {
	return &Repo{Storage: s}
}

// INSERT INTO posts (1, "SOME TEXT")

func (r *Repo) Create(u *model.UserDB) error {
	if err := r.Storage.DB.QueryRow(
		"INSERT INTO users (username, first_name, last_name, avatar, email, password, is_author, about) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		u.Username,
		u.FirstName,
		u.LastName,
		u.Avatar,
		u.Email,
		u.Password,
		u.IsAuthor,
		u.About,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func (r *Repo) FindByUsername(username string) (*model.UserDB, error) {
	u := &model.UserDB{}
	if err := r.Storage.DB.QueryRow(
		"SELECT id, username, first_name, last_name, avatar, email, password, is_author, about FROM users WHERE username = $1",
		username,
	).Scan(
		&u.ID,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.Avatar,
		&u.Email,
		&u.Password,
		&u.IsAuthor,
		&u.About,
	); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *Repo) FindByID(id uint) (*model.UserDB, error) {
	u := &model.UserDB{}
	if err := r.Storage.DB.QueryRow(
		"SELECT id, username, first_name, last_name, avatar, email, password, is_author, about FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.Avatar,
		&u.Email,
		&u.Password,
		&u.IsAuthor,
		&u.About,
	); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *Repo) FindByEmail(email string) (*model.UserDB, error) {
	u := &model.UserDB{}
	if err := r.Storage.DB.QueryRow(
		"SELECT id, username, first_name, last_name, avatar, email, password, is_author, about FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.Avatar,
		&u.Email,
		&u.Password,
		&u.IsAuthor,
		&u.About,
	); err != nil {
		return nil, err
	}

	return u, nil
}
