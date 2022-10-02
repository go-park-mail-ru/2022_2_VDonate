package user_repo

import (
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models/user"
	storage "github.com/go-park-mail-ru/2022_2_VDonate/internal/storages"
)

type Repo struct {
	Storage *storage.Storage
}

func New(s *storage.Storage) *Repo {
	return &Repo{Storage: s}
}

func (r *Repo) Create(u *model.User) error {
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

func (r *Repo) FindByUsername(username string) (*model.User, error) {
	u := &model.User{}
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

func (r *Repo) FindByID(id uint) (*model.User, error) {
	u := &model.User{}
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

func (r *Repo) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
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
