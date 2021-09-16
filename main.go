package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

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
}

func create(args []string) {
	if len(args) == 0 {
		fmt.Println("missing project name")
		os.Exit(0)
	}

	c := []byte("#!/bin/bash\n\n")
	n := args[0]

	err := os.WriteFile(n, c, 0644)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	p, _ := exec.LookPath("vi")

	cmd := &exec.Cmd{
		Path:   p,
		Args:   []string{p, n},
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
