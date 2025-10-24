package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func CiliumBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{
				{
					Group:    "cilium.io",
					Version:  "v2",
					Resource: "ciliumidentities",
				},
				{
					Group:    "cilium.io",
					Version:  "v2alpha1",
					Resource: "ciliumendpointslices",
				},
			},
		}
	}
}
