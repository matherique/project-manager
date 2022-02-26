package main

import (
	"fmt"

	"github.com/matherique/project-manager/internal/project"
	"github.com/matherique/project-manager/pkg/cmd"
	"github.com/matherique/project-manager/pkg/config"
)

const doc_open string = `
Usage: project open [name] 

Execute the project script
`

func cmd_open(args []string, c config.Config, p project.Project) error {
	if len(args) == 0 {
		return fmt.Errorf("missing project name")
	}

	c.Load()

	name := args[0]

	if !project.Exists(name) {
		return fmt.Errorf("project not found")
	}

	return cmd.Exec(p.Path(n), "up")
}
