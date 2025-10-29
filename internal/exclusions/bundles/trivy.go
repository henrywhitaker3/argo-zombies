package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func TrivyBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Name:    "scan-vulnerabilityreport-.*",
					Kind:    "Job",
					Version: "batch/v1",
				},
			},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{
				{
					Group:    "aquasecurity.github.io",
					Version:  "v1alpha1",
					Resource: "clustersbomreports",
				},
			},
		}
	}
}
