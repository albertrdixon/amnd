package amnd

import (
	"bytes"
	"os"
	"time"

	"github.com/ghodss/yaml"
	"github.com/imdario/mergo"
)

var (
	defaultInterval = 24 * time.Hour
	defaultTmp      = "/tmp"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Config struct {
	Creds    *Credentials  `yaml:"plex-pass,omitempty"`
	TmpDir   string        `yaml:"tmpdir,omitempty"`
	Cmd      []string      `yaml:"restart-command"`
	Interval time.Duration `yaml:"update-interval"`
}

type Credentials struct {
	Username string
	Password string
}

func ReadConfig(path string) (*Config, error) {
	d := &Config{Interval: defaultInterval, TmpDir: defaultTmp}
	c, e := fromFile(path)
	if e != nil {
		return d, e
	}

	if e := mergo.Merge(c, d); e != nil {
		return d, e
	}
	return c, nil
}

func fromFile(p string) (*Config, error) {
	f, e := os.Open(p)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	return fromStream(f)
}

func fromStream(r Reader) (*Config, error) {
	b := new(bytes.Buffer)
	if _, e := b.ReadFrom(r); e != nil {
		return nil, e
	}

	c := new(Config)
	e := yaml.Unmarshal(b.Bytes(), c)
	return c, e
}
