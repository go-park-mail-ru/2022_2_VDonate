package userRepository

import (
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
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

func (r Postgres) Create(user model.User) (uint64, error) {
	var id uint64
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}

	if err = tx.QueryRow(
		`
			INSERT INTO users (username, email) 
			VALUES ($1, $2) 
			RETURNING id;`,
		user.Username,
		user.Email,
	).Scan(&id); err != nil {
		return 0, err
	}

	if _, err = tx.Exec(
		`
			INSERT INTO user_info (user_id, avatar, password, is_author, about) 
			VALUES ($1, $2, $3, $4, $5);`,
		id,
		user.Avatar,
		user.Password,
		user.IsAuthor,
		user.About,
	); err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return 0, errTx
		}
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (r Postgres) GetByUsername(username string) (model.User, error) {
	var u model.User
	if err := r.DB.Get(
		&u,
		`
		SELECT id, username, email 
		FROM users 
		WHERE username = $1;`,
		username,
	); err != nil {
		return model.User{}, err
	}

	if err := r.DB.Get(
		&u,
		`
		SELECT avatar, password, is_author, about
		FROM user_info 
		WHERE user_id = $1;`,
		u.ID,
	); err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (r Postgres) GetByID(id uint64) (model.User, error) {
	var u model.User
	if err := r.DB.Get(
		&u,
		`
		SELECT id, username, email 
		FROM users 
		WHERE id = $1;`,
		id,
	); err != nil {
		return model.User{}, err
	}

	if err := r.DB.Get(
		&u,
		`
		SELECT avatar, password, is_author, about
		FROM user_info 
		WHERE user_id = $1;`,
		u.ID,
	); err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (r Postgres) GetByEmail(email string) (model.User, error) {
	var u model.User
	if err := r.DB.Get(
		&u,
		`
		SELECT id, username, email 
		FROM users 
		WHERE email = $1;`,
		email,
	); err != nil {
		return model.User{}, err
	}

	if err := r.DB.Get(
		&u,
		`
		SELECT avatar, password, is_author, about
		FROM user_info 
		WHERE user_id = $1;`,
		u.ID,
	); err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (r Postgres) GetBySessionID(sessionID string) (model.User, error) {
	var u model.User
	if err := r.DB.Get(
		&u,
		`
		SELECT id, username, email
		FROM users JOIN sessions ON sessions.user_id = id
    	WHERE sessions.value = $1;`,
		sessionID,
	); err != nil {
		return model.User{}, err
	}

	if err := r.DB.Get(
		&u,
		`
		SELECT avatar, password, is_author, about
		FROM user_info 
		WHERE user_id = $1;`,
		u.ID,
	); err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (r Postgres) GetUserByPostID(postID uint64) (model.User, error) {
	var user model.User
	if err := r.DB.Get(&user, `
		SELECT id, username, email
		FROM posts 
		JOIN users on users.id = posts.user_id 
		WHERE posts.post_id = $1`, postID,
	); err != nil {
		return model.User{}, err
	}

	if err := r.DB.Get(
		&user,
		`
		SELECT avatar, password, is_author, about
		FROM user_info 
		WHERE user_id = $1;`,
		user.ID,
	); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r Postgres) Update(user model.User) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(
		`
		UPDATE users 
		SET username=$1,
		    email=$2
		WHERE id = $3`,
		user.Username,
		user.Email,
		user.ID,
	); err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return errTx
		}
		return err
	}

	if _, err = tx.Exec(
		`
		UPDATE user_info 
		SET avatar=$1,
		    password=$2,
		    is_author=$3,
		    about=$4
		WHERE user_id = $5`,
		user.Avatar,
		user.Password,
		user.IsAuthor,
		user.About,
		user.ID,
	); err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return errTx
		}
		return err
	}

	return tx.Commit()
}

func (r Postgres) DeleteByID(id uint64) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	if _, err = tx.Exec(`
		DELETE FROM users WHERE id=$1;`,
		id,
	); err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return errTx
		}
		return err
	}

	if _, err = tx.Exec(`
		DELETE FROM user_info WHERE user_id=$1;`,
		id,
	); err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return errTx
		}
		return err
	}

	return tx.Commit()
}

func (r Postgres) GetAllAuthors() ([]model.User, error) {
	var u []model.User
	if err := r.DB.Select(
		&u,
		`
		SELECT * FROM users;`,
	); err != nil {
		return nil, err
	}

	if err := r.DB.Get(
		&u,
		`
		SELECT avatar, is_author, about
		FROM user_info 
		WHERE is_author = true;`,
	); err != nil {
		return nil, err
	}

	return u, nil
}

func (r Postgres) GetAuthorByUsername(username string) ([]model.User, error) {
	var u []model.User
	if err := r.DB.Select(
		&u,
		`
		SELECT *
		FROM users
		WHERE similarity(username, $1) > 0
		ORDER BY similarity(username, $1) DESC, username;`,
		username,
	); err != nil {
		return []model.User{}, err
	}

	for index, user := range u {
		if err := r.DB.Get(
			&u[index],
			`
			SELECT avatar, is_author, about
			FROM user_info 
			WHERE user_id = $1 AND is_author = true;`,
			user.ID,
		); err != nil {
			return []model.User{}, err
		}
	}

	return u, nil
}
