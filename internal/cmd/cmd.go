package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec(c string, args ...string) error {
	p, err := exec.LookPath(c)

	if err != nil {
		return fmt.Errorf("command not found, set a valid editor")
	}

	a := append([]string{c}, args...)

	cmd := &exec.Cmd{
		Path:   p,
		Args:   a,
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

	return cmd.Run()
}
