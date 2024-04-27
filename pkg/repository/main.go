package repository

type Repos struct {
	Fs     FsRepositoryInterface
}

func New() Repos {
	return Repos{
		Fs:     &FsRepository{},
	}
}

func NewTestRepos() Repos {
	return Repos{
		Fs:     &FsMockRepository{},
	}
}
