package main

import (
	"fmt"
	"os"
	"strings"

	fc "github.com/matherique/project-manager/internal/file_config"
)

func cmd_list(args []string, c fc.FileConfig) error {
	c.Load()

	p := strings.Split(c.Get("projects"), ";")

	if len(p) == 0 {
		fmt.Fprintln(os.Stdout, "no project found")
	}

	fmt.Fprintln(os.Stdout, strings.Join(p, "\n"))

	return nil
}
