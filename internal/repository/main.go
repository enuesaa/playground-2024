package repository

import "github.com/enuesaa/sample-books-api/internal/repository/db"

func New() Repos {
	repos := Repos{
		DB: &db.Repo{},
	}

	return repos
}

type Repos struct {
	DB db.RepoInterface
}
