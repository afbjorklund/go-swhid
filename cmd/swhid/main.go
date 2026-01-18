package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "swhid",
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	return nil
}

func main() {
	rootCmd.AddCommand(contentCmd)
	rootCmd.AddCommand(directoryCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
