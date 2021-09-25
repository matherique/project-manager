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

func (c *config) Exec(a []string) {
	c.c.Load()
	if len(a) == 0 {
		fmt.Fprintln(os.Stdout, c.GetAll())
		return
	}

	if len(a) == 1 {
		if a[0] == "edit" {
		}

		if c.c.HasKey(a[0]) {
			fmt.Println(c.c.Get(a[0]))
			return
		}
	}

}

func (c *config) GetAll() string {
	c.c.Load()

	return c.c.All()
}

func (c *config) Edit() {

	cf := c.c.FilePath()

	if cf == "" {
		fmt.Println("config file not found")
		os.Exit(1)
	}

	cmd.Exec(c.c.Get("editor"), cf)
}
