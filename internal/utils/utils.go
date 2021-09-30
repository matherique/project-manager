package utils

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

func CreateFile(file string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}

	name := path.Base(file)

	t := template.Must(template.New("project").ParseFiles("./template/project.tmpl"))

	return t.Execute(f, name)
}
