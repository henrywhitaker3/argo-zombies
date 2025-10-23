package app

import (
	"context"

	"github.com/henrywhitaker3/argo-zombies/internal/config"
	"github.com/henrywhitaker3/argo-zombies/internal/dashboard"
	"github.com/henrywhitaker3/argo-zombies/internal/detector"
)

type App struct {
	Version string
	Config  *config.Config

	Detector *detector.Detector
	Trackers []Tracker
}

func New(cfg *config.Config) (*App, error) {
	app := &App{Config: cfg}
	if cfg.Dashboards.Github.Enabled {
		app.Trackers = append(app.Trackers, dashboard.NewGithub(cfg.Dashboards.Github.Repo, cfg.Dashboards.Github.Token))
	}
	if cfg.Dashboards.Gitlab.Enabled {
		t, err := dashboard.NewGitlab(cfg.Dashboards.Gitlab.Repo, cfg.Dashboards.Gitlab.Token)
		if err != nil {
			return nil, err
		}
		app.Trackers = append(app.Trackers, t)
	}
	return app, nil
}

type Tracker interface {
	CreateOrUpdateDashboard(ctx context.Context, body string) error
}
