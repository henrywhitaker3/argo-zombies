package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func AksBundle() BundleFunc {
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
					Name:    "managed",
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
				{
					Name:    "azure-ip-masq-agent-config-reconciled",
					Kind:    "ConfigMap",
					Version: "v1",
				},
				{
					Name:      "overlay-upgrade-data",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "konnectivity-certs",
					Namespace: "kube-system",
					Kind:      "Secret",
					Version:   "v1",
				},
				{
					Name:      "azure-cloud-provider",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "bootstrap-signer",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "cloud-node-manager",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "konnectivity-agent",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "kube-proxy",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "metrics-server",
					Namespace: "kube-system",
					Kind:      "Service",
					Version:   "v1",
				},
				{
					Name:      "konnectivity-agent",
					Namespace: "kube-system",
					Kind:      "NetworkPolicy",
					Version:   "networking.k8s.io/v1",
				},
				{
					Name:      "coredns-pdb",
					Namespace: "kube-system",
					Kind:      "PodDisruptionBudget",
					Version:   "policy/v1",
				},
				{
					Name:      "konnectivity-agent",
					Namespace: "kube-system",
					Kind:      "PodDisruptionBudget",
					Version:   "policy/v1",
				},
				{
					Name:      "metrics-server-pdb",
					Namespace: "kube-system",
					Kind:      "PodDisruptionBudget",
					Version:   "policy/v1",
				},
				{
					Name:    "aks-cluster-admin-binding",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "aks-service-rolebinding",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "cloud-node-manager",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "container-health-read-logs-global",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "metrics-server:system:auth-delegator",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:aks-client-node-proxier",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:azure-cloud-provider",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:azure-cloud-provider-secret-getter",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:coredns-autoscaler",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:kube-proxy",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:metrics-server",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:persistent-volume-binding",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "aks-service",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "cloud-node-manager",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "container-health-log-reader",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:azure-cloud-provider",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:azure-cloud-provider-secret-getter",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:coredns-autoscaler",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:kube-proxy",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:metrics-server",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:persistent-volume-secret-operator",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:      "metrics-server-auth-reader",
					Namespace: "kube-system",
					Kind:      "RoleBinding",
					Version:   "rbac.authorization.k8s.io/v1",
				},
				{
					Name:      "metrics-server-binding",
					Namespace: "kube-system",
					Kind:      "RoleBinding",
					Version:   "rbac.authorization.k8s.io/v1",
				},
				{
					Name:      "system:metrics-server",
					Namespace: "kube-system",
					Kind:      "Role",
					Version:   "rbac.authorization.k8s.io/v1",
				},
				{
					Name:      "metrics-server",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:    "acn-multitenancy-editor-binding",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "appmonitoringconfig-user-global",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "pod-reader-all-namespaces-binding",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:konnectivity-agent-autoscaler",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "acn-multitenancy-editor",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "pod-reader-all-namespaces",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "system:konnectivity-agent-autoscaler",
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:      "nodeNetConfigEditorRoleBinding",
					Namespace: "kube-system",
					Kind:      "RoleBinding",
					Version:   "rbac.authorization.k8s.io/v1",
				},
				{
					Name:      "nodeNetConfigEditorRole",
					Namespace: "kube-system",
					Kind:      "Role",
					Version:   "rbac.authorization.k8s.io/v1",
				},
				{
					Name:      "cluster-autoscaler-status",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "cluster-autoscaler-status",
					Namespace: "kube-system",
					Kind:      "cns-config",
					Version:   "v1",
				},
				{
					Name:      "cluster-autoscaler-status",
					Namespace: "kube-system",
					Kind:      "cns-win-config",
					Version:   "v1",
				},
				{
					Name:      "cluster-autoscaler-status",
					Namespace: "kube-system",
					Kind:      "konnectivity-agent-autoscaler",
					Version:   "v1",
				},
				{
					Name:      "azure-cns",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "konnectivity-agent-autoscaler",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
			},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{},
			Namespaces:            []string{},
			Selectors: []exclusions.ExcludedMetadata{
				{
					Labels: map[string]string{
						"kubernetes.azure.com/managedby": "aks",
					},
				},
			},
		}

		return e
	}
}
