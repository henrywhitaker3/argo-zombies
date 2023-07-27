/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package main

import (
	"fmt"
	"os"

	"github.com/henrywhitaker3/argo-zombies/cmd"
)

func main() {
	cmd.Execute()
}

func getKubeconfigPath() string {
	for i, arg := range os.Args {
		if arg == "--kubeconfig" {
			return os.Args[i+1]
		}
	}

	return fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
}
