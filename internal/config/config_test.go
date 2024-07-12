package config

import (
	"testing"

	"github.com/henrywhitaker3/argo-zombies/internal/exclusions"
	"github.com/stretchr/testify/assert"
)

func TestItReturnsAnEmptyConfigWhenNoFileProvided(t *testing.T) {
	cfg, err := LoadConfig("bongo")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, []string{}, cfg.Exclusions.Namespaces)
	assert.Equal(t, []string{}, cfg.Exclusions.Bundles)
	assert.Equal(t, []exclusions.ExcludedGroupVersionResource{}, cfg.Exclusions.GroupVersionResources)
	assert.Equal(t, []exclusions.ExcludedMetadata{}, cfg.Exclusions.Selectors)
}

func TestItSetsGithubDashboardAsDisabledByDefault(t *testing.T) {
	cfg, err := LoadConfig("bongo")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, false, cfg.Dashboards.Github.Enabled)
}

func TestItSetsGitlabDashboardAsDisabledByDefault(t *testing.T) {
	cfg, err := LoadConfig("bongo")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, false, cfg.Dashboards.Gitlab.Enabled)
}
