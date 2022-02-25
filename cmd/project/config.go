package main

import (
	"fmt"

	"github.com/matherique/project-manager/pkg/cmd"
	"github.com/matherique/project-manager/pkg/config"
)

const doc_config string = `
Usage:
  project config                     Get all configuration
  project config [name]              Get the configuration by name
  project config [name] [value]      Set the configuration value

Configurations avaliable:
  editor     Editor binary name
  scripts    Path to folder used to save projects

Set or get configuration value
`

func cmd_config(args []string, c config.Config) (string, error) {
	c.Load()

	argsize := len(args)

	// project config
	if argsize == 0 {
		return c.All(), nil
	}

	// project config edit
	if args[0] == "edit" {
		return "", edit(c)
	}

	// project config invalid_key
	if !c.HasKey(args[0]) {
		return "", fmt.Errorf("key not found")
	}

	switch argsize {
	case 1:
		// project config key
		return c.Get(args[0]), nil
	case 2:
		// project config key value
		c.Set(args[0], args[1])
		return "", c.Save()
	}

	return "", nil
}

func edit(c config.Config) error {
	file := c.Path()

	if file == "" {
		return fmt.Errorf("config file not found")
	}

	return cmd.Exec(c.Get("editor"), file)
}
