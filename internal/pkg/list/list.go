package list

import (
	"fmt"
	"os"
	"strings"

	fc "github.com/matherique/project-manager/pkg/file_config"
)

type list struct {
	c fc.FileConfig
}

func NewList(c fc.FileConfig) *list {
	ct := new(list)
	ct.c = c
	return ct
}

func (c *list) Exec(a []string) error {
	c.c.Load()

	p := strings.Split(c.c.Get("projects"), ";")

	if len(p) == 0 {
		fmt.Fprintln(os.Stdout, "no project found")
	}

	fmt.Fprintln(os.Stdout, strings.Join(p, "\n"))

	return nil
}
