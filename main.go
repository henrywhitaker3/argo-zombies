/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package main

import (
	"os"

	"github.com/henrywhitaker3/argo-zombies/cmd"
	"github.com/henrywhitaker3/argo-zombies/internal/config"
)

func main() {
	path := getConfigPath()

	if err := config.LoadConfig(path); err != nil {
		panic(err)
	}

	// fmt.Println(config.Cfg.Exclusions)

	cmd.Execute()
}

func getConfigPath() string {
	for i, arg := range os.Args {
		if arg == "--config" {
			return os.Args[i+1]
		}
	}

	return ".argo-zombies.yaml"
}
