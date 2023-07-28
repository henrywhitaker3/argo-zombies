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
				{
					Name:    "managed-premium",
					Kind:    "StorageClass",
					Version: "storage.k8s.io/v1",
				},
				{
					Name:    "managed-csi-premium",
					Kind:    "StorageClass",
					Version: "storage.k8s.io/v1",
				},
				{
					Name:    "managed-csi",
					Kind:    "StorageClass",
					Version: "storage.k8s.io/v1",
				},
				{
					Name:    "default",
					Kind:    "StorageClass",
					Version: "storage.k8s.io/v1",
				},
				{
					Name:    "azurefile-premium",
					Kind:    "StorageClass",
					Version: "storage.k8s.io/v1",
				},
				{
					Name:    "azurefile-csi-premium",
					Kind:    "StorageClass",
					Version: "storage.k8s.io/v1",
				},
				{
					Name:    "azurefile-csi",
					Kind:    "StorageClass",
					Version: "storage.k8s.io/v1",
				},
				{
					Name:    "azurefile",
					Kind:    "StorageClass",
					Version: "storage.k8s.io/v1",
				},
				{
					Name:    "file.csi.azure.com",
					Kind:    "CSIDriver",
					Version: "storage.k8s.io/v1",
				},
				{
					Name:    "disk.csi.azure.com",
					Kind:    "CSIDriver",
					Version: "storage.k8s.io/v1",
				},
			},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{},
			Namespaces:            []string{},
			Selectors:             []exclusions.ExcludedMetadata{},
		}

		return e
	}
}
