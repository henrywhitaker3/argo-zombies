package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

// Add exclusions for k3s resources
func K3sBundle() bundleFunc {
	return func() exclusions.Exclusions {
		e := exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Name:      "coredns",
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
					Name:      "cluster-dns",
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
			},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{
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
			Selectors:  []exclusions.ExcludedMetadata{},
		}

		return e
	}
}
