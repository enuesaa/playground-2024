package command

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/router"
	"github.com/urfave/cli/v2"
)

func NewWriteCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name: "write",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "docs",
				Value:       "docs",
				Usage:       "docs dir path",
				Destination: &repos.Config.DocsPath,
			},
			// TODO
			// &cli.StringFlag{
			// 	Name:        "project",
			// 	Value:       "project",
			// 	Usage:       "project dir path",
			// 	Destination: &docspath,
			// },
		},
		Action: func(c *cli.Context) error {
			app := router.New(repos)

			return app.Start(":3000")
		},
	}

	return cmd
}
