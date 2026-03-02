package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swhid",
	Short: "Compute and parse SWHIDs (ISO/IEC 18670)",
}

func main() {
	rootCmd.AddCommand(contentCmd)
	rootCmd.AddCommand(directoryCmd)
	rootCmd.AddCommand(parseCmd)
	rootCmd.AddCommand(verifyCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
