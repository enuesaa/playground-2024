package srvbook

import "github.com/enuesaa/sample-books-api/internal/repository"

func New(repos repository.Repos) Srv {
	return Srv{
		repos: repos,
	}
}

type Srv struct {
	repos repository.Repos
}
