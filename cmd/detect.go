/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package cmd

import (
	"github.com/henrywhitaker3/argo-zombies/internal/app"
	"github.com/henrywhitaker3/argo-zombies/internal/detector"
	"github.com/henrywhitaker3/argo-zombies/internal/k8s"
	"github.com/henrywhitaker3/argo-zombies/internal/views/detect"
	"github.com/spf13/cobra"
)

func NewDetectCommand(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "detect",
		Short: "Detect and list any resources not managed by ArgoCD",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := k8s.NewClient(cmd.Flag("kubeconfig").Value.String())
			if err != nil {
				return err
			}
			dynClient, err := k8s.NewDynamicClient(cmd.Flag("kubeconfig").Value.String())
			if err != nil {
				return err
			}

			d := detector.NewDetector(client, dynClient, app.Config)

			c, err := d.Detect(cmd.Context())
			if err != nil {
				return err
			}

			tbl := detect.NewTable(c)
			tbl.Print()

			iB, err := detect.NewMarkdown(c)
			if err != nil {
				return err
			}
			for _, tracker := range app.Trackers {
				tracker.CreateOrUpdateDashboard(cmd.Context(), iB)
			}

			return nil
		},
	}
}
