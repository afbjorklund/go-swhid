//go:build !git

package main

import (
	"github.com/spf13/cobra"
)

var gitCmd *cobra.Command = nil
