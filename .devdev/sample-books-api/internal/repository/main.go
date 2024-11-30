package repository

func New() Repos {
	repos := Repos{
		DB: &DBRepo{},
	}

	return repos
}

type Repos struct {
	DB DBRepoInterface
}
