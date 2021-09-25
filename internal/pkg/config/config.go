package config

import (
	"fmt"
	"os"

	"github.com/matherique/project-manager/pkg/cmd"
	fc "github.com/matherique/project-manager/pkg/file_config"
)

type config struct {
	c fc.FileConfig
}

func NewConfig(c fc.FileConfig) *config {
	cfg := new(config)
	cfg.c = c
	return cfg
}

func (c *config) Exec(a []string) error {
	c.c.Load()

	if len(a) == 0 {
		fmt.Fprintln(os.Stdout, c.GetAll())
		return nil
	}

	switch len(a) {
	case 1:
		if a[0] == "edit" {
			return c.Edit()
		}
		return c.GetKey(a[0])
	case 2:
		return c.SetValue(a[0], a[1])
	}

	return nil
}

func (c *config) GetAll() string {
	return c.c.All()
}

func (c *config) SetValue(key, value string) error {
	if !c.c.HasKey(key) {
		return fmt.Errorf("key not found")
	}

	c.c.Set(key, value)

	return c.c.Save()
}

func (c *config) GetKey(key string) error {
	if !c.c.HasKey(key) {
		return fmt.Errorf("key not found")
	}

	fmt.Println(c.c.Get(key))

	return nil
}

func (c *config) Edit() error {
	cf := c.c.FilePath()

	if cf == "" {
		return fmt.Errorf("config file not found")
	}

	cmd.Exec(c.c.Get("editor"), cf)
	return nil
}
