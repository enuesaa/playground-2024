package command

import (
	"fmt"
	"time"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/router"
	"github.com/enuesaa/codetrailer/internal/usecase"
	"github.com/urfave/cli/v2"
)

func NewExportCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:      "export",
		Args:      true,
		ArgsUsage: "<filename>",
		Before: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("<filename> is required.")
			}
			repos.Config.DocPath = c.Args().Get(0)

			return repos.Pw.Install()
		},
		Action: func(c *cli.Context) error {
			app := router.New(repos)
			go app.Start(":3000")

			time.Sleep(1 * time.Second)

			return usecase.ExportPdf(repos)
		},
	}

	return cmd
}
