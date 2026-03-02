package main

import (
	"fmt"

	"github.com/afbjorklund/go-swhid/pkg/swhid"
	"github.com/spf13/cobra"
)

var directoryCmd = &cobra.Command{
	Use:     "directory PATH",
	Aliases: []string{"dir"},
	Args:    cobra.ExactArgs(1),
	Short:   "Compute SWHID for a directory recursively",
	Run: func(cmd *cobra.Command, args []string) {
		swhid.DirectoryExclude = exclude
		for _, arg := range args {
			directory, err := swhid.NewDirectoryFromPath(arg)
			if err != nil {
				fmt.Printf("%s: %v\n", arg, err)
				continue
			}
			fmt.Printf("%s\n", directory.Swhid())
		}
	},
}

var exclude string

func init() {
	directoryCmd.PersistentFlags().StringVar(&exclude, "exclude", "", "Exclude files matching these suffixes (e.g., .tmp, .log)")

	directoryCmd.Flags().BoolVarP(&swhid.WriteObjects, "write", "w", false, "Write objects")
}
