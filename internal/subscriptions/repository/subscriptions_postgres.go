package subscriptionsRepository

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgres(URL string) (*Postgres, error) {
	db, err := sqlx.Open("postgres", URL)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}

func (p *Postgres) GetSubscriptionsByUserID(userID uint64) ([]*models.AuthorSubscription, error) {
	var s []*models.AuthorSubscription
	if err := p.DB.Select(&s, `
		SELECT author_subscriptions.id,
		       author_subscriptions.author_id,
		       author_subscriptions.tier,
		       author_subscriptions.text,
		       author_subscriptions.price
		FROM subscriptions JOIN author_subscriptions on author_subscriptions.id = subscriptions.subscription_id
		WHERE subscriber_id = $1`,
		userID,
	); err != nil {
		return nil, err
	}

	return s, nil
}

func (p *Postgres) GetSubscriptionsByAuthorID(authorID uint64) ([]*models.AuthorSubscription, error) {
	var s []*models.AuthorSubscription
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

func (p *Postgres) GetSubscriptionsByID(ID uint64) (*models.AuthorSubscription, error) {
	var s models.AuthorSubscription
	if err := p.DB.Get(&s, `
		SELECT * 
		FROM author_subscriptions
		WHERE id = $1`,
		ID,
	); err != nil {
		return nil, err
	}

	return &s, nil
}

func (p *Postgres) AddSubscription(sub models.AuthorSubscription) (*models.AuthorSubscription, error) {
	err := p.DB.QueryRowx(`
		INSERT INTO author_subscriptions (author_id, tier, text, price) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id`,
		sub.AuthorID,
		sub.Tier,
		sub.Text,
		sub.Price,
	).Scan(&sub.ID)
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (p *Postgres) UpdateSubscription(sub *models.AuthorSubscription) (*models.AuthorSubscription, error) {
	_, err := p.DB.NamedExec(
		`
		UPDATE author_subscriptions 
		SET author_id=:author_id,
		    tier=:tier,
		    text=:text,
		    price=:price
		WHERE id = :id`, sub)
	if err != nil {
		return nil, err
	}
	return sub, err
}

func (p *Postgres) DeleteSubscription(subID uint64) error {
	_, err := p.DB.Query(`
		DELETE FROM author_subscriptions 
		WHERE id=$1`,
		subID,
	)
	if err != nil {
		return err
	}

	return nil
}
