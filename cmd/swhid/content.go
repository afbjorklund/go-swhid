package main

import (
	"fmt"
	"io"
	"os"

	"github.com/afbjorklund/go-swhid/pkg/swhid"
	"github.com/spf13/cobra"
)

var contentCmd = &cobra.Command{
	Use:   "content",
	Short: "Compute a content SWHID from stdin",
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return
		}
		content := swhid.NewContent(bytes)
		fmt.Printf("%s\n", content.Swhid())
	},
}
