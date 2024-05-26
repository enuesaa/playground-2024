package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:     "sample",
		Version: "0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello")
		},
	}

	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	app.Execute()
}
