package project

import (
	"fmt"
	"os"
	"path"

	"github.com/matherique/project-manager/pkg/config"
)

type Project interface {
	Exists(name string) bool
	All() []string
	Path(name string) string
	Remove(name string) error
}

type project struct {
	c *config.FileConfig
}

func NewProject(c config.FileConfig) *project {
	p := new(project)
	p.config = c

	return p
}

func (p project) Exists(name string) bool {
	files := All(c)

	for _, f := range files {
		if f == name {
			return true
		}
	}

	return false
}

func (p project) All() []string {
	c.Load()
	var fnames []string

	s := c.Get("scripts")

	files, err := os.ReadDir(s)

	if err != nil {
		return fnames
	}

	for _, f := range files {
		fnames = append(fnames, f.Name())
	}

	return fnames
}

func (p project) Path(name string) string {
	return path.Join(c.Get("scripts"), name)
}

func (p project) Remove(name string) error {
	c.Load()

	if !Exists(c, name) {
		return fmt.Errorf("project not found")
	}

	return os.Remove(Path(c, name))
}
