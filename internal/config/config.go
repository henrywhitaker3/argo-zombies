package config

import (
	"errors"
	"os"

	"github.com/fatih/structs"
	"github.com/henrywhitaker3/argo-zombies/internal/dashboard"
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions"
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

type LogLevel string

func (l LogLevel) Level() zap.AtomicLevel {
	switch l {
	case "info":
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "debug":
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "error":
		fallthrough
	default:
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	}
}

type Config struct {
	LogLevel   LogLevel              `yaml:"logLevel"`
	Exclusions exclusions.Exclusions `yaml:"exclusions"`
	Dashboards struct {
		Github dashboard.GithubDashboard `yaml:"github"`
		Gitlab dashboard.GitlabDashboard `yaml:"gitlab"`
	} `yaml:"dashboards"`
}

func LoadConfig(path string) (*Config, error) {
	c := &Config{}
	c.setDefaults()

	file, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return c, nil
		}

		return nil, err
	}

	if err := yaml.Unmarshal(file, c); err != nil {
		return nil, err
	}

	c.addBundles()
	c.loadEnvVars()

	return c, nil
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
			Gitlab dashboard.GitlabDashboard "yaml:\"gitlab\""
		}{
			Github: dashboard.GithubDashboard{Enabled: false},
			Gitlab: dashboard.GitlabDashboard{Enabled: false},
		}
	}
}

func (c *Config) loadEnvVars() {
	if os.Getenv("GITHUB_TOKEN") != "" {
		c.Dashboards.Github.Token = os.Getenv("GITHUB_TOKEN")
	}
	if os.Getenv("GITLAB_TOKEN") != "" {
		c.Dashboards.Gitlab.Token = os.Getenv("GITLAB_TOKEN")
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
