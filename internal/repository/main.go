package repository

type Repos struct {
	Cmd    CmdRepositoryInterface
	Fs     FsRepositoryInterface
	Pw     PwRepositoryInterface
	Prompt PromptRepositoryInterface
	Config Config
}

func New() Repos {
	return Repos{
		Cmd:    &CmdRepository{},
		Fs:     &FsRepository{},
		Pw:     &PwRepository{},
		Prompt: &PromptRepository{},
		Config: Config{
			DocsPath: "docs",
		},
	}
}
