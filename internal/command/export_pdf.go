package command

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/urfave/cli/v2"
)

func NewExportPdfCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "export-pdf",
		Usage: "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "project",
				Aliases: []string{"p"},
				Value: ".",
				Usage: "Project path",
			},
		},
		Action: func(c *cli.Context) error {
			// port := c.Int("port")
			return nil
		},
	}

	return cmd
}
