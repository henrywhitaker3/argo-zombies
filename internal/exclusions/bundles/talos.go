package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func TalosBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			Selectors: []exclusions.ExcludedMetadata{
				{
					Annotations: map[string]string{
						"config.k8s.io/owning-inventory": "talos-bootstrap-manifests-inventory",
					},
				},
			},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{
				{
					Group:    "certificates.k8s.io",
					Version:  "v1",
					Resource: "clustertrustbundles",
				},
			},
			Resources: []exclusions.ExcludedResource{
				{
					Name:      "talos-ipv4",
					Namespace: "default",
					Kind:      "EndpointSlice",
					Version:   "discovery.k8s.io/v1",
				},
				{
					Name:      "talos-bootstrap-manifests-inventory",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "talos",
					Namespace: "default",
					Kind:      "Service",
					Version:   "v1",
				},
			},
		}
	}
}
