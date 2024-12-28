package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(&printCmd{}, "")
  
	flag.Parse()
	ctx := context.Background()
	status := subcommands.Execute(ctx)

	os.Exit(int(status))
}

type printCmd struct {
	capitalize bool
}

func (c *printCmd) Name() string {
	return "print"
}

func (c *printCmd) Synopsis() string {
	return "Print args to stdout."
}

func (c *printCmd) Usage() string {
	return `print [-capitalize] <some text>:
	Print args to stdout.
  `
}

func (c *printCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.capitalize, "capitalize", false, "capitalize output")
}

func (c *printCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if c.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()

	return subcommands.ExitSuccess
}
