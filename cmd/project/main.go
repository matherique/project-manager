package main

import (
	"fmt"
	"log"
	"os"

	fc "github.com/matherique/project-manager/internal/file_config"

	command "github.com/matherique/cmd"
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

	init := command.New("project")
	init.SetLongDesc(doc_help)
	init.SetHandler(func(args []string) error {
		return cmd_fzf(args, c)
	})

	create := command.New("create", "new")
	create.SetLongDesc(doc_create)
	create.SetHandler(func(args []string) error {
		return cmd_create(args, c)
	})

	open := command.New("open")
	open.SetLongDesc(doc_open)
	open.SetHandler(func(args []string) error {
		return cmd_open(args, c)
	})

	config := command.New("config")
	config.SetLongDesc(doc_config)
	config.SetHandler(func(args []string) error {
		return cmd_config(args, c)
	})

	list := command.New("list", "ls")
	list.SetLongDesc(doc_list)
	list.SetHandler(func(args []string) error {
		return cmd_list(args, c)
	})

	remove := command.New("remove", "rm")
	remove.SetLongDesc(doc_remove)
	remove.SetHandler(func(args []string) error {
		return cmd_remove(args, c)
	})

	fzf := command.New("fzf")
	fzf.SetLongDesc(doc_fzf)
	fzf.SetHandler(func(args []string) error {
		return cmd_fzf(args, c)
	})

	stop := command.New("stop")
	stop.SetLongDesc(doc_stop)
	stop.SetHandler(func(args []string) error {
		return cmd_stop(args, c)
	})

	init.AddSub(create)
	init.AddSub(open)
	init.AddSub(config)
	init.AddSub(list)
	init.AddSub(remove)
	init.AddSub(fzf)

	if err := init.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
