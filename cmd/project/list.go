package main

import (
	"fmt"
	"strings"

	"github.com/matherique/project-manager/internal/project"
	"github.com/matherique/project-manager/pkg/config"
)

const doc_list string = `
Usage: project list|ls

List all available projects
`

func cmd_list(args []string, c config.Config, p project.Project) (string, error) {
	c.Load()

	projects := p.All()

	if len(projects) == 0 {
		return "", fmt.Errorf("no project found")
	}

	return strings.Join(projects, "\n"), nil
}
