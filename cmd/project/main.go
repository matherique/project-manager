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
	var doc string

	switch os.Args[1] {
	case "help":
		cmd = cmd_help
	case "create":
		cmd = cmd_create
		doc = doc_create
	case "open":
		cmd = cmd_open
	case "new":
		cmd = cmd_create
		doc = doc_create
	case "config":
		cmd = cmd_config
	case "list":
		cmd = cmd_list
	case "ls":
		cmd = cmd_list
	case "edit":
		cmd = cmd_edit
	case "remove":
		cmd = cmd_remove
	case "rm":
		cmd = cmd_remove
	case "fzf":
		cmd = cmd_fzf
	default:
		cmd = nil
	}

	if cmd == nil {
		fmt.Fprintln(os.Stdout, "subcommand not found, try: create|open|edit|config|remove")
		os.Exit(1)
	}

	if len(os.Args) > 1 && os.Args[2] == "help" {
		fmt.Fprint(os.Stdout, doc)
		return
	}

	err = cmd(os.Args[2:], c)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
