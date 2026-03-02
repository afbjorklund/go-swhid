//go:build git
package main

import (
	"fmt"

	"github.com/afbjorklund/go-swhid/pkg/swhid"
	"github.com/spf13/cobra"
)

var gitRevisionCmd = &cobra.Command{
	Use:   "revision REPO",
	Args:  cobra.MinimumNArgs(1),
	Short: "Compute SWHID for a revision/commit",
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := swhid.NewRepository(args[0])
		if err != nil {
			fmt.Printf("%s: %v\n", args[0], err)
			return
		}
		var revision *swhid.Revision
		if len(args) < 2 {
			/*
				head, err := repo.Head()
				if err != nil {
					return
				}
				fmt.Printf("%s\n", head)
			*/
			revision, err = repo.NewRevisionFromHead()
			if err != nil {
				return
			}
		} else {
			revision, err = repo.NewRevisionFromHash(args[1])
			if err != nil {
				return
			}

		}
		fmt.Printf("%s\n", revision.Swhid())
	},
}

var gitReleaseCmd = &cobra.Command{
	Use:   "release REPO TAG",
	Args:  cobra.ExactArgs(2),
	Short: "Compute SWHID for a release/tag",
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := swhid.NewRepository(args[0])
		if err != nil {
			fmt.Printf("%s: %v\n", args[0], err)
			return
		}
		/*
			tag, err := repo.Tag(args[1])
			if err != nil {
				fmt.Printf("%s: %v\n", args[1], err)
				return
			}
			fmt.Printf("%s\n", tag)
		*/
		release, err := repo.NewReleaseFromTag(args[1])
		if err != nil {
			return
		}
		fmt.Printf("%s\n", release.Swhid())
	},
}

var gitSnapshotCmd = &cobra.Command{
	Use:   "snapshot REPO",
	Args:  cobra.ExactArgs(1),
	Short: "Compute snapshot SWHID for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := swhid.NewRepository(args[0])
		if err != nil {
			fmt.Printf("%s: %v\n", args[0], err)
			return
		}
		/*
			branches, err := repo.Branches()
			if err != nil {
				fmt.Printf("%s: %v\n", args[1], err)
				return
			}
			for _, branch := range branches {
				fmt.Printf("%s\n", branch)
			}
		*/
		snapshot, err := repo.NewSnapshot()
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Printf("%s\n", snapshot.Swhid())
	},
}

var gitTagsCmd = &cobra.Command{
	Use:   "tags REPO",
	Args:  cobra.ExactArgs(1),
	Short: "List all tags in a repository",
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := swhid.NewRepository(args[0])
		if err != nil {
			fmt.Printf("%s: %v\n", args[0], err)
			return
		}
		tags, err := repo.Tags()
		if err != nil {
			fmt.Printf("%s: %v\n", args[1], err)
			return
		}
		for _, tag := range tags {
			fmt.Printf("%s\n", tag)
		}
	},
}

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Git repository SWHID computation (requires -tags git)",
}

func init() {
	/*
	        rootCmd.AddCommand(gitRevisionCmd)
		rootCmd.AddCommand(gitReleaseCmd)
		rootCmd.AddCommand(gitSnapshotCmd)
	*/

	gitCmd.AddCommand(gitRevisionCmd)
	gitCmd.AddCommand(gitReleaseCmd)
	gitCmd.AddCommand(gitSnapshotCmd)
	gitCmd.AddCommand(gitTagsCmd)
}
