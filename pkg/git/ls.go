package git

import (
	"github.com/antonito/git-remover/internal/exec"
)

// List the files commited on the repository
// TODO: Implement in pure Go
func List() error {
	err := exec.Shell("git rev-list --objects --all | git cat-file --batch-check='%(objecttype) %(objectname) %(objectsize) %(rest)' | sed -n 's/^blob //p' | sort --numeric-sort --key=2 | cut -c 1-12,41- | $(command -v gnumfmt || echo numfmt) --field=2 --to=iec-i --suffix=B --padding=7 --round=nearest")
	return err
}
