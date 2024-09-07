package command

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/usecase"
	"github.com/urfave/cli/v2"
)

func NewWriteCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "write",
		Args: true,
		Before: func(ctx *cli.Context) error {
			return usecase.CreateRegistry(repos)
		},
		Action: func(c *cli.Context) error {
			go usecase.WriteStep(repos)

			return usecase.LaunchMenu()
		},
	}

	return cmd
}
