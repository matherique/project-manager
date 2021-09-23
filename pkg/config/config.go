package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Getter interface {
	Get(key string) string
}

type Setter interface {
	Set(key, value string)
}

type Saver interface {
	Save() error
}

type Loader interface {
	Load()
}

type Reader interface {
	Read() error
}

type Config interface {
	Getter
	Setter
	Saver
	Loader
	Reader
	parse(r io.Reader)
	All() string
}

type config struct {
	f string
	m map[string]string
}

func NewConfig(f string) (*config, error) {
	m := make(map[string]string)
	c := new(config)
	c.m = m
	c.f = f

	err := c.Read()

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

	return strings.Trim(v, "\"")
}

func (c *config) Set(key, value string) {
	c.m[key] = value
}

func (c *config) Save() error {
	return nil
}

func (c *config) Load() {
	f := c.ConfigFile()
	r, _ := os.Open(f)

	defer r.Close()

	c.parse(r)
}

func (c *config) Read() error {
	f := c.ConfigFile()

	r, err := os.Open(f)

	if err != nil {
		return fmt.Errorf("could not open file %v", err)
	}

	defer r.Close()

	c.parse(r)

	return nil
}

func (c *config) parse(r io.Reader) {
	s := bufio.NewScanner(r)

	c.m = make(map[string]string)

	for s.Scan() {
		re := regexp.MustCompile(`^(.*)=(.*)$`)
		res := re.FindStringSubmatch(s.Text())

		if len(res) == 0 || len(res) < 3 {
			continue
		}

		c.m[res[1]] = res[2]
	}

}

func (c *config) All() string {
	var all []string

	for k, v := range c.m {
		all = append(all, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(all, "\n")
}

func (c *config) ConfigFile() string {
	if c.f == "" {
		// TODO: get config path
		return "config"
	}

	return c.f
}
