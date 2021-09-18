package create

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
)

const tpl = `#!/bin/bash

project={{.}}

tmux new-session -s $project
`

func Execute(args []string) {
	if len(args) == 0 {
		fmt.Println("missing project name")
		os.Exit(0)
	}

	fn := args[0]
  f, err := os.Create(fn)

	if err != nil {
		log.Fatal(err)
	}

  t := template.Must(template.New("project").Parse(tpl))
  
  err = t.Execute(f, fn)

	if err != nil {
		log.Fatal(err)
	}

  err = os.Chmod(fn, 0777)

	if err != nil {
		log.Fatal(err)
	}

	p, _ := exec.LookPath("vi")

	cmd := &exec.Cmd{
		Path:   p,
		Args:   []string{p, fn},
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

  err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
