package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func AksBundle() bundleFunc {
	return func() exclusions.Exclusions {
		e := exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Name:    "aks-node-validating-webhook",
					Kind:    "ValidatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "aks-node-mutating-webhook",
					Kind:    "MutatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "aks-webhook-admission-controller",
					Kind:    "MutatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
			},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{},
			Namespaces:            []string{},
			Selectors:             []exclusions.ExcludedMetadata{},
		}

		return e
	}
}
