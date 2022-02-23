package main

import (
	"fmt"
	"os"

	fc "github.com/matherique/project-manager/pkg/file_config"
	"github.com/matherique/project-manager/pkg/project"
)

const doc_remove string = `
Usage: project remove|rm [name] 

Remove the project from the list
`

func cmd_remove(args []string, c fc.FileConfig) error {
	if len(args) == 0 {
		return fmt.Errorf("missing project name")
	}

	err := project.Remove(c, args[0])

	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "project removed with successfully")

	return nil
}
