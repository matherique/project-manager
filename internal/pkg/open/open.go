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


}
