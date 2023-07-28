package config

import (
	"errors"
	"os"

	"github.com/fatih/structs"
	"github.com/henrywhitaker3/argo-zombies/internal/dashboard"
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions"
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"gopkg.in/yaml.v3"
)

var (
	Cfg *Config
)

type Config struct {
	Exclusions exclusions.Exclusions `yaml:"exclusions"`
	Dashboards struct {
		Github dashboard.GithubDashboard `yaml:"github"`
	} `yaml:"dashboard"`
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
		c.Exclusions = exclusions.Exclusions{
			Resources:             []exclusions.ExcludedResource{},
			Namespaces:            []string{},
			Selectors:             []exclusions.ExcludedMetadata{},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{},
			Bundles:               []string{},
		}
	}
	if structs.IsZero(c.Dashboards) {
		c.Dashboards = struct {
			Github dashboard.GithubDashboard "yaml:\"github\""
		}{
			Github: dashboard.GithubDashboard{},
		}
	}
}

func (c *Config) addBundles() {
	c.Exclusions.Bundles = append(c.Exclusions.Bundles, "k8s")
	for _, bundle := range c.Exclusions.Bundles {
		if f, valid := bundles.Bundles[bundle]; valid {
			excls := f()

			c.mergeResources(excls.Resources)
			c.mergeNamespaces(excls.Namespaces)
			c.mergeSelectors(excls.Selectors)
			c.mergeGroupVersionResources(excls.GroupVersionResources)
		}
	}
}

func (c *Config) mergeResources(r []exclusions.ExcludedResource) {
	c.Exclusions.Resources = append(c.Exclusions.Resources, r...)
}

func (c *Config) mergeNamespaces(ns []string) {
	c.Exclusions.Namespaces = append(c.Exclusions.Namespaces, ns...)
}

func (c *Config) mergeSelectors(s []exclusions.ExcludedMetadata) {
	c.Exclusions.Selectors = append(c.Exclusions.Selectors, s...)
}

func (c *Config) mergeGroupVersionResources(g []exclusions.ExcludedGroupVersionResource) {
	c.Exclusions.GroupVersionResources = append(c.Exclusions.GroupVersionResources, g...)
}
