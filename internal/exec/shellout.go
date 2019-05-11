package exec

import (
	"bytes"
	"os/exec"
)

const shellToUse = "/bin/bash"

// ShellOut execute a command and returns its stdout and stderr
func ShellOut(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
