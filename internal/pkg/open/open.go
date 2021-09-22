package open

import (
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/matherique/project-manager/pkg/config"
)

type open struct {
  c config.Config
}

func NewOpen(c config.Config) *open {
  o := new(open)
  o.c = c
  return o
}

func (o *open) Exec(a []string) {
  if len(a) == 0 {
    log.Fatalf("missing project name")
  }
  
  o.c.Load()
  sp := o.c.Get("scripts")
  fp := path.Join(sp, a[0])

	p, err := exec.LookPath(fp)

	if err != nil {
    log.Fatalf("file not found: %v", err)
	}

	cmd := &exec.Cmd{
		Path:   p,
		Args:   []string{p, fp},
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

  err = cmd.Run()

	if err != nil {
    log.Fatalf("could not run the command: %v", err)
	}
}
