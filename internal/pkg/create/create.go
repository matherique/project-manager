package create

import (
	"fmt"
	"path"

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

	fp := path.Join(c.c.Get("scripts"), a[0])

	err := createFile(fp)

	if err != nil {
		return err
	}

	edt := c.c.Get("editor")

	return cmd.Exec(edt, fp)
}
