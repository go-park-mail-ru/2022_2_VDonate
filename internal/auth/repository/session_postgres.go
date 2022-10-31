package sessionsRepository

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

func (r Postgres) Close() error {
	if err := r.DB.Close(); err != nil {
		return err
	}

	return nil
}

func (r Postgres) GetByUserID(id uint64) (models.Cookie, error) {
	var c models.Cookie
	err := r.DB.Get(&c, "SELECT value, user_id, expire_date FROM sessions WHERE user_id=$1;", id)
	if err != nil {
		return models.Cookie{}, err
	}

	return c, err
}

func (r Postgres) GetBySessionID(sessionID string) (models.Cookie, error) {
	var c models.Cookie
	err := r.DB.Get(&c, "SELECT value, user_id, expire_date FROM sessions WHERE value=$1;", sessionID)
	if err != nil {
		return models.Cookie{}, err
	}

	return c, err
}

func (r Postgres) GetByUsername(username string) (models.Cookie, error) {
	var c models.Cookie
	err := r.DB.Get(&c, `
		SELECT value, user_id, expire_date 
		FROM sessions 
		JOIN users on users.username = $1`,
		username,
	)
	if err != nil {
		return models.Cookie{}, err
	}

	return c, nil
}

func (r Postgres) CreateSession(cookie models.Cookie) (models.Cookie, error) {
	_, err := r.DB.Exec(
		"INSERT INTO sessions (value, user_id, expire_date) VALUES ($1, $2, $3);",
		cookie.Value,
		cookie.UserID,
		cookie.Expires,
	)
	if err != nil {
		return models.Cookie{}, err
	}

	return cookie, nil
}

func (r Postgres) DeleteByUserID(id uint64) error {
	_, err := r.DB.Exec("DELETE FROM sessions WHERE user_id=$1;", id)
	if err != nil {
		return err
	}

	return nil
}

func (r Postgres) DeleteBySessionID(sessionID string) error {
	_, err := r.DB.Exec("DELETE FROM sessions WHERE value=$1;", sessionID)
	if err != nil {
		return err
	}

	return nil
}
