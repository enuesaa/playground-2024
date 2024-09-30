package command

import (
	"fmt"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/router"
	"github.com/urfave/cli/v2"
)

func NewWriteCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name: "write",
		Args: true,
		ArgsUsage: "<filename>",
		Before: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("<filename> is required.")
			}
			repos.Config.DocPath = c.Args().Get(0)

			return nil
		},
		Action: func(c *cli.Context) error {
			app := router.New(repos)

			return app.Start(":3000")
		},
	}

	return cmd
}
