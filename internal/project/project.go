package project

import (
	"fmt"
	"os"
	"path"

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

func Remove(c fc.FileConfig, name string) error {
	c.Load()

	if !Exists(c, name) {
		return fmt.Errorf("project not found")
	}

	err := os.Remove(Path(c, name))

	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "project removed with successfully")

	return nil
}
