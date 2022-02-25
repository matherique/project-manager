package main

import (
	"fmt"

	"github.com/matherique/project-manager/internal/project"
	"github.com/matherique/project-manager/pkg/cmd"
	"github.com/matherique/project-manager/pkg/config"
	"github.com/matherique/project-manager/pkg/utils"
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

if [ "$1" = "up" ]; then
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

func cmd_create(args []string, c config.Config, p project.Project) error {
	c.Load()

	if len(args) == 0 {
		return fmt.Errorf("missing project name")
	}

	name := args[0]

	if p.Exists(name) {
		return fmt.Errorf("a project with this name already exists")
	}

	fp := p.Path(name)

	if err := utils.CreateFile(fp, tpl); err != nil {
		return err
	}

	if err := cmd.Exec(c.Get("editor"), fp); err != nil {
		return p.Remove(name)
	}

	return nil
}
