package command

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/urfave/cli/v2"
)

func NewInitCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "init",
		Usage: "Create new project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "project",
				Aliases: []string{"p"},
				Value: ".",
				Usage: "Project path",
			},
		},
		Action: func(c *cli.Context) error {
			// project := c.String("project")

			if err := repos.Pw.Install(); err != nil {
				return err
			}

			// download headless browser here.
			// touch README.md

			return nil
		},
	}

	return cmd
}
