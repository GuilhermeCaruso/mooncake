package config

import (
	"fmt"
	"io/ioutil"
	"log"
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
}

type Config struct {
	Package string
	Files   []ConfigFile
}

type ConfigFile struct {
	Original string
	New      string
}

func NewConfig(p string) Config {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalf("Something went wrong: %s", err.Error())
	}

	var mf MooncakeFile

	if err := yaml.Unmarshal(b, &mf); err != nil {
		log.Fatalf("Fail to parse file content: %s", err.Error())
	}

	config := new(Config)
	config.setFiles(mf.Mocks.Path, mf.Mocks.Output, mf.Mocks.Files)
	return *config
}

func (c *Config) setFiles(path string, newPath string, files []string) {
	for _, f := range files {
		c.Files = append(c.Files, ConfigFile{
			Original: filepath.Join(path, f),
			New:      filepath.Join(newPath, fmt.Sprintf("generated_%s", f)),
		})
	}
}
