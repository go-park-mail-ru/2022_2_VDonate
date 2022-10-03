package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Open(dbName, dbURL string) error {
	db, err := sql.Open(dbName, dbURL)
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
