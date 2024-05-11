package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

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
