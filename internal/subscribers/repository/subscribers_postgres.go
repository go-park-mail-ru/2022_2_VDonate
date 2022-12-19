package subscribersRepository

import (
	"database/sql"

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

func (p Postgres) PayAndSubscribe(payment models.Payment) error {
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}

	if err = tx.QueryRow(
		"INSERT INTO payments (id, to_id, from_id, sub_id, status) VALUES ($1, $2, $3, $4, $5) RETURNING time;",
		payment.ID,
		payment.ToID,
		payment.FromID,
		payment.SubID,
		payment.Status,
	).Scan(&payment.Time); err != nil {
		return err
	}

	var sub models.Subscription
	if err = tx.QueryRow(
		`SELECT author_id, subscriber_id, subscription_id FROM subscriptions WHERE author_id=$1 AND subscriber_id=$2`, payment.ToID, payment.FromID,
	).Scan(&sub.AuthorID, &sub.SubscriberID, &sub.AuthorSubscriptionID); err != nil && err != sql.ErrNoRows {
		if errRollback := tx.Rollback(); errRollback != nil {
			return errRollback
		}

		return err
	}

	if err == sql.ErrNoRows {
		_, err = tx.Exec(`
			INSERT INTO subscriptions (author_id, subscriber_id, subscription_id) 
			VALUES ($1, $2, $3);`,
			payment.ToID,
			payment.FromID,
			payment.SubID,
		)
	} else {
		_, err = tx.Exec(`
			UPDATE subscriptions SET subscription_id=$3, date_created=now() WHERE author_id=$1 AND subscriber_id=$2`,
			payment.ToID,
			payment.FromID,
			payment.SubID,
		)
	}
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return errRollback
		}

		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (p Postgres) UpdateStatus(status string, id string) error {
	return p.DB.QueryRow(
		`
		UPDATE payments
		SET status=$1
		WHERE id=$2`, status, id,
	).Err()
}
