package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func ReflectorBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Name:    "reloader-meta-info",
					Kind:    "ConfigMap",
					Version: "v1",
				},
			},
			Selectors: []exclusions.ExcludedMetadata{
				{
					Annotations: map[string]string{
						"reflector.v1.k8s.emberstack.com/auto-reflects": "True",
					},
				},
				{
					Annotations: map[string]string{
						"reflector.v1.k8s.emberstack.com/reflects": "True",
					},
				},
			},
		}
	}
}
