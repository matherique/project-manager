package config

import (
	"io"
)

type Getter interface {
  Get(key string) string
}

type Setter interface {
  Set(key, value string) error
}

type Saver interface {
  Save() error
}

type Reader interface {
  Read(r io.Reader) error 
}

type Loader interface {
  Load() error 
}

type Config interface {
  Getter
  Setter
  Saver
  Reader
  Loader
}

type config struct {
  m map[string]string
}

func NewConfig(r io.Reader) (*config, error) {
  m := make(map[string]string)
  c := new(config)
  c.m = m

  err := c.Read(r)

  if err != nil {
    return nil, err
  }

  return c, nil
}

func (c *config) Get(key string) string {
  v, ok := c.m[key]

  if !ok {
    return ""
  }

  return v
}

func (c *config) Set(key, value string) error { return nil }
func (c *config) Save() error { return nil }
func (c *config) Read(r io.Reader) error { return nil }
func (c *config) Load() error  { return nil }

/*
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
  
  return &config{ m: d }
}
*/
