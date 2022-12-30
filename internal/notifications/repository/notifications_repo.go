package notificationsRepository

import (
	"os"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sqlx.DB
}

func New(url string, maxOpenConns int) (*Postgres, error) {
	url += " user=" + os.Getenv("PG_USER") + " password=" + os.Getenv("PG_PASSWORD")

	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)

	return &Postgres{
		DB: db,
	}, nil
}

func (p *Postgres) GetNotificationsByUserID(userID uint64) ([]models.Notification, error) {
	notifications := make([]models.Notification, 0)
	err := p.DB.Select(&notifications, "SELECT * FROM notification WHERE data->'user_id' = $1", userID)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (p *Postgres) DeleteNotificationByUserID(userID uint64) error {
	_, err := p.DB.Exec("DELETE FROM notification WHERE data->'user_id' = $1", userID)
	if err != nil {
		return err
	}

	return nil
}
