package main

import (
	"fmt"

	"github.com/matherique/project-manager/internal/cmd"
	fc "github.com/matherique/project-manager/internal/file_config"
	"github.com/matherique/project-manager/internal/project"
)

const doc_stop string = `
Usage: project stop [name]

Stop de project with the given name.
`

func cmd_stop(args []string, c fc.FileConfig) error {
	if len(args) == 0 {
		return fmt.Errorf("missing project name")
	}

	c.Load()

	n := args[0]

	if !project.Exists(c, n) {
		return fmt.Errorf("project not found")
	}

	fp := project.Path(c, n)

	return cmd.Exec(fp, "down")
}
