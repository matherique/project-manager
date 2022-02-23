package main

import (
	"fmt"
	"os"

	"github.com/matherique/project-manager/pkg/cmd"
	fc "github.com/matherique/project-manager/pkg/file_config"
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

func cmd_config(a []string, c fc.FileConfig) error {
	c.Load()

	if len(a) == 0 {
		fmt.Fprintln(os.Stdout, GetAll(c))
		return nil
	}

	switch len(a) {
	case 1:
		if a[0] == "edit" {
			return Edit(c)
		}
		return GetKey(c, a[0])
	case 2:
		return SetValue(c, a[0], a[1])
	}

	return nil
}

func GetAll(c fc.FileConfig) string {
	return c.All()
}

func SetValue(c fc.FileConfig, key, value string) error {
	if !c.HasKey(key) {
		return fmt.Errorf("key not found")
	}

	c.Set(key, value)

	return c.Save()
}

func GetKey(c fc.FileConfig, key string) error {
	if !c.HasKey(key) {
		return fmt.Errorf("key not found")
	}

	fmt.Fprintln(os.Stdout, c.Get(key))

	return nil
}

func Edit(c fc.FileConfig) error {
	cf := c.Path()

	if cf == "" {
		return fmt.Errorf("config file not found")
	}

	return cmd.Exec(c.Get("editor"), cf)
}
