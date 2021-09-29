package utils

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

func CreateFile(file, tpl string) error {
	f, err := os.Create(file)

	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}

	name := path.Base(file)

	t := template.Must(template.New("project").Parse(tpl))

	err = t.Execute(f, name)

	if err != nil {
		return fmt.Errorf("could not save template file: %v", err)
	}

	err = os.Chmod(name, 0777)

	if err != nil {
		return fmt.Errorf("could not save template file: %v", err)
	}

	return nil
}
