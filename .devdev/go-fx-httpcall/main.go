package main

import (
	"fmt"

	"github.com/enuesaa/.devdev/go-fx-httpcall/callerfx"
	"github.com/enuesaa/.devdev/go-fx-httpcall/clifx"
	"go.uber.org/fx"
)

func main() {
	fxapp := fx.New(
		clifx.Module,
		callerfx.Module,
		fx.Invoke(func(cli clifx.ICli) error {
			return cli.Launch()
		}),
		fx.Invoke(func(cli clifx.ICli, caller callerfx.ICaller, shutdowner fx.Shutdowner) error {
			fmt.Println(cli.GetUrl())
			body, err := caller.Run(cli.GetUrl())
			if err != nil {
				return err
			}
			fmt.Printf("%s", body)

			// shutdown fx.App
			if err := shutdowner.Shutdown(); err != nil {
				return err
			}

			return nil
		}),
		fx.NopLogger,
	)
	fxapp.Run()
}
