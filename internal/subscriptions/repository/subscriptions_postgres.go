package subscriptionsRepository

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

func (p *Postgres) GetSubscriptionsByUserID(userID uint64) ([]models.AuthorSubscription, error) {
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
	); err != nil {
		return nil, err
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
