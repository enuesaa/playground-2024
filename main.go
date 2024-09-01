package main

import (
	"log"
	"os"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/usecase"
	"github.com/urfave/cli/v2"
)

// Dir structure
//   README.md
//   steps/0-aa.md
//   steps/0.1-aa.md
//   steps/0.2.1-bb.md
//   steps/1-bb.md
//   steps/0-image-name.png
// 各ステップは細かく刻む
//
// ステップに必要な情報
//   priority
//   name, description

func main() {
	repos := repository.New()

	app := &cli.App{
		Name:    "codetrailer",
		Usage:   "A CLI tool to capture stdin/stdout and generate a step-by-step document.",
		Version: "0.0.1",
		Commands: []*cli.Command{
			// command.NewRecordCommand(repos),
			// command.NewPreviewCommand(repos),
			// command.NewExportPdfCommand(repos),
		},
		Action: func(c *cli.Context) error {
			if err := repos.Pw.Install(); err != nil {
				return err
			}

			go usecase.Prompt(repos)

			return usecase.LaunchMenu()
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

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
