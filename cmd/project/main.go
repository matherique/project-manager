package main

import (
	"fmt"
	"log"
	"os"

	"github.com/matherique/project-manager/internal/pkg/create"
	"github.com/matherique/project-manager/internal/pkg/open"
	"github.com/matherique/project-manager/pkg/config"
)

type Teste struct {
  Editor string
  Scripts string
}

func main() {
  c, err := config.NewConfig("config")

  if err != nil {
    log.Fatalf("could not load config file: %v", err)
  }

	if len(os.Args) < 2 {
		fmt.Println("subcommands: create")
		return
	}

  crt := create.NewCreate(c)
  op := open.NewOpen(c)

	switch os.Args[1] {
	case "create":
		crt.Exec(os.Args[2:])
	case "open":
	  op.Exec(os.Args[2:])
	default:
		fmt.Println("subcommand not found, try: create|open|edit|config")
	}

}
