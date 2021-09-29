package main

import (
	"fmt"
	"strings"

	"github.com/matherique/project-manager/pkg/cmd"
	fc "github.com/matherique/project-manager/pkg/file_config"
)

func cmd_edit(args []string, c fc.FileConfig) error {
	if len(args) == 0 {
		return fmt.Errorf("missing project name")
	}

	c.Load()

	pl := strings.Split(c.Get("projects"), ";")

	if len(pl) == 0 {
		return fmt.Errorf("no project created")
	}

	var p string
	for _, v := range pl {
		if args[0] == v {
			p = v
		}
	}

	if p == "" {
		return fmt.Errorf("no project found with this name: %s", args[0])
	}

	return cmd.Exec(c.Get("editor"), args[0])
}
