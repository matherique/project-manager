package project

import (
	"os"

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
