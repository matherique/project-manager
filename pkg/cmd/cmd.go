package cmd

import (
	"os"
	"os/exec"
)

func Exec(c string, args ...string) error {
	p, err := exec.LookPath(c)

	if err != nil {
		return err
	}

	a := append([]string{c}, args...)

	cmd := &exec.Cmd{
		Path:   p,
		Args:   a,
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

	err = cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
