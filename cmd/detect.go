/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package cmd

import (
	"fmt"

	"github.com/henrywhitaker3/argo-zombies/internal/config"
	"github.com/henrywhitaker3/argo-zombies/internal/dashboard"
	"github.com/henrywhitaker3/argo-zombies/internal/detector"
	"github.com/henrywhitaker3/argo-zombies/internal/k8s"
	"github.com/henrywhitaker3/argo-zombies/internal/views/detect"
	"github.com/spf13/cobra"
	"k8s.io/client-go/dynamic"
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
		// dynClient, err :=

		d := detector.NewDetector(detector.DetectorOpts{
			Client: client,
			DynClient: func() (*dynamic.DynamicClient, error) {
				return k8s.NewDynamicClient(cmd.Flag("kubeconfig").Value.String())
			},
		})

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
				gh, err := dashboard.NewGithub(cmd.Context(), dashboard.GithubOpts{
					Title:             config.Cfg.Dashboards.Title,
					Repo:              config.Cfg.Dashboards.Github.Repo,
					Token:             config.Cfg.Dashboards.Github.Token,
					AppClientID:       config.Cfg.Dashboards.Github.ClientID,
					AppInstallationID: config.Cfg.Dashboards.Github.InstallationID,
					AppPrivateKey:     config.Cfg.Dashboards.Github.PrivateKey,
				})
				if err != nil {
					return fmt.Errorf("build github client: %w", err)
				}
				if err := gh.CreateOrUpdateDashboard(iB); err != nil {
					return err
				}
			}
			if config.Cfg.Dashboards.Gitlab.Enabled {
				gl, err := dashboard.NewGitlab(cmd.Context(), dashboard.GitlabOpts{
					Title: config.Cfg.Dashboards.Title,
					Repo:  config.Cfg.Dashboards.Gitlab.Repo,
					Token: config.Cfg.Dashboards.Gitlab.Token,
				})
				if err != nil {
					return err
				}
				if err := gl.CreateOrUpdateDashboard(iB); err != nil {
					return err
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
