package db

import (
	"database/sql"
	_ "embed"

	"github.com/enuesaa/sample-books-api/internal/repository/db/dbq"
	_ "modernc.org/sqlite"
)

type RepoInterface interface {
	Migrate() error
	Open() error
	Close() error
	Query() *dbq.Queries
}

//go:embed dbschema.sql
var dbMigrateQuery string

type Repo struct {
	db *sql.DB
}
