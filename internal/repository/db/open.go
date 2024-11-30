package db

import "database/sql"

func (repo *Repo) Open() error {
	if repo.db != nil {
		return nil
	}
	db, err := sql.Open("sqlite", repo.dsn())
	if err != nil {
		return err
	}
	repo.db = db
	return nil
}
