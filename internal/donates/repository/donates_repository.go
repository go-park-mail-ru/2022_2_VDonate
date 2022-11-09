package donatesRepository

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jmoiron/sqlx"
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

func (p *Postgres) SendDonate(donate models.Donate) (models.Donate, error) {
	err := p.DB.QueryRowx(`
		INSERT INTO donates (user_id, author_id, price)
		VALUES ($1, $2, $3) RETURNING id;`,
		donate.UserID,
		donate.AuthorID,
		donate.Price,
	).Scan(&donate.ID)
	if err != nil {
		return models.Donate{}, err
	}
	return donate, nil
}

func (p Postgres) GetDonatesByUserID(userID uint64) ([]models.Donate, error) {
	var donates []models.Donate
	if err := p.DB.Select(&donates, "SELECT * FROM donates WHERE user_id=$1;", userID); err != nil {
		return nil, err
	}
	// err := p.DB.Get(&donates, `SELECT * FROM donates WHERE user_id=$1;`, userID)
	// if err != nil {
	// 	return []models.Donate{}, err
	// }
	return donates, nil
}

func (p Postgres) GetDonateByID(donateID uint64) (models.Donate, error) {
	var donate models.Donate
	err := p.DB.Get(&donate, `SELECT * FROM donates WHERE id=$1;`, donateID)
	if err != nil {
		return models.Donate{}, err
	}
	return donate, nil
}
