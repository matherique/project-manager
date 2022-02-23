package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	fc "github.com/matherique/project-manager/pkg/config"
	"github.com/matherique/project-manager/internal/project"
)

const doc_fzf string = `
Usage: project fzf

List all project with fzf, and open if selected
`

func cmd_fzf(_ []string, c fc.FileConfig) error {
	projects := project.All(c)

	if len(projects) == 0 {
		return fmt.Errorf("no project found")
	}

	p, err := exec_fzf(projects)

	if err != nil {
		return err
	}

	if p == "" {
		return nil
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
