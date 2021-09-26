package create

import (
	"fmt"
	"html/template"
	"os"
)

func createFile(file string) error {
	f, err := os.Create(file)

	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}

	t := template.Must(template.New("project").Parse(tpl))

	err = t.Execute(f, f.Name())

	if err != nil {
		return fmt.Errorf("could not save template file: %v", err)
	}

	err = os.Chmod(f.Name(), 0777)

	if err != nil {
		return fmt.Errorf("could not save template file: %v", err)
	}

	return nil
}
