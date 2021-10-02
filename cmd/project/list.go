package main

import (
	"fmt"
	"os"
	"strings"

	fc "github.com/matherique/project-manager/internal/file_config"
	"github.com/matherique/project-manager/internal/project"
)

func cmd_list(args []string, c fc.FileConfig) error {
	c.Load()

	all := project.All(c)

	if len(all) == 0 {
		fmt.Fprintln(os.Stderr, "no project found")
	}

	fmt.Fprintln(os.Stdout, strings.Join(all, "\n"))

	return nil
}
