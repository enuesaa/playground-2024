package repository

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/enuesaa/sample-books-api/internal/repository/dbq"
	_ "modernc.org/sqlite"
)

type DBRepoInterface interface {
	Migrate() error
	Open() error
	Close() error
	Query() *dbq.Queries
	Tx(ctx context.Context, fn func(ctx context.Context) error) error
}

//go:embed dbschema.sql
var dbMigrateQuery string

type DBRepo struct {
	db *sql.DB
    tx *sql.Tx
}

func (repo *DBRepo) dsn() string {
	return "file:data.db?_fk=1"
}

func (repo *DBRepo) Migrate() error {
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

func (repo *DBRepo) Open() error {
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

func (repo *DBRepo) Close() error {
	if repo.db == nil {
		return nil
	}
	if err := repo.db.Close(); err != nil {
		return err
	}
	repo.db = nil
	return nil
}

func (repo *DBRepo) Query() *dbq.Queries {
    if repo.tx != nil {
        return dbq.New(repo.tx)
    }
    return dbq.New(repo.db)
}

// see https://dev.appswingby.com/golang/golang%E3%81%A7db%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E5%AE%9F%E8%A3%85%E3%82%92%E8%BF%BD%E5%8A%A0%E3%81%99%E3%82%8B%E6%96%B9%E6%B3%95%E3%80%9Csqlc%E3%82%92/
func (repo *DBRepo) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
    repo.tx = tx

    defer func() {
        repo.tx = nil
    }()

    if err = fn(ctx); err != nil {
        if rberr := tx.Rollback(); rberr != nil {
			return rberr
        }
        return err
    }

    if err := tx.Commit(); err != nil {
        return err
    }
	return nil
}
