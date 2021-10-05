package main

import (
	"fmt"

	fc "github.com/matherique/project-manager/internal/file_config"
	"github.com/matherique/project-manager/internal/project"
)

func cmd_remove(args []string, c fc.FileConfig) error {
	if len(args) == 0 {
		return fmt.Errorf("missing project name")
	}

	return project.Remove(c, args[0])
}
