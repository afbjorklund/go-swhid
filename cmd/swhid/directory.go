package main

import (
	"fmt"

	"github.com/afbjorklund/go-swhid/pkg/swhid"
	"github.com/spf13/cobra"
)

var directoryCmd = &cobra.Command{
	Use:     "directory PATH",
	Aliases: []string{"dir"},
	Args:    cobra.MinimumNArgs(0),
	Short:   "Compute SWHID for a directory recursively",
	Run: func(cmd *cobra.Command, args []string) {
		swhid.DirectoryExcludes = excludes
		if archive != "" {
			directory, err := swhid.NewDirectoryFromTar(archive)
			if err != nil {
				fmt.Printf("%s: %v\n", archive, err)
				return
			}
			fmt.Printf("%s\n", directory.Swhid())
			return
		}
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

var archive string
var excludes []string

func init() {
	directoryCmd.PersistentFlags().StringVar(&archive, "archive", "", "Path to archive")
	directoryCmd.PersistentFlags().StringArrayVar(&excludes, "exclude", []string{".git"}, "Exclude files matching these suffixes (e.g., .tmp, .log)")

	directoryCmd.Flags().BoolVarP(&swhid.WriteObjects, "write", "w", false, "Write objects")
}
