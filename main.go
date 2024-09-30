package main

import (
	"log"
	"os"

	"github.com/enuesaa/codetrailer/internal/command"
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/urfave/cli/v2"
)

func main() {
	repos := repository.New()

	app := &cli.App{
		Name:    "codetrailer",
		Usage:   "A CLI tool to generate a step-by-step document.",
		Version: "0.0.1",
		Commands: []*cli.Command{
			command.NewWriteCommand(repos),
			command.NewExportCommand(repos),
		},
	}

	// disable default
	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		return err
	}
	app.HideHelpCommand = true
	cli.AppHelpTemplate = `{{.Usage}}

USAGE:
	{{.HelpName}} {{if .VisibleFlags}}[flags]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
FLAGS:{{range .Flags}}
	{{.}}{{end}}
{{end}}`

	cli.CommandHelpTemplate = `{{template "helpNameTemplate" .}}
 
USAGE:
	{{template "usageTemplate" .}}{{if .Category}}
 
CATEGORY:
	{{.Category}}{{end}}{{if .Description}}
 
DESCRIPTION:
	{{template "descriptionTemplate" .}}{{end}}

FLAGS:{{range $i, $e := .VisibleFlags}}
	{{wrap $e.String 6}}
{{end}}`

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
