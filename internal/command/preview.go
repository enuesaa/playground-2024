package command

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/urfave/cli/v2"
)

func NewPreviewCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name: "preview",
		Args: true,
		Before: func(ctx *cli.Context) error {
			return repos.Pw.Install()
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	return cmd
}
