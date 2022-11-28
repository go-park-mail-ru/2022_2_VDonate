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
	return r.DB.Close()
}

func (r Postgres) GetByUserID(id uint64) (models.Cookie, error) {
	var c models.Cookie
	err := r.DB.Get(&c, "SELECT value, user_id, expire_date FROM sessions WHERE user_id=$1;", id)

	return c, err
}

func (r Postgres) GetBySessionID(sessionID string) (models.Cookie, error) {
	var c models.Cookie
	err := r.DB.Get(&c, "SELECT value, user_id, expire_date FROM sessions WHERE value=$1;", sessionID)

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

	return c, err
}

func (r Postgres) CreateSession(cookie models.Cookie) (models.Cookie, error) {
	_, err := r.DB.Exec(
		"INSERT INTO sessions (value, user_id, expire_date) VALUES ($1, $2, $3);",
		cookie.Value,
		cookie.UserID,
		cookie.Expires,
	)

	return cookie, err
}

func (r Postgres) DeleteByUserID(id uint64) error {
	_, err := r.DB.Exec("DELETE FROM sessions WHERE user_id=$1;", id)

	return err
}

func (r Postgres) DeleteBySessionID(sessionID string) error {
	_, err := r.DB.Exec("DELETE FROM sessions WHERE value=$1;", sessionID)

	return err
}
