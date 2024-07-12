/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package cmd

import (
	"github.com/henrywhitaker3/argo-zombies/internal/app"
	"github.com/spf13/cobra"
)

func NewRootCommand(app *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "argo-zombies",
		Short:   "Find kubernetes resources which are not managed by ArgoCD",
		Version: app.Version,
	}

	cmd.PersistentFlags().String("config", ".argo-zombies.yaml", "Path to the argo-zombies config file")
	cmd.PersistentFlags().String("kubeconfig", "~/.kube/config", "Path to the kubeconfig file")

	cmd.AddCommand(NewDetectCommand(app))

	return cmd
}
