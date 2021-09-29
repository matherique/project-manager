package main

import (
	"fmt"
	"log"
	"os"

	fc "github.com/matherique/project-manager/pkg/file_config"
)

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

	var cmd func([]string, fc.FileConfig) error

	switch os.Args[1] {
	case "create":
		cmd = cmd_create
	// case "open":
	// 	cmd = op
	// case "config":
	// 	cmd = cfg
	// case "list":
	// 	cmd = list
	// case "edit":
	// 	cmd = edit
	default:
		cmd = nil
	}

	if cmd == nil {
		fmt.Println("subcommand not found, try: create|open|edit|config")
		os.Exit(1)
	}

	err = cmd(os.Args[2:], c)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
