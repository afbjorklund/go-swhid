package main

import (
	"fmt"
	"os"

	"github.com/afbjorklund/go-swhid/pkg/swhid"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify PATH SWHID",
	Args:  cobra.ExactArgs(2),
	Short: "Verify that a file or directory matches a given SWHID",
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		st, err := os.Stat(path)
		if err != nil {
			fmt.Printf("%s: %v\n", path, err)
			return
		}
		var id *swhid.Swhid
		if st.IsDir() {
			dir, err := swhid.NewDirectoryFromPath(path)
			if err != nil {
				fmt.Printf("%s: %v\n", path, err)
				return
			}
			id = dir.Swhid()
		} else {
			cnt, err := swhid.NewContentFromPath(path)
			if err != nil {
				fmt.Printf("%s: %v\n", path, err)
				return
			}
			id = cnt.Swhid()
		}
		actual := id.String()
		expected := args[1]
		if actual == expected {
			fmt.Printf("✓ Verification successful: %s matches %s\n", path, expected)
		} else {
			fmt.Printf("✗ Verification failed: %s does not match %s\n", path, expected)
		}
	},
}
