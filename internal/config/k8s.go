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
			},
			GroupVersionResources: []ExcludedGroupVersionResource{},
			Namespaces:            []string{},
			Selectors:             []ExcludedMetadata{},
		}

		return e
	}
}
