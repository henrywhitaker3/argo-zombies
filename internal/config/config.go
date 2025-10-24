package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/structs"
	"github.com/henrywhitaker3/argo-zombies/internal/dashboard"
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions"
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"gopkg.in/yaml.v3"
)

var (
	Cfg *Config
)

type Dashboards struct {
	Title  string                    `yaml:"title"`
	Github dashboard.GithubDashboard `yaml:"github"`
	Gitlab dashboard.GitlabDashboard `yaml:"gitlab"`
}

type Config struct {
	Exclusions exclusions.Exclusions `yaml:"exclusions"`
	Dashboards Dashboards            `yaml:"dashboards"`
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
	Cfg.loadEnvVars()

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
		c.Dashboards = Dashboards{
			Github: dashboard.GithubDashboard{Enabled: false},
			Gitlab: dashboard.GitlabDashboard{Enabled: false},
		}
	}
	if c.Dashboards.Title == "" {
		c.Dashboards.Title = "Argo Zombies Dashboards"
	}
}

func (c *Config) loadEnvVars() {
	if os.Getenv("GITHUB_TOKEN") != "" {
		c.Dashboards.Github.Token = os.Getenv("GITHUB_TOKEN")
	}
	if id := os.Getenv("GITHUB_APP_CLIENT_ID"); id != "" {
		c.Dashboards.Github.ClientID = id
	}
	if id := os.Getenv("GITHUB_APP_INSTALLATION_ID"); id != "" && id != "0" {
		instID, err := strconv.Atoi(id)
		if err != nil {
			panic(fmt.Errorf("convert installation id to int: %w", err))
		}
		c.Dashboards.Github.InstallationID = instID
	}
	if priv := os.Getenv("GITHUB_APP_PRIVATE_KEY"); priv != "" {
		c.Dashboards.Github.PrivateKey = priv
	}
	if os.Getenv("GITLAB_TOKEN") != "" {
		c.Dashboards.Gitlab.Token = os.Getenv("GITLAB_TOKEN")
	}
	if title := os.Getenv("DASHBOARD_TITLE"); title != "" {
		c.Dashboards.Title = title
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
