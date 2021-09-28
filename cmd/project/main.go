package main

import (
	"fmt"
	"log"
	"os"

	"github.com/matherique/project-manager/internal/pkg/config"
	"github.com/matherique/project-manager/internal/pkg/create"
	"github.com/matherique/project-manager/internal/pkg/list"
	"github.com/matherique/project-manager/internal/pkg/open"
	fc "github.com/matherique/project-manager/pkg/file_config"
)

type Teste struct {
	Editor  string
	Scripts string
}

func main() {
	c, err := fc.NewConfig("config")

	if err != nil {
		log.Fatalf("could not load config file: %v", err)
	}

	if len(os.Args) < 2 {
		fmt.Println("subcommand not found, try: create|open|edit|config")
		os.Exit(1)
	}

	err = c.Read()

	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}

	crt := create.NewCreate(c)
	op := open.NewOpen(c)
	cfg := config.NewConfig(c)
	list := list.NewList(c)

	switch os.Args[1] {
	case "create":
		err = crt.Exec(os.Args[2:])
	case "open":
		err = op.Exec(os.Args[2:])
	case "config":
		err = cfg.Exec(os.Args[2:])
	case "list":
		err = list.Exec(os.Args[2:])
	default:
		fmt.Println("subcommand not found, try: create|open|edit|config")
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
