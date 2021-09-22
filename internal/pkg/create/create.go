package create

import (
	"html/template"
	"log"
	"os"
	"path"

	"github.com/matherique/project-manager/pkg/cmd"
	"github.com/matherique/project-manager/pkg/config"
)

const tpl = `#!/bin/bash

project={{.}}

tmux new-session -s $project
`

type create struct {
  c config.Config
}

func NewCreate(c config.Config) *create {
  ct := new(create)
  ct.c = c
  return ct
}

func (c *create) Exec(a []string) {
	if len(a) == 0 {
		log.Fatalf("missing project name")
	}

  c.c.Load()

  fn := a[0]
  fp := path.Join(c.c.Get("scripts"), fn)
  f, err := os.Create(fp)

	if err != nil {
    log.Fatalf("could not create file: %v", err)
	}

  t := template.Must(template.New("project").Parse(tpl))
  
  err = t.Execute(f, fn)

	if err != nil {
    log.Fatalf("could not save template file: %v", err)
	}

  err = os.Chmod(fn, 0777)

	if err != nil {
    log.Fatalf("could not save template file: %v", err)
	}

  edt := c.c.Get("editor")

  err = cmd.Exec(edt, fp)

	if err != nil {
    log.Fatalf("could not run the command: %v", err)
	}

}

