package command

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/usecase"
	"github.com/urfave/cli/v2"
)

func NewWriteCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "write",
		Usage: "",
		Args: true,
		Action: func(c *cli.Context) error {
			// start prompt
			return usecase.LaunchMenu()
		},
	}

	return cmd
}
