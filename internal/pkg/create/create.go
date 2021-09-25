package create

import (
	"fmt"
	"html/template"
	"os"
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

	fn := a[0]
	fp := path.Join(c.c.Get("scripts"), fn)
	f, err := os.Create(fp)

	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}

	t := template.Must(template.New("project").Parse(tpl))

	err = t.Execute(f, fn)

	if err != nil {
		return fmt.Errorf("could not save template file: %v", err)
	}

	err = os.Chmod(fn, 0777)

	if err != nil {
		return fmt.Errorf("could not save template file: %v", err)
	}

	edt := c.c.Get("editor")

	return cmd.Exec(edt, fp)
}

func (c *create) CreateFile() error {
	return nil
}
