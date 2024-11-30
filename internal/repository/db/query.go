package db

import "github.com/enuesaa/sample-books-api/internal/repository/db/dbq"

func (repo *Repo) Query() *dbq.Queries {
	return dbq.New(repo.db)
}
