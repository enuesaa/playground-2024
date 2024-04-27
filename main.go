package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

type Data struct {
	A string `json:"a"`
}

func main() {
	app := &cobra.Command{
		Use:     "txtsout",
		Short:   "template engine",
		Version: "0.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			fbyte, err := os.ReadFile("testdata/data.json")
			if err != nil {
				return err
			}

			var data Data
			if err := json.Unmarshal(fbyte, &data); err != nil {
				return err
			}

			tmplbyte, err := os.ReadFile("testdata/a.txt")
			if err != nil {
				return err
			}

			tmpl, err := template.New("try").Parse(string(tmplbyte))
			if err != nil {
				return err
			}
		
			return tmpl.Execute(os.Stdout, data)
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
