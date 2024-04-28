package cli

import (
	"encoding/json"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

type Data struct {
	A string `json:"a"`
}

func Run() error {
	cmd := &cobra.Command{
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
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	cmd.PersistentFlags().SortFlags = false
	cmd.PersistentFlags().BoolP("help", "", false, "Show help information")
	cmd.PersistentFlags().BoolP("version", "", false, "Show version")

	return cmd.Execute()
}
