package main

import (
	"fmt"

	"github.com/afbjorklund/go-swhid/pkg/swhid"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:     "parse",
	Aliases: []string{"dir"},
	Short:   "Parse/pretty-print a (qualified) SWHID",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		swhid, err := swhid.Parse(args[0])
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("%s\n", swhid)
	},
}
