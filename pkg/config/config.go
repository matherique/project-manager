package config

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

// Parse config file and get all the config data
//
// Possible configs:
//   - editor: editor used to open the script file and edit it
//   - scripts: folder to store all scripts

type config struct {
  config map[string]string
}

func (c *config) Get(key string) (string, error) {
  v, ok := c.config[key]
  
  if !ok {
    return "", fmt.Errorf("key not found")
  }

  return v, nil
}

// Receive io.Reader and return a new config instance
func New(r io.Reader) *config {
  s := bufio.NewScanner(r)

  d := make(map[string]string)

  for s.Scan() {
    re := regexp.MustCompile(`^(.*)=(.*)$`)
    res := re.FindStringSubmatch(s.Text())
    if len(res) == 0 || len(res) < 3 {
      continue
    }

    d[res[1]] = res[2] 
  }
  
  return &config{ config: d }
}
