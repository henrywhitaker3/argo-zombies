package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func DatadogBundle() bundleFunc {
	return func() exclusions.Exclusions {
		e := exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Name:    "datadog-cluster-id",
					Kind:    "ConfigMap",
					Version: "v1",
				},
				{
					Name:    "datadog-leader-election",
					Kind:    "ConfigMap",
					Version: "v1",
				},
				{
					Name:    "datadogtoken",
					Kind:    "ConfigMap",
					Version: "v1",
				},
			},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{},
			Namespaces:            []string{},
			Selectors:             []exclusions.ExcludedMetadata{},
		}

		return e
	}
}
