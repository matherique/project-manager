package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Teste struct {
  Editor string
  Scripts string
}

func main() {
  f, err := os.Open("config")

  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()

  /*
	if len(os.Args) < 2 {
		fmt.Println("subcommands: create")
		return
	}

	switch os.Args[1] {
	case "create":
		create(os.Args[2:])
	default:
		fmt.Println("subcommand not found, try: create|edit|config")
	}

  */
}
