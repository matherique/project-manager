package main

import (
	"fmt"
	"os"

	fc "github.com/matherique/project-manager/internal/file_config"
)

func cmd_help(_ []string, c fc.FileConfig) error {
	doc := `
Usage: project [command] [args]

> A project setup manager

Commands:
 create   Create a new project
 list     List all created projects
 config   List all configuration
 edit     Edit an project setup file
 remove   Remove and project
 fzf      List all created projects using fzf and execute what you choose

Use 'project [command] help' for more information about the command`
	fmt.Fprintln(os.Stdout, doc)
	return nil
}
