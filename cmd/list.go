package cmd

import (
	"fmt"

	"github.com/antonito/git-remover/internal/exec"
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
		stdout, stderr, err := exec.ShellOut("git rev-list --objects --all | git cat-file --batch-check='%(objecttype) %(objectname) %(objectsize) %(rest)' | sed -n 's/^blob //p' | sort --numeric-sort --key=2 | cut -c 1-12,41- | $(command -v gnumfmt || echo numfmt) --field=2 --to=iec-i --suffix=B --padding=7 --round=nearest")
		if err != nil {
			fmt.Println(stderr)
		} else {
			fmt.Println(stdout)
		}
	},
}
