package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/matherique/project-manager/internal/cmd"
	fc "github.com/matherique/project-manager/internal/file_config"
	"github.com/matherique/project-manager/internal/utils"
)

func cmd_create(a []string, c fc.FileConfig) error {
	c.Load()

	if len(a) == 0 {
		return fmt.Errorf("missing project name")
	}

	name := a[0]

	fp := path.Join(c.Get("scripts"), name)

	err := utils.CreateFile(fp)

	if err != nil {
		return err
	}

	edt := c.Get("editor")

	err = cmd.Exec(edt, fp)

	if err != nil {
		return err
	}

	p := c.Get("projects")

	var pl []string
	if p == "" {
		pl = []string{}
	} else {
		pl = strings.Split(p, ";")
	}

	pl = append(pl, name)

	c.Set("projects", strings.Join(pl, ";"))

	return c.Save()
}
