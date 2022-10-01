package storage

import (
	"database/sql"
	storage_config "github.com/go-park-mail-ru/2022_2_VDonate/internal/storages/config"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB     *sql.DB
	config *storage_config.Config
}

func New(config *storage_config.Config) *Storage {
	return &Storage{config: config}
}

func (s *Storage) Open(dbName string) error {
	db, err := sql.Open(dbName, s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.DB = db

	return nil
}

func (s *Storage) Close() error {
	if err := s.DB.Close(); err != nil {
		return err
	}
	return nil
}
