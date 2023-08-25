package config

import (
	"testing"

	"github.com/henrywhitaker3/argo-zombies/internal/exclusions"
	"github.com/stretchr/testify/assert"
)

func TestItReturnsAnEmptyConfigWhenNoFileProvided(t *testing.T) {
	if err := LoadConfig("bongo"); err != nil {
		t.Error(err)
	}

	assert.Equal(t, []string{}, Cfg.Exclusions.Namespaces)
	assert.Equal(t, []string{}, Cfg.Exclusions.Bundles)
	assert.Equal(t, []exclusions.ExcludedGroupVersionResource{}, Cfg.Exclusions.GroupVersionResources)
	assert.Equal(t, []exclusions.ExcludedMetadata{}, Cfg.Exclusions.Selectors)
}

func TestItSetsGithubDashboardAsDisabledByDefault(t *testing.T) {
	if err := LoadConfig("bongo"); err != nil {
		t.Error(err)
	}

	assert.Equal(t, false, Cfg.Dashboards.Github.Enabled)
}

func TestItSetsGitlabDashboardAsDisabledByDefault(t *testing.T) {
	if err := LoadConfig("bongo"); err != nil {
		t.Error(err)
	}

	assert.Equal(t, false, Cfg.Dashboards.Gitlab.Enabled)
}
