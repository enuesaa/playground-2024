package command

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/usecase"
	"github.com/urfave/cli/v2"
)

func NewExportCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "export",
		Before: func(ctx *cli.Context) error {
			return repos.Pw.Install()
		},
		Action: func(c *cli.Context) error {
			return usecase.ExportPdf()
		},
	}

	return cmd
}
