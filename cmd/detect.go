/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package cmd

import (
	"github.com/henrywhitaker3/argo-zombies/internal/config"
	"github.com/henrywhitaker3/argo-zombies/internal/dashboard"
	"github.com/henrywhitaker3/argo-zombies/internal/detector"
	"github.com/henrywhitaker3/argo-zombies/internal/k8s"
	"github.com/henrywhitaker3/argo-zombies/internal/views/detect"
	"github.com/spf13/cobra"
)

// detectCmd represents the detect command
var detectCmd = &cobra.Command{
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

		d := detector.NewDetector(client, dynClient)

		c, err := d.Detect(cmd.Context())
		if err != nil {
			return err
		}

		tbl := detect.NewTable(c)
		tbl.Print()

		if dashboardEnabled(config.Cfg) {
			iB, err := detect.NewMarkdown(c)
			if err != nil {
				return err
			}

			if config.Cfg.Dashboards.Github.Enabled {
				gh := dashboard.NewGithub(cmd.Context(), config.Cfg.Dashboards.Github.Repo, config.Cfg.Dashboards.Github.Token)
				if err := gh.CreateOrUpdateDashboard(iB); err != nil {
					panic(err)
				}
			}
			if config.Cfg.Dashboards.Gitlab.Enabled {
				gl, err := dashboard.NewGitlab(cmd.Context(), config.Cfg.Dashboards.Gitlab.Repo, config.Cfg.Dashboards.Gitlab.Token)
				if err != nil {
					panic(err)
				}
				if err := gl.CreateOrUpdateDashboard(iB); err != nil {
					panic(err)
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func dashboardEnabled(cfg *config.Config) bool {
	return config.Cfg.Dashboards.Github.Enabled || config.Cfg.Dashboards.Gitlab.Enabled
}
