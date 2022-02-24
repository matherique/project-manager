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
	c config.Config
}

func NewProject(c config.Config) *project {
	p := new(project)
	p.c = c

	return p
}

func (p project) Exists(name string) bool {
	for _, f := range p.All() {
		if f == name {
			return true
		}
	}

	return false
}

func (p project) All() []string {
	p.c.Load()
	var fnames []string

	s := p.c.Get("scripts")

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
	return path.Join(p.c.Get("scripts"), name)
}

func (p project) Remove(name string) error {
	p.c.Load()

	if !p.Exists(name) {
		return fmt.Errorf("project not found")
	}

	return os.Remove(p.Path(name))
}
