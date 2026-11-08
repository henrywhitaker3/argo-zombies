package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func GithubActionsBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			Selectors: []exclusions.ExcludedMetadata{
				{
					Labels: map[string]string{
						"app.kubernetes.io/part-of": "gha-runner-scale-set",
					},
				},
			},
		}
	}
}
