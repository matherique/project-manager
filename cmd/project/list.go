package main

import (
	"fmt"
	"os"
	"strings"

	fc "github.com/matherique/project-manager/pkg/file_config"
	"github.com/matherique/project-manager/pkg/project"
)

const doc_list string = `
Usage: project list|ls

List all available projects
`

func cmd_list(args []string, c fc.FileConfig) error {
	c.Load()

	all := project.All(c)

	if len(all) == 0 {
		return fmt.Errorf("no project found")
	}

	fmt.Fprintln(os.Stdout, strings.Join(all, "\n"))

	return nil
}
