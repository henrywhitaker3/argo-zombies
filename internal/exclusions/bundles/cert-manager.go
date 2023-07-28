package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func CertManagerBundle() bundleFunc {
	return func() exclusions.Exclusions {
		e := exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Name:      ".*-issuer-key",
					Namespace: "cert-manager",
					Kind:      "Secret",
					Version:   "v1",
				},
				{
					Name:      "cert-manager-webhook-ca",
					Namespace: "cert-manager",
					Kind:      "Secret",
					Version:   "v1",
				},
			},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{},
			Namespaces:            []string{},
			Selectors: []exclusions.ExcludedMetadata{
				{
					Labels: map[string]string{
						"controller.cert-manager.io/fao": "true",
					},
				},
			},
		}

		return e
	}
}
