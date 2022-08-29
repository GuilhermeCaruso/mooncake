package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type MooncakeFile struct {
	Mocks Mock `yaml:"mocks"`
}

type Mock struct {
	Package string   `yaml:"package"`
	Path    string   `yaml:"path"`
	Files   []string `yaml:"files"`
	Output  string   `yaml:"output"`
	Prefix  string   `yaml:"prefix"`
}

type Config struct {
	Package string
	Prefix  string
	Output  string
	Files   []ConfigFile
}

type ConfigFile struct {
	Original string
	New      string
}

func NewConfig(p string) Config {
	dir := filepath.Dir(p)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalf("Something went wrong: %s", err.Error())
	}

	var mf MooncakeFile

	if err := yaml.Unmarshal(b, &mf); err != nil {
		log.Fatalf("Fail to parse file content: %s", err.Error())
	}

	config := new(Config)
	config.Package = mf.Mocks.Package
	config.Output = filepath.Join(dir, mf.Mocks.Output)
	config.Prefix = "generated"
	if mf.Mocks.Prefix != "" {
		config.Prefix = mf.Mocks.Prefix
	}
	config.setFiles(filepath.Join(dir, mf.Mocks.Path), filepath.Join(dir, mf.Mocks.Output), mf.Mocks.Files)
	return *config
}

func (c *Config) PrepareFolder() {
	if _, err := os.Stat(c.Output); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(c.Output, os.ModePerm); err != nil {
			log.Fatalf("Fail to create output folder: %s", err.Error())

		}

	}
}

func (c *Config) setFiles(path string, newPath string, files []string) {

	for _, f := range files {
		c.Files = append(c.Files, ConfigFile{
			Original: filepath.Join(path, f),
			New:      filepath.Join(newPath, fmt.Sprintf("%s_%s", c.Prefix, f)),
		})
	}
}
