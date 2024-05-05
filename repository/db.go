package repository

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

type DBRepositoryInterface interface {
	Open(path string) error
	Query(query string) (*sql.Rows, error)
	Close() error
}

type DBRepository struct {
	db *sql.DB
}

func (repo *DBRepository) dsn(path string) string {
	return fmt.Sprintf("file:%s?_fk=1", path)
}

func (repo *DBRepository) Open(path string) error {
	if repo.db != nil {
		return nil
	}
	db, err := sql.Open("sqlite", repo.dsn(path))
	if err != nil {
		return err
	}
	repo.db = db
	return nil
}

func (repo *DBRepository) Query(query string) (*sql.Rows, error) {
	return repo.db.Query(query)
}

func (repo *DBRepository) Close() error {
	if repo.db == nil {
		return nil
	}
	if err := repo.db.Close(); err != nil {
		return err
	}
	repo.db = nil
	return nil
}
