package subscriptionsRepository

import (
	"database/sql"
	"errors"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ztrue/tracerr"
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgres(url string, maxIdleConns, maxOpenConns int) (*Postgres, error) {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)

	return &Postgres{DB: db}, nil
}

func (p Postgres) GetSubscriptionByUserAndAuthorID(userID, authorID uint64) (models.AuthorSubscription, error) {
	var s models.AuthorSubscription
	if err := p.DB.Get(&s, `
		SELECT "as".id, "as".author_id, "as".tier, "as".text, "as".title, "as".price, "as".img
		FROM subscriptions s 
		JOIN author_subscriptions "as" on "as".id = s.subscription_id
		WHERE s.subscriber_id = $1 
		AND s.author_id = $2
		`,
		userID,
		authorID,
	); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return models.AuthorSubscription{}, err
	}

	var f models.Follower
	if s.ID == 0 {
		if err := p.DB.Get(&f, `
			SELECT * 
			FROM followers
			WHERE follower_id = $1 and author_id = $2
			`,
			userID,
			authorID,
		); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return models.AuthorSubscription{}, err
		}

		return models.AuthorSubscription{
			AuthorID: f.AuthorID,
		}, nil
	}

	return s, nil
}

func (p Postgres) GetSubscriptionsByUserID(userID uint64) ([]models.AuthorSubscription, error) {
	var s []models.AuthorSubscription
	if err := p.DB.Select(&s, `
		SELECT author_subscriptions.id,
		       author_subscriptions.author_id,
		       author_subscriptions.title,
		       author_subscriptions.tier,
		       author_subscriptions.text,
		       author_subscriptions.price
		FROM subscriptions JOIN author_subscriptions on author_subscriptions.id = subscriptions.subscription_id
		WHERE subscriber_id = $1`,
		userID,
	); err != nil && err != sql.ErrNoRows {
		return nil, tracerr.Wrap(err)
	}

	var f []models.Follower
	if err := p.DB.Select(&f, `
		SELECT follower_id, author_id
		FROM followers
		WHERE follower_id = $1;
	`, userID); err != nil && err != sql.ErrNoRows {
		return nil, tracerr.Wrap(err)
	}

	for _, follow := range f {
		s = append(s, models.AuthorSubscription{
			ID:       0,
			AuthorID: follow.AuthorID,
		})
	}

	return s, nil
}

func (p Postgres) GetSubscriptionsByAuthorID(authorID uint64) ([]models.AuthorSubscription, error) {
	var s []models.AuthorSubscription
	if err := p.DB.Select(&s, `
		SELECT * 
		FROM author_subscriptions
		WHERE author_id = $1`,
		authorID,
	); err != nil {
		return nil, err
	}

	return s, nil
}

func (p Postgres) GetSubscriptionByID(id uint64) (models.AuthorSubscription, error) {
	var s models.AuthorSubscription
	if err := p.DB.Get(&s, `
		SELECT * 
		FROM author_subscriptions
		WHERE id = $1`,
		id,
	); err != nil {
		return models.AuthorSubscription{}, err
	}

	return s, nil
}

func (p Postgres) AddSubscription(sub models.AuthorSubscription) (uint64, error) {
	var id uint64
	err := p.DB.QueryRowx(`
			INSERT INTO author_subscriptions (author_id, img, title, tier, text, price) 
			VALUES ($1, $2, $3, $4, $5, $6) 
			RETURNING id`,
		sub.AuthorID,
		sub.Img,
		sub.Title,
		sub.Tier,
		sub.Text,
		sub.Price,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p Postgres) UpdateSubscription(sub models.AuthorSubscription) error {
	_, err := p.DB.NamedExec(
		`
		UPDATE author_subscriptions 
		SET author_id=:author_id,
		    img=:img,
		    tier=:tier,
		    title=:title,
		    text=:text,
		    price=:price
		WHERE id = :id`, sub)

	return err
}

func (p Postgres) DeleteSubscription(subID uint64) error {
	_, err := p.DB.Exec(`
		DELETE FROM author_subscriptions 
		WHERE id=$1`,
		subID,
	)

	return err
}
