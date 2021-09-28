package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
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

type FilePather interface {
	FilePath() string
}

type FileConfig interface {
	Getter
	Setter
	Saver
	Loader
	Reader
	FilePather
	parse(r io.Reader)
	All() string
	HasKey(k string) bool
	Keys() []string
	Values() []string
	Raw() []string
}

type fileConfig struct {
	f string
	m map[string]string
}

func NewConfig(f string) (*fileConfig, error) {
	m := make(map[string]string)
	c := new(fileConfig)
	c.m = m
	c.f = f

	err := c.Read()

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *fileConfig) HasKey(key string) bool {
	for _, v := range c.Keys() {
		if v == key {
			return true
		}
	}

	return false
}

func (c *fileConfig) Keys() []string {
	keys := make([]string, len(c.m))

	i := 0
	for k := range c.m {
		keys[i] = k
		i += 1
	}

	return keys
}

func (c *fileConfig) Values() []string {
	values := make([]string, len(c.f))

	i := 0
	for _, v := range c.m {
		values[i] = v
		i += 1
	}

	return values
}

func (c *fileConfig) Get(key string) string {
	v, ok := c.m[key]

	if !ok {
		return ""
	}

	return strings.Trim(v, "\"")
}

func (c *fileConfig) Raw() []string {
	r := make([]string, len(c.m))

	for i, k := range c.Keys() {
		r[i] = fmt.Sprintf("%s=%s", k, c.m[k])
	}

	sort.Strings(r)
	return r
}

func (c *fileConfig) Set(key, value string) { c.m[key] = value }

func (c *fileConfig) Save() error {
	err := os.Truncate(c.FilePath(), 0)

	if err != nil {
		return err
	}

	f, err := os.OpenFile(c.FilePath(), os.O_WRONLY, 0755)

	if err != nil {
		return err
	}

	defer f.Close()

	for _, r := range c.Raw() {
		f.Write([]byte(fmt.Sprintln(r)))
	}

	return nil
}

func (c *fileConfig) Load() {
	f := c.ConfigFile()
	r, _ := os.Open(f)

	defer r.Close()

	c.parse(r)
}

func (c *fileConfig) FilePath() string {
	fp, err := filepath.Abs(c.f)

	if err != nil {
		return ""
	}

	return fp
}

func (c *fileConfig) Read() error {
	f := c.ConfigFile()

	r, err := os.Open(f)

	if err != nil {
		return fmt.Errorf("could not open file %v", err)
	}

	defer r.Close()

	c.parse(r)

	return nil
}

func (c *fileConfig) parse(r io.Reader) {
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

func (c *fileConfig) All() string { return strings.Join(c.Raw(), "\n") }

func (c *fileConfig) ConfigFile() string {
	if c.f == "" {
		// TODO: get config path
		return "config"
	}

	return c.f
}
