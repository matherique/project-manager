package main

import (
	"fmt"
	"log"
	"os"

	"github.com/matherique/project-manager/internal/project"
	fc "github.com/matherique/project-manager/pkg/config"

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

	p := project.NewProject(c)

	init := command.New("project")
	init.SetLongDesc(doc_help)
	init.SetHandler(func(args []string) error {
		return cmd_fzf(args, c, p)
	})

	create := command.New("create", "new")
	create.SetLongDesc(doc_create)
	create.SetHandler(func(args []string) error {
		return cmd_create(args, c, p)
	})

	open := command.New("open")
	open.SetLongDesc(doc_open)
	open.SetHandler(func(args []string) error {
		return cmd_open(args, c, p)
	})

	edit := command.New("edit")
	edit.SetLongDesc(doc_edit)
	edit.SetHandler(func(args []string) error {
		return cmd_edit(args, c, p)
	})

	config := command.New("config")
	config.SetLongDesc(doc_config)
	config.SetHandler(func(args []string) error {
		cnf, err := cmd_config(args, c)

		if err != nil {
			return err
		}

		fmt.Fprint(os.Stdout, cnf)
		return nil
	})

	list := command.New("list", "ls")
	list.SetLongDesc(doc_list)
	list.SetHandler(func(args []string) error {
		l, err := cmd_list(args, c, p)

		if err != nil {
			return err
		}

		fmt.Fprint(os.Stdout, l)
		return nil
	})

	remove := command.New("remove", "rm")
	remove.SetLongDesc(doc_remove)
	remove.SetHandler(func(args []string) error {
		r, err := cmd_remove(args, c, p)

		if err != nil {
			return err
		}

		fmt.Fprint(os.Stdout, r)
		return nil
	})

	fzf := command.New("fzf")
	fzf.SetLongDesc(doc_fzf)
	fzf.SetHandler(func(args []string) error {
		return cmd_fzf(args, c, p)
	})

	stop := command.New("stop")
	stop.SetLongDesc(doc_stop)
	stop.SetHandler(func(args []string) error {
		return cmd_stop(args, c, p)
	})

	init.AddSub(create)
	init.AddSub(open)
	init.AddSub(config)
	init.AddSub(list)
	init.AddSub(remove)
	init.AddSub(fzf)
	init.AddSub(stop)
	init.AddSub(edit)

	if err := init.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
