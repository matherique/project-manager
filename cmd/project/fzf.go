package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	fc "github.com/matherique/project-manager/internal/file_config"
	"github.com/matherique/project-manager/internal/project"
)

func cmd_fzf(_ []string, c fc.FileConfig) error {
	projects := project.All(c)

	p, err := exec_fzf(projects)

	if err != nil {
		return err
	}

	return cmd_open([]string{p}, c)
}

func exec_fzf(source []string) (string, error) {
	path, err := exec.LookPath("fzf")

	if err != nil {
		return "", fmt.Errorf("fzf not instaled")
	}

	cmd := exec.Command(path, "--reverse")
	cmd.Stderr = os.Stderr
	in, _ := cmd.StdinPipe()

	errCh := make(chan error, 1)

	go func() {
		for _, src := range source {
			fmt.Fprintln(in, src)
		}

		errCh <- nil
		in.Close()
	}()

	err = <-errCh

	if err != nil {
		return "", err
	}

	result, _ := cmd.Output()

	return strings.Trim(string(result), "\n"), nil
}
