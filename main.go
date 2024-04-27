package main

import (
	"log"
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:     "txtsout",
		Short:   "template engine",
		Version: "0.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("a")
			return nil
		},
	}

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceErrors = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	if err := app.Execute(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
