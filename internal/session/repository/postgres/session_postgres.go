package sessionPostgres

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/session/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sqlx.DB
}

func New(URL string) (sessionRepository.API, error) {
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

func (r *Repository) GetByUserID(id uint) (*models.Cookie, error) {
	var c models.Cookie
	err := r.DB.Get(&c, "SELECT value, user_id, expire_date FROM sessions WHERE user_id=$1;", id)
	if err != nil {
		return nil, err
	}
	return &c, err
}

func (r *Repository) GetByValue(value string) (*models.Cookie, error) {
	var c models.Cookie
	err := r.DB.Get(&c, "SELECT value, user_id, expire_date FROM sessions WHERE value=$1;", value)
	if err != nil {
		return nil, err
	}
	return &c, err
}

func (r *Repository) Create(cookie *models.Cookie) (*models.Cookie, error) {
	_, err := r.DB.Exec(
		"INSERT INTO sessions (value, user_id, expire_date) VALUES ($1, $2, $3);",
		cookie.Value,
		cookie.UserID,
		cookie.Expires,
	)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func (r *Repository) DeleteByUserID(id uint) error {
	_, err := r.DB.Query("DELETE FROM sessions WHERE user_id=$1;", id)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) DeleteByValue(value string) error {
	_, err := r.DB.Query("DELETE FROM sessions WHERE value=$1;", value)
	if err != nil {
		return err
	}
	return nil
}
