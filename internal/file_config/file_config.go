package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"text/template"
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

type FileConfig interface {
	Getter
	Setter
	Saver
	Loader
	Reader
	parse(r io.Reader)
	All() string
	HasKey(k string) bool
	Keys() []string
	Values() []string
	Raw() []string
	HasConfigFile() bool
	Path() string
	Home() string
	Create() error
	Default() (defaultConfig, error)
	Template() string
}

type fileConfig struct {
	h string
	f string
	m map[string]string
}

type defaultConfig struct {
	Editor   string
	Scripts  string
	Projects string
}

func NewConfig() (*fileConfig, error) {
	m := make(map[string]string)
	c := new(fileConfig)
	c.m = m
	c.f = "config"

	h, err := os.UserConfigDir()

	if err != nil {
		return nil, err
	}

	exe, _ := os.Executable()
	c.h = path.Join(h, path.Base(exe))

	if !c.HasConfigFile() {
		err = c.Create()

		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *fileConfig) HasConfigFile() bool {
	f, _ := os.Stat(c.Path())
	return f != nil
}

func (c *fileConfig) Path() string { return path.Join(c.h, c.f) }

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
	err := os.Truncate(c.Path(), 0)

	if err != nil {
		return err
	}

	f, err := os.OpenFile(c.Path(), os.O_WRONLY|os.O_TRUNC, 0755)

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
	r, _ := os.Open(c.Path())

	defer r.Close()

	c.parse(r)
}

func (c *fileConfig) Read() error {
	r, err := os.Open(c.Path())

	if err != nil {
		return err
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

func (c *fileConfig) Home() string { return c.h }

func (c *fileConfig) Create() error {
	defaults, err := c.Default()

	if err != nil {
		return err
	}

	f, _ := os.OpenFile(c.Path(), os.O_CREATE|os.O_WRONLY, 0666)

	t := template.Must(template.New("project").Parse(c.Template()))

	return t.Execute(f, defaults)
}

func (c *fileConfig) Default() (defaultConfig, error) {
	e := os.Getenv("EDITOR")
	s := path.Join(c.Home(), "scripts")

	if err := os.MkdirAll(s, 0700); err != nil {
		return defaultConfig{}, err
	}

	return defaultConfig{
		Editor:   e,
		Scripts:  s,
		Projects: "",
	}, nil
}

func (c *fileConfig) Template() string {
	return `editor={{.Editor}}
projects={{.Projects}}
scripts={{.Scripts}}
`
}
