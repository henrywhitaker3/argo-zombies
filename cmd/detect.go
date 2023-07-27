/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package cmd

import (
	"context"

	"github.com/henrywhitaker3/argo-zombies/internal/detector"
	"github.com/henrywhitaker3/argo-zombies/internal/k8s"
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

		return d.Detect(context.Background())
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
