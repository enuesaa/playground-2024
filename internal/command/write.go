package command

import (
	"fmt"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/usecase"
	"github.com/urfave/cli/v2"
)

func NewWriteCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "write",
		Args: true,
		ArgsUsage: "<path>",
		Before: func(ctx *cli.Context) error {
			if ctx.Args().Len() == 0 {
				return fmt.Errorf("path is required")
			}
			return usecase.CreateRegistry(repos)
		},
		Action: func(c *cli.Context) error {
			path := c.Args().First()

			go usecase.Write(repos, path)

			return usecase.LaunchMenu()
		},
	}

	return cmd
}
