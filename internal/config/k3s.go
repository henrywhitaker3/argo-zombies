package config

// Add exclusions for k3s resources
func k3sBundle() bundleFunc {
	return func() Exclusions {
		e := Exclusions{
			Resources: []ExcludedResource{
				{
					Name:      "attachdetach-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "certificate-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "clusterrole-aggregation-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "coredns",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "cronjob-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "daemon-set-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "deployment-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "disruption-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "endpoint-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "endpointslice-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "endpointslicemirroring-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "ephemeral-volume-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "expand-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "generic-garbage-collector",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "horizontal-pod-autoscaler",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "job-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "namespace-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "node-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "persistent-volume-binder",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "pod-garbage-collector",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "pv-protection-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "pvc-protection-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "replicaset-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "replication-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "resourcequota-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "root-ca-cert-publisher",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "service-account-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "statefulset-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "token-cleaner",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "ttl-after-finished-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "ttl-controller",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      ".*node-password.k3s",
					Namespace: "kube-system",
					Kind:      "Secret",
					Version:   "v1",
				},
			},
			GroupVersionResources: []ExcludedGroupVersionResource{
				{
					Group:    "k3s.cattle.io",
					Version:  "v1",
					Resource: "addons",
				},
				{
					Group:    "helm.cattle.io",
					Version:  "v1",
					Resource: "helmchartconfigs",
				},
				{
					Group:    "helm.cattle.io",
					Version:  "v1",
					Resource: "helmcharts",
				},
			},
			Namespaces: []string{},
			Selectors:  []ExcludedMetadata{},
		}

		return e
	}
}
