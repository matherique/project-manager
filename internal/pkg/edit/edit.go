package edit

import (
	"fmt"
	"strings"

	"github.com/matherique/project-manager/pkg/cmd"
	fc "github.com/matherique/project-manager/pkg/file_config"
)

type edit struct {
	c fc.FileConfig
}

func NewEdit(c fc.FileConfig) *edit {
	e := new(edit)
	e.c = c
	return e
}

func (e *edit) Exec(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing project name")
	}

	e.c.Load()

	pl := strings.Split(e.c.Get("projects"), ";")

	if len(pl) == 0 {
		return fmt.Errorf("no project created")
	}

	var p string
	for _, v := range pl {
		if args[0] == v {
			p = v
		}
	}

	if p == "" {
		return fmt.Errorf("no project found with this name: %s", args[0])
	}

	return cmd.Exec(e.c.Get("editor"), args[0])
}
