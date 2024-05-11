package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type Version struct{}
func (v *Version) Decode(ctx *kong.DecodeContext) error {
	return nil
}
func (v *Version) IsBool() bool {
	return true
}
func (v *Version) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println("0.0.1")
	app.Exit(0)
	return nil
}
