package main

import (
	"fmt"

	"github.com/matherique/project-manager/internal/project"
	"github.com/matherique/project-manager/pkg/config"
)

const doc_remove string = `
Usage: project remove|rm [name] 

Remove the project from the list
`

func cmd_remove(args []string, c config.Config, p project.Project) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("missing project name")
	}

	if err := p.Remove(args[0]); err != nil {
		return "", err
	}

	return "project removed with successfully", nil
}
