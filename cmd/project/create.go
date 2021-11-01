package main

import (
	"fmt"

	"github.com/matherique/project-manager/internal/cmd"
	fc "github.com/matherique/project-manager/internal/file_config"
	"github.com/matherique/project-manager/internal/project"
	"github.com/matherique/project-manager/internal/utils"
)

const tpl = `#!/bin/bash

project={{.}}
location=

cd "$location"

_up() {
	tmux new-session -s $project -c $location -d

	# all tmux code goes here
}

_down() {
	# script to kill the tmux session gracifully
}

# ======= dont remove ======= 

if [ "$1" = "up ]; then
	_up
elif [ "$1" = "down" ]; then
	_down
else
	echo "usage: up|down"
fi

if [ ! -z $TMUX ];
then 
  tmux switch -t $project
else
  tmux attach -t $project
fi 

`
const doc_create string = `
Usage: project create [name] 

Create new project script
`

func cmd_create(a []string, c fc.FileConfig) error {
	c.Load()

	if len(a) == 0 {
		return fmt.Errorf("missing project name")
	}

	name := a[0]

	if project.Exists(c, name) {
		return fmt.Errorf("a project with this name already exists")
	}

	fp := project.Path(c, name)

	err := utils.CreateFile(fp, tpl)

	if err != nil {
		return err
	}

	edt := c.Get("editor")

	return cmd.Exec(edt, fp)
}
