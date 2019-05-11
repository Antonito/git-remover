package exec

import "fmt"

// Shell executes a command, and prints on Stderr or Stdout
func Shell(cmd string) error {
	stdout, stderr, err := ShellOut(cmd)
	if err != nil {
		fmt.Println(stderr)
	} else {
		fmt.Println(stdout)
	}
	return err
}
