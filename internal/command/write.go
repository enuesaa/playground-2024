package command

import (

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/router"
	"github.com/urfave/cli/v2"
)

func NewWriteCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "write",
		Action: func(c *cli.Context) error {
			app := router.New(repos)

			return app.Start(":3000")
		},
	}

	return cmd
}
