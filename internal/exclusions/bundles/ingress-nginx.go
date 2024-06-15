package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func IngressNginxBundle() BundleFunc {
	return func() exclusions.Exclusions {
		e := exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Name:    "ingress-nginx-admission",
					Kind:    "Secret",
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
