package create

import (
	"fmt"
	"path"
	"strings"

	"github.com/matherique/project-manager/pkg/cmd"
	fc "github.com/matherique/project-manager/pkg/file_config"
)

const tpl = `#!/bin/bash

project={{.}}

tmux new-session -s $project
`

type create struct {
	c fc.FileConfig
}

func NewCreate(c fc.FileConfig) *create {
	ct := new(create)
	ct.c = c
	return ct
}

func (c *create) Exec(a []string) error {
	c.c.Load()

	if len(a) == 0 {
		return fmt.Errorf("missing project name")
	}

	name := a[0]

	fp := path.Join(c.c.Get("scripts"), name)

	err := createFile(fp)

	if err != nil {
		return err
	}

	edt := c.c.Get("editor")

	err = cmd.Exec(edt, fp)

	if err != nil {
		return err
	}

	p := c.c.Get("projects")

	var pl []string
	if p == "" {
		pl = []string{}
	} else {
		pl = strings.Split(p, ";")
	}

	pl = append(pl, name)

	c.c.Set("projects", strings.Join(pl, ";"))

	return c.c.Save()
}
