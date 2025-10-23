/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package main

import (
	"context"
	"os"

	"github.com/henrywhitaker3/argo-zombies/cmd"
	"github.com/henrywhitaker3/argo-zombies/internal/app"
	"github.com/henrywhitaker3/argo-zombies/internal/config"
	"github.com/henrywhitaker3/argo-zombies/internal/logger"
)

var (
	version = "unknown"
)

func main() {
	path := getConfigPath()

	cfg, err := config.LoadConfig(path)
	if err != nil {
		panic(err)
	}
	app, err := app.New(cfg)
	if err != nil {
		panic(err)
	}
	app.Version = version

	ctx := logger.Wrap(context.Background(), cfg.LogLevel.Level())
	defer logger.Logger(ctx).Sync()

	root := cmd.NewRootCommand(app)
	root.SetContext(ctx)
	if root.Execute() != nil {
		os.Exit(1)
	}
}

func getConfigPath() string {
	for i, arg := range os.Args {
		if arg == "--config" {
			return os.Args[i+1]
		}
	}

	return ".argo-zombies.yaml"
}
