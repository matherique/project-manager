package config

import (
	"fmt"
	"os"

	"github.com/matherique/project-manager/pkg/cmd"
	cnf "github.com/matherique/project-manager/pkg/config"
)

type config struct {
	c cnf.Config
}

func NewConfig(c cnf.Config) *config {
	cfg := new(config)
	cfg.c = c
	return cfg
}

func (c *config) Exec(a []string) {
	c.c.Load()
	if len(a) == 0 {
		fmt.Fprintln(os.Stdout, c.GetAll())
		return
	}

	if len(a) == 1 && a[0] == "edit" {
		cmd.Exec(c.c.Get("editor"))
		return
	}
}

func (c *config) GetAll() string {
	c.c.Load()

	return c.c.All()
}
