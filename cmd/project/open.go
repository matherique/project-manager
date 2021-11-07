package main

import (
	"fmt"

	"github.com/matherique/project-manager/internal/cmd"
	fc "github.com/matherique/project-manager/internal/file_config"
	"github.com/matherique/project-manager/internal/project"
)

const doc_open string = `
Usage: project open [name] 

Execute the project script
`

func cmd_open(a []string, c fc.FileConfig) error {
	if len(a) == 0 {
		return fmt.Errorf("missing project name")
	}

	c.Load()

	n := a[0]

	if !project.Exists(c, n) {
		return fmt.Errorf("project not found")
	}

	fp := project.Path(c, n)

	return cmd.Exec(fp, "up")
}
