package subscribersRepository

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

func (p Postgres) GetSubscribers(authorID uint64) ([]uint64, error) {
	var subscribers []uint64
	err := p.DB.Select(&subscribers, `
		SELECT subscriber_id 
		FROM subscriptions 
		WHERE author_id=$1`,
		authorID,
	)
	if err != nil {
		return nil, err
	}

	return subscribers, nil
}

func (p Postgres) Subscribe(subscription models.Subscription) error {
	_, err := p.DB.Exec(`
		INSERT INTO subscriptions (author_id, subscriber_id, subscription_id) 
		VALUES ($1, $2, $3)`,
		subscription.AuthorID,
		subscription.SubscriberID,
		subscription.AuthorSubscriptionID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p Postgres) Unsubscribe(userID, authorID uint64) error {
	_, err := p.DB.Exec(`
		DELETE FROM subscriptions 
		WHERE author_id=$1 AND subscriber_id=$2`,
		authorID,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}
