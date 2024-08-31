package repository

type Repos struct {
	Cmd CmdRepositoryInterface
	Fs  FsRepositoryInterface
	Log LogRepositoryInterface
}

func New() Repos {
	return Repos{
		Cmd: &CmdRepository{},
		Fs:  &FsRepository{},
		Log: &LogRepository{},
	}
}
