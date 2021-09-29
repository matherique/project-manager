package main

import (
	"fmt"
	"path"

	"github.com/matherique/project-manager/internal/cmd"
	fc "github.com/matherique/project-manager/internal/file_config"
)

func cmd_open(a []string, c fc.FileConfig) error {
	if len(a) == 0 {
		return fmt.Errorf("missing project name")
	}

	c.Load()
	sp := c.Get("scripts")
	fp := path.Join(sp, a[0])

	return cmd.Exec(fp)
}
