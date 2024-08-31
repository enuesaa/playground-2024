package repository

type Repos struct {
	Cmd CmdRepositoryInterface
	Fs  FsRepositoryInterface
	Log LogRepositoryInterface
	Pw PwRepositoryInterface
}

func New() Repos {
	return Repos{
		Cmd: &CmdRepository{},
		Fs:  &FsRepository{},
		Log: &LogRepository{},
		Pw: &PwRepository{},
	}
}
