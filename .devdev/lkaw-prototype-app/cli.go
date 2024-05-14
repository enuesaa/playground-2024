package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/enuesaa/lkaw/repository"
)

func init() {
	log.SetFlags(0)
}

// see https://github.com/alecthomas/kong/blob/2ab5733f117949ba9d1e24c1123d9c977d1808e6/global.go
func Parse(repos repository.Repos) (Cli, error) {
	var cli Cli

	parser, err := kong.New(
		&cli,
		kong.Name("lkaw"),
		kong.Description("A CLI tool to look at files anyway."),
	)
	if err != nil {
		return cli, err
	}
	if _, err = parser.Parse(os.Args[1:]); err != nil {
		return cli, err
	}

	if cli.Filenames[0] == "." {
		filenames, err := repos.Fs.ListFiles(".")
		if err != nil {
			return cli, err
		}
		cli.Filenames = filenames
	}

	return cli, nil
}


type Cli struct {
	Port      int         `help:"port to serve." default:"3000"`
	Filenames []string    `arg:"" required:"" name:"filename" help:"Files to open"`
	Version   VersionFlag `name:"version" short:"v" help:"Print version"`
}

type VersionFlag struct{}

func (v *VersionFlag) Decode(ctx *kong.DecodeContext) error {
	return nil
}
func (v *VersionFlag) IsBool() bool {
	return true
}
func (v *VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println("0.0.1")
	app.Exit(0)
	return nil
}
