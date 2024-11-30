package db

import (
	"context"
	"database/sql"
)

func (repo *Repo) Migrate() error {
	db, err := sql.Open("sqlite", repo.dsn())
	if err != nil {
		return err
	}
	ctx := context.Background()
	if _, err := db.ExecContext(ctx, dbMigrateQuery); err != nil {
		return err
	}
	return nil
}
