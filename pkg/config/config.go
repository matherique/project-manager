package config

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

// Parse config file and get all the config data
//
// Possible configs:
//   - editor: editor used to open the script file and edit it
//   - scripts: folder to store all scripts
//   - updated: datetime from last update

type config struct {
  editor string
  scripts string
  updated time.Time
}

func Get(filepath string) *config {
  f, err := os.Open(filepath)

  if err != nil {
    log.Fatalf("erro while open config file: %v", err)
  }
  
  defer f.Close()

  s := bufio.NewScanner(f)

  d := make(map[string]string)

  for s.Scan() {
    line := strings.Split(s.Text(), "=")
    k := line[0]
    d[k] = line[1]
  }
  
  c := config{}

  if _, ok := d["editor"]; ok {
    c.editor = d["editor"]
  }

  if _, ok := d["scripts"]; ok {
    c.scripts = d["scripts"]
  }

  if _, ok := d["updated"]; ok {
    time, err := time.Parse(time.RFC3339, d["updated"]);

    if err != nil {
      log.Fatalf("could not parse 'updated' time %v", err)
      return &c
    }

    c.updated = time
  }

  return &c
}
