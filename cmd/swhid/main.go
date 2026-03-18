package main

import (
	"fmt"
	"os"

	"github.com/afbjorklund/go-swhid/pkg/swhid"
	"github.com/spf13/cobra"
)

var version string

var hashName string
var hashEncoding string

var rootCmd = &cobra.Command{
	Use:   "swhid",
	Short: "Compute and parse SWHIDs (ISO/IEC 18670)",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		swhid.Version = version
		if err := swhid.SetHash(hashName); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		if err := swhid.SetEncoding(hashEncoding); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(contentCmd)
	rootCmd.AddCommand(directoryCmd)
	if gitCmd != nil {
		rootCmd.AddCommand(gitCmd)
	}
	rootCmd.AddCommand(parseCmd)
	rootCmd.AddCommand(verifyCmd)

	rootCmd.PersistentFlags().StringVar(&version, "version", "1", "SWH version")
	rootCmd.PersistentFlags().StringVar(&hashName, "hash", "sha1cd", "Hash name")
	rootCmd.PersistentFlags().StringVar(&hashEncoding, "encoding", "hex", "Hash encoding")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
