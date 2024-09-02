package command

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/urfave/cli/v2"
)

func NewPreviewCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "preview",
		Usage: "",
		Args: true,
		Action: func(c *cli.Context) error {
			if err := repos.Pw.Install(); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
