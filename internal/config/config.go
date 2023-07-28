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

	Cfg.addBundles()

	return nil
}

func (c *Config) setDefaults() {
	if structs.IsZero(c.Exclusions) {
		c.Exclusions = Exclusions{
			Resources:             []ExcludedResource{},
			Namespaces:            []string{},
			Selectors:             []ExcludedMetadata{},
			GroupVersionResources: []ExcludedGroupVersionResource{},
			Bundles:               []string{},
		}
	}
}

func (c *Config) addBundles() {
	c.Exclusions.Bundles = append(c.Exclusions.Bundles, "k8s")
	for _, bundle := range c.Exclusions.Bundles {
		if f, valid := bundles[bundle]; valid {
			excls := f()

			c.mergeResources(excls.Resources)
			c.mergeNamespaces(excls.Namespaces)
			c.mergeSelectors(excls.Selectors)
			c.mergeGroupVersionResources(excls.GroupVersionResources)
		}
	}
}

func (c *Config) mergeResources(r []ExcludedResource) {
	c.Exclusions.Resources = append(c.Exclusions.Resources, r...)
}

func (c *Config) mergeNamespaces(ns []string) {
	c.Exclusions.Namespaces = append(c.Exclusions.Namespaces, ns...)
}

func (c *Config) mergeSelectors(s []ExcludedMetadata) {
	c.Exclusions.Selectors = append(c.Exclusions.Selectors, s...)
}

func (c *Config) mergeGroupVersionResources(g []ExcludedGroupVersionResource) {
	c.Exclusions.GroupVersionResources = append(c.Exclusions.GroupVersionResources, g...)
}
