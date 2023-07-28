package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func AksBundle() bundleFunc {
	return func() exclusions.Exclusions {
		e := exclusions.Exclusions{
			Resources:             []exclusions.ExcludedResource{},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{},
			Namespaces:            []string{},
			Selectors:             []exclusions.ExcludedMetadata{},
		}

		return e
	}
}
