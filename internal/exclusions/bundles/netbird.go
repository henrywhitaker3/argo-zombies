package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func NetbirdBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{
				{
					Group:    "netbird.io",
					Version:  "v1",
					Resource: "nbresources",
				},
			},
		}
	}
}
