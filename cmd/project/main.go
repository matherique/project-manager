package main

import (
	"fmt"
	"log"
	"os"

	"github.com/matherique/project-manager/pkg/config"
	"github.com/matherique/project-manager/internal/pkg/create"
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

	switch os.Args[1] {
	case "create":
		crt.Exec(os.Args[2:])
	default:
		fmt.Println("subcommand not found, try: create|edit|config")
	}

}
