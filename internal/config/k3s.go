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
				{
					Name:      "coredns",
					Namespace: "kube-system",
					Kind:      "Deployment",
					Version:   "apps/v1",
				},
				{
					Name:      "cluster-dns",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "coredns",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "extension-apiserver-authentication",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "k3s",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "k3s-etcd",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "k3s-etcd-snapshots",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "kube-apiserver-legacy-service-account-token-tracking",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "k3s-serving",
					Namespace: "kube-system",
					Kind:      "Secret",
					Version:   "v1",
				},
				{
					Name:    "clustercidrs-node",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "k3s-cloud-controller-manager",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "k3s-cloud-controller-manager-auth-delegator",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "kube-apiserver-kubelet-admin",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:k3s-controller",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "clustercidrs-node",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "k3s-cloud-controller-manager",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:k3s-controller",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:      "k3s-cloud-controller-manager-authentication-reader",
					Namespace: "kube-system",
					Kind:      "RoleBinding",
					Version:   "rbac.authorization.k8s.io/v1",
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
