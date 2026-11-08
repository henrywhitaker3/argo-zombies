package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func MetalLBBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{
				{
					Resource: "frrnodestates",
					Group:    "frrk8s.metallb.io",
					Version:  "v1beta1",
				},
			},
		}
	}
}
