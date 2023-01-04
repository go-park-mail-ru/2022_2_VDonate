package subscribersRepository

import (
	"database/sql"
	"os"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ztrue/tracerr"
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgres(url string, maxOpenConns int) (*Postgres, error) {
	url += " user=" + os.Getenv("PG_USER") + " password=" + os.Getenv("PG_PASSWORD")

	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)

	return &Postgres{DB: db}, nil
}

func (p Postgres) GetSubscribers(authorID uint64) ([]uint64, error) {
	var subscribers []uint64
	var followers []uint64
	err := p.DB.Select(&subscribers, `
		SELECT subscriber_id 
		FROM subscriptions 
		WHERE author_id=$1;`,
		authorID,
	)
	if err != nil {
		return nil, err
	}

	if err = p.DB.Select(&followers, `
		SELECT follower_id
		FROM followers
		WHERE author_id=$1;`,
		authorID,
	); err != nil {
		return nil, err
	}

	return append(subscribers, followers...), nil
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

	if _, err = p.DB.Exec(`
		DELETE FROM followers 
		WHERE author_id=$1 AND follower_id=$2`,
		authorID,
		userID,
	); err != nil {
		return err
	}

	return nil
}

func (p Postgres) Follow(subscriberID, authorID uint64) error {
	var sub models.Subscription
	if err := p.DB.Get(&sub, `
		SELECT author_id, subscriber_id, subscription_id 
		FROM subscriptions 
		WHERE subscriber_id=$1 AND author_id=$2;`,
		subscriberID,
		authorID,
	); err != nil && err != sql.ErrNoRows {
		return tracerr.Wrap(err)
	}

	if sub.SubscriberID == 0 && sub.AuthorID == 0 && sub.AuthorSubscriptionID == 0 {
		_, err := p.DB.Exec(`
			INSERT INTO followers (author_id, follower_id) 
			VALUES ($1, $2);`,
			authorID,
			subscriberID,
		)
		if err != nil {
			return tracerr.Wrap(err)
		}
	} else {
		return tracerr.Wrap(domain.ErrAlreadySubscribed)
	}

	return nil
}

func (p Postgres) PayAndSubscribe(payment models.Payment) error {
	var f models.Follower
	err := p.DB.Get(&f, `
			SELECT author_id, follower_id
			FROM followers
			WHERE follower_id=$1 AND author_id=$2;
		`,
		payment.FromID,
		payment.ToID,
	)
	if err != nil && err != sql.ErrNoRows {
		return tracerr.Wrap(err)
	}

	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}

	if f.AuthorID != 0 || f.FollowerID != 0 {
		if _, err = tx.Exec(`
			DELETE FROM followers 
        	WHERE follower_id=$1 AND author_id=$2`,
			payment.FromID,
			payment.ToID,
		); err != nil {
			return tracerr.Wrap(err)
		}
	}

	if err = tx.QueryRow(
		"INSERT INTO payments (id, to_id, from_id, sub_id, status) VALUES ($1, $2, $3, $4, $5) RETURNING time;",
		payment.ID,
		payment.ToID,
		payment.FromID,
		payment.SubID,
		payment.Status,
	).Scan(&payment.Time); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return errRollback
		}

		return err
	}

	if payment.Status == "REJECTED" || payment.Status == "EXPIRED" {
		if err = tx.Commit(); err != nil {
			return err
		}

		return nil
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
	_, err := p.DB.Exec(
		`
		UPDATE payments
		SET status=$1
		WHERE id=$2`, status, id,
	)
	if err != nil {
		return err
	}

	return nil
}
