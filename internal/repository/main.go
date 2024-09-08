package repository

type Repos struct {
	Cmd CmdRepositoryInterface
	Fs  FsRepositoryInterface
	Pw PwRepositoryInterface
	Prompt PromptRepositoryInterface
}

func New() Repos {
	return Repos{
		Cmd: &CmdRepository{},
		Fs:  &FsRepository{},
		Pw: &PwRepository{},
		Prompt: &PromptRepository{},
	}
}
