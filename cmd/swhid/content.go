package main

import (
	"fmt"
	"io"
	"os"

	"github.com/afbjorklund/go-swhid/pkg/swhid"
	"github.com/spf13/cobra"
)

var contentCmd = &cobra.Command{
	Use:     "content",
	Aliases: []string{"cnt"},
	Short:   "Compute SWHID from content from stdin",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		f := os.Stdin
		var err error
		if file != "" {
			f, err = os.Open(file)
			if err != nil {
				return
			}
			defer f.Close()
		}
		bytes, err := io.ReadAll(f)
		if err != nil {
			return
		}
		content := swhid.NewContent(bytes)
		fmt.Printf("%s\n", content.Swhid())
	},
}

var file string

func init() {
	contentCmd.PersistentFlags().StringVar(&file, "file", "", "Path to file")
}
