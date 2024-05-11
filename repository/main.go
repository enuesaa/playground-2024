package repository

type Repos struct {
	Fs     FsRepositoryInterface
}

func New() Repos {
	return Repos{
		Fs:     &FsRepository{},
	}
}
