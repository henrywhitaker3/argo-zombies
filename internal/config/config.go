package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	Cfg *Config
)

type Config struct {
	Exclusions Exclusions `yaml:"exclusions"`
}

func LoadConfig(path string) error {
	c := &Config{}

	file, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if err := yaml.Unmarshal(file, c); err != nil {
		return err
	}

	Cfg = c

	return nil
}
