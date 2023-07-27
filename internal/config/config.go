package config

import (
	"errors"
	"os"

	"github.com/fatih/structs"
	"gopkg.in/yaml.v3"
)

var (
	Cfg *Config
)

type Config struct {
	Exclusions Exclusions `yaml:"exclusions"`
}

func LoadConfig(path string) error {
	Cfg = &Config{}
	Cfg.setDefaults()

	file, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if err := yaml.Unmarshal(file, Cfg); err != nil {
		return err
	}

	return nil
}

func (c *Config) setDefaults() {
	if structs.IsZero(c.Exclusions) {
		c.Exclusions = Exclusions{
			Resources:             []ExcludedResource{},
			Namespaces:            []string{},
			Selectors:             []ExcludedMetadata{},
			GroupVersionResources: []ExcludedGroupVersionResource{},
		}
	}
}
