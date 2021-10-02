package project

import (
	"os"
	"path"
	"strings"

	fc "github.com/matherique/project-manager/internal/file_config"
)

func Exists(c fc.FileConfig, name string) bool {
	files := All(c)

	for _, f := range files {
		if f == name {
			return true
		}
	}

	return false
}

func All(c fc.FileConfig) []string {
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

func Path(c fc.FileConfig, name string) string {
	if !Exists(c, name) {
		return ""
	}

	return path.Join(c.Get("scripts"), name)
}

func Add(c fc.FileConfig, name string) error {
	p := c.Get("projects")

	var pl []string

	if p != "" {
		pl = strings.Split(p, ";")
	}

	pl = append(pl, name)

	c.Set("projects", strings.Join(pl, ";"))

	return c.Save()
}
