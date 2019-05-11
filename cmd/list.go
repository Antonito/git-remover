package cmd

import (
	"os"

	"github.com/antonito/git-remover/pkg/git"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files in repository",
	Long:  `List all the files commited on the repository`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := git.List(); err != nil {
			os.Exit(1)
		}
	},
}
