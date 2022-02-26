package main

import (
	"fmt"

	"github.com/matherique/project-manager/internal/project"
	"github.com/matherique/project-manager/pkg/cmd"
	"github.com/matherique/project-manager/pkg/config"
)

const doc_edit string = `
Usage: project edit [name] 

Open project script file to edit
`

func cmd_edit(args []string, c config.Config, p project.Project) error {
	if len(args) == 0 {
		return fmt.Errorf("missing project name")
	}

	c.Load()

	projects := p.All()

	if len(projects) == 0 {
		return fmt.Errorf("no project created")
	}

	name := args[0]

	if !p.Exists(name) {
		return fmt.Errorf("no project found with this name: %s", name)
	}

	return cmd.Exec(c.Get("editor"), p.Path(name))
}
