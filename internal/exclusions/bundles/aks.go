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
					Name:      "bootstrap-token-.*",
					Namespace: "kube-system",
					Kind:      "Secret",
					Version:   "v1",
				},
				{
					Name:      "konnectivity-certs",
					Namespace: "kube-system",
					Kind:      "Secret",
					Version:   "v1",
				},
				{
					Name:      "aks-secrets-store-csi-driver",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "aks-secrets-store-csi-driver",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "aks-secrets-store-provider-azure",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
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
					Name:      "aks-secrets-store-csi-driver",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "aks-secrets-store-csi-driver-windows",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "aks-secrets-store-provider-azure",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "aks-secrets-store-provider-windows",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "aks-ip-masq-agent",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "cloud-node-manager",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "cloud-node-manager-windows",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "csi-azuredisk-node",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "csi-azuredisk-node-win",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "csi-azurefile-node",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "csi-azurefile-node-win",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "kube-proxy",
					Namespace: "kube-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "konnectivity-agent",
					Namespace: "kube-system",
					Kind:      "Deployment",
					Version:   "apps/v1",
				},
				{
					Name:      "metrics-server",
					Namespace: "kube-system",
					Kind:      "Deployment",
					Version:   "apps/v1",
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
					Name:    "aks-secretproviderclasses-rolebinding",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "aks-secretprovidersyncing-rolebinding",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "aks-secretprovidertokenrequest-rolebinding",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "aks-service-rolebinding",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "auto-approve-csr-for-group",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "auto-approve-renewals-for-nodes",
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
					Name:    "create-csrs-for-bootstrapping",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "metrics-server:system:auth-delegator",
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
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
