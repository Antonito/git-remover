package cmd

import (
	"fmt"
	"strings"

	"github.com/antonito/git-remover/internal/exec"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

func execWrapper(cmd string) error {
	stdout, stderr, err := exec.ShellOut(cmd)
	if err != nil {
		fmt.Println(stderr)
	} else {
		fmt.Println(stdout)
	}
	return err
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a file from the git history",
	Long:  `Remove a file from the git history`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := rmFromHistory(args); err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		resetLocalRepo()
	},
}

func rmFromHistory(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no file specified")
	}
	fileList := strings.Join(args, " ")

	cmd := fmt.Sprintf("git filter-branch -f --index-filter \"git rm -rf --cached --ignore-unmatch %s\" -- --all", fileList)
	return execWrapper(cmd)
}

func resetLocalRepo() {
	cmds := []string{
		"rm -rf .git/refs/original/",
		"git reflog expire --expire=now --all",
		"git gc --prune=now",
		"git gc --aggressive --prune=now",
	}
	for _, currentCmd := range cmds {
		if err := execWrapper(currentCmd); err != nil {
			fmt.Printf("Cannot reset repo ('%s' failed)\n", currentCmd)
			return
		}
	}

	fmt.Println("Successful, please update your remote with `git push --all -f`")
}
