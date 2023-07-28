package config

func k8sBundle() bundleFunc {
	return func() Exclusions {
		e := Exclusions{
			Resources: []ExcludedResource{
				{
					Name:    "kube-root-ca.crt",
					Kind:    "ConfigMap",
					Version: "v1",
				},
				{
					Name:      "extension-apiserver-authentication",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:    "default",
					Kind:    "ServiceAccount",
					Version: "v1",
				},
				{
					Name:    "system-cluster-critical",
					Kind:    "PriorityClass",
					Version: "scheduling.k8s.io/v1",
				},
				{
					Name:    "system-node-critical",
					Kind:    "PriorityClass",
					Version: "scheduling.k8s.io/v1",
				},
				{
					Name:      "kubernetes",
					Namespace: "default",
					Kind:      "EndpointSlice",
					Version:   "discovery.k8s.io/v1",
				},
				{
					Name:      "kubernetes",
					Namespace: "default",
					Kind:      "Service",
					Version:   "v1",
				},
				{
					Name:      "kube-dns",
					Namespace: "kube-system",
					Kind:      "Service",
					Version:   "v1",
				},
			},
			GroupVersionResources: []ExcludedGroupVersionResource{},
			Namespaces:            []string{},
			Selectors:             []ExcludedMetadata{},
		}

		return e
	}
}
