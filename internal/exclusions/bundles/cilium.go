package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func CiliumBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{
				{
					Group:    "cilium.io",
					Version:  "v2",
					Resource: "ciliumidentities",
				},
				{
					Group:    "cilium.io",
					Version:  "v2alpha1",
					Resource: "ciliumendpointslices",
				},
			},
			Resources: []exclusions.ExcludedResource{
				{
					Name:      "cilium",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "cilium-config",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "cilium-ca",
					Namespace: "kube-system",
					Kind:      "Secret",
					Version:   "v1",
				},
				{
					Name:      "cilium",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "cilium-operator",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:    "cilium-operator",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "cilium",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "cilium-operator",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "cilium",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
			},
		}
	}
}
