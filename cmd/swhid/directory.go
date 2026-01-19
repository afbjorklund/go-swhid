package main

import (
	"fmt"

	"github.com/afbjorklund/go-swhid/pkg/swhid"
	"github.com/spf13/cobra"
)

var directoryCmd = &cobra.Command{
	Use:     "directory",
	Aliases: []string{"dir"},
	Short:   "Compute a directory SWHID recursively",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			directory, err := swhid.NewDirectoryFromDir(arg)
			if err != nil {
				fmt.Printf("%s: %v\n", arg, err)
				continue
			}
			fmt.Printf("%s\n", directory.Swhid())
		}
	},
}
