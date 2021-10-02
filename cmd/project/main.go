package main

import (
	"fmt"
	"log"
	"os"

	fc "github.com/matherique/project-manager/internal/file_config"
)

func main() {
	c, err := fc.NewConfig()

	if err != nil {
		log.Fatalf("could not load config file: %v", err)
	}

	err = c.Read()

	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}

	var cmd func([]string, fc.FileConfig) error

	switch os.Args[1] {
	case "create":
		cmd = cmd_create
	case "open":
		cmd = cmd_open
	case "config":
		cmd = cmd_config
	case "list":
		cmd = cmd_list
	case "edit":
		cmd = cmd_edit
	default:
		cmd = nil
	}

	if cmd == nil {
		fmt.Fprintln(os.Stdout, "subcommand not found, try: create|open|edit|config")
		os.Exit(1)
	}

	err = cmd(os.Args[2:], c)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
