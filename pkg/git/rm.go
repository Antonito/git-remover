package git

import (
	"fmt"
	"strings"

	"github.com/antonito/git-remover/internal/exec"
)

// Remove deletes a list of files from a git history
func Remove(files []string) error {
	if len(files) == 0 {
		return fmt.Errorf("no file specified")
	}
	fileList := strings.Join(files, " ")
	rmCmd := fmt.Sprintf("git filter-branch -f --index-filter \"git rm -rf --cached --ignore-unmatch %s\" -- --all", fileList)

	cmds := []string{
		// Remove files
		rmCmd,

		// Update history
		"rm -rf .git/refs/original/",
		"git reflog expire --expire=now --all",
		"git gc --prune=now",
		"git gc --aggressive --prune=now",
	}

	for _, currentCmd := range cmds {
		if err := exec.Shell(currentCmd); err != nil {
			fmt.Printf("Cannot reset repo ('%s' failed)\n", currentCmd)
			return err
		}
	}

	fmt.Println("Successful, please update your remote with `git push --all -f`")
	return nil
}
