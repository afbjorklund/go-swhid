package main

import (
	"fmt"

	"github.com/afbjorklund/go-swhid/pkg/swhid"
	"github.com/spf13/cobra"
)

var directoryCmd = &cobra.Command{
	Use:   "directory",
	Short: "Compute a directory SWHID recursively",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			directory, err := swhid.NewDirectoryFromDir(arg)
			if err != nil {
				continue
			}
			fmt.Printf("%s\n", directory.Swhid())
		}
	},
}
