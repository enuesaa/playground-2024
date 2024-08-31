package command

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/usecase"
	"github.com/urfave/cli/v2"
)

func NewRecordCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "record",
		Usage: "Create new step or update existing steps.",
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

			go usecase.Prompt(repos)

			return usecase.LaunchMenu()
		},
	}

	return cmd
}
