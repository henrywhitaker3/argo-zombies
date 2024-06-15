package bundles_test

import (
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"github.com/henrywhitaker3/argo-zombies/internal/test"
)

func aksTestCases() []testCase {
	cases := []testCase{
		{
			name: "filters out aks node validating webhook",
			item: test.NewItem().
				SetName("aks-node-validating-webhook").
				SetKind("ValidatingWebhookConfiguration").
				SetAPIVersion("admissionregistration.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks node mutating webhook",
			item: test.NewItem().
				SetName("aks-node-mutating-webhook").
				SetKind("MutatingWebhookConfiguration").
				SetAPIVersion("admissionregistration.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks webhook admission controller",
			item: test.NewItem().
				SetName("aks-webhook-admission-controller").
				SetKind("MutatingWebhookConfiguration").
				SetAPIVersion("admissionregistration.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks managed premium storage class",
			item: test.NewItem().
				SetName("managed-premium").
				SetKind("StorageClass").
				SetAPIVersion("storage.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks managed csi premium storage class",
			item: test.NewItem().
				SetName("managed-csi-premium").
				SetKind("StorageClass").
				SetAPIVersion("storage.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks managed csi storage class",
			item: test.NewItem().
				SetName("managed-csi").
				SetKind("StorageClass").
				SetAPIVersion("storage.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks default storage class",
			item: test.NewItem().
				SetName("default").
				SetKind("StorageClass").
				SetAPIVersion("storage.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks azurefile-premium storage class",
			item: test.NewItem().
				SetName("azurefile-premium").
				SetKind("StorageClass").
				SetAPIVersion("storage.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks azurefile-csi-premium storage class",
			item: test.NewItem().
				SetName("azurefile-csi-premium").
				SetKind("StorageClass").
				SetAPIVersion("storage.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks azurefile-csi storage class",
			item: test.NewItem().
				SetName("azurefile-csi").
				SetKind("StorageClass").
				SetAPIVersion("storage.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks azurefile storage class",
			item: test.NewItem().
				SetName("azurefile").
				SetKind("StorageClass").
				SetAPIVersion("storage.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks managed storage class",
			item: test.NewItem().
				SetName("managed").
				SetKind("StorageClass").
				SetAPIVersion("storage.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks file csidriver",
			item: test.NewItem().
				SetName("file.csi.azure.com").
				SetKind("CSIDriver").
				SetAPIVersion("storage.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks disk csidriver",
			item: test.NewItem().
				SetName("disk.csi.azure.com").
				SetKind("CSIDriver").
				SetAPIVersion("storage.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks ip masq configmap",
			item: test.NewItem().
				SetName("azure-ip-masq-agent-config-reconciled").
				SetKind("ConfigMap").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks overlay-upgrade-data configmap",
			item: test.NewItem().
				SetName("overlay-upgrade-data").
				SetKind("ConfigMap").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks konnectivity-certs secret",
			item: test.NewItem().
				SetName("konnectivity-certs").
				SetKind("Secret").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks azure-cloud-provider service account",
			item: test.NewItem().
				SetName("azure-cloud-provider").
				SetKind("ServiceAccount").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks bootstrap-signer service account",
			item: test.NewItem().
				SetName("bootstrap-signer").
				SetKind("ServiceAccount").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks cloud-node-manager service account",
			item: test.NewItem().
				SetName("cloud-node-manager").
				SetKind("ServiceAccount").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks konnectivity-agent service account",
			item: test.NewItem().
				SetName("konnectivity-agent").
				SetKind("ServiceAccount").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks kube-proxy service account",
			item: test.NewItem().
				SetName("kube-proxy").
				SetKind("ServiceAccount").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks metrics-server service account",
			item: test.NewItem().
				SetName("metrics-server").
				SetKind("ServiceAccount").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks konnectivity-agent netpol",
			item: test.NewItem().
				SetName("konnectivity-agent").
				SetKind("NetworkPolicy").
				SetAPIVersion("networking.k8s.io/v1").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks coredns-pdb",
			item: test.NewItem().
				SetName("coredns-pdb").
				SetKind("PodDisruptionBudget").
				SetAPIVersion("policy/v1").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks konnectivity-agent pdb",
			item: test.NewItem().
				SetName("konnectivity-agent").
				SetKind("PodDisruptionBudget").
				SetAPIVersion("policy/v1").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks metrics-server pdb",
			item: test.NewItem().
				SetName("metrics-server-pdb").
				SetKind("PodDisruptionBudget").
				SetAPIVersion("policy/v1").
				SetNamespace("kube-system").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks aks-cluster-admin-binding crb",
			item: test.NewItem().
				SetName("aks-cluster-admin-binding").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks aks-service-rolebinding crb",
			item: test.NewItem().
				SetName("aks-service-rolebinding").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks cloud-node-manager crb",
			item: test.NewItem().
				SetName("cloud-node-manager").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks container-health-read-logs-global crb",
			item: test.NewItem().
				SetName("container-health-read-logs-global").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks metrics-server:system:auth-delegator crb",
			item: test.NewItem().
				SetName("metrics-server:system:auth-delegator").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:aks-client-node-proxier crb",
			item: test.NewItem().
				SetName("system:aks-client-node-proxier").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:azure-cloud-provider crb",
			item: test.NewItem().
				SetName("system:azure-cloud-provider").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:azure-cloud-provider-secret-getter crb",
			item: test.NewItem().
				SetName("system:azure-cloud-provider-secret-getter").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:coredns-autoscaler crb",
			item: test.NewItem().
				SetName("system:coredns-autoscaler").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:kube-proxy crb",
			item: test.NewItem().
				SetName("system:kube-proxy").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:metrics-server crb",
			item: test.NewItem().
				SetName("system:metrics-server").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:persistent-volume-binding crb",
			item: test.NewItem().
				SetName("system:persistent-volume-binding").
				SetKind("ClusterRoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks aks-service cr",
			item: test.NewItem().
				SetName("aks-service").
				SetKind("ClusterRole").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks cloud-node-manager crb",
			item: test.NewItem().
				SetName("cloud-node-manager").
				SetKind("ClusterRole").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks container-health-log-reader crb",
			item: test.NewItem().
				SetName("container-health-log-reader").
				SetKind("ClusterRole").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:azure-cloud-provider crb",
			item: test.NewItem().
				SetName("system:azure-cloud-provider").
				SetKind("ClusterRole").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:azure-cloud-provider-secret-getter crb",
			item: test.NewItem().
				SetName("system:azure-cloud-provider-secret-getter").
				SetKind("ClusterRole").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:kube-proxy crb",
			item: test.NewItem().
				SetName("system:kube-proxy").
				SetKind("ClusterRole").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:metrics-server crb",
			item: test.NewItem().
				SetName("system:metrics-server").
				SetKind("ClusterRole").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks system:persistent-volume-secret-operator crb",
			item: test.NewItem().
				SetName("system:persistent-volume-secret-operator").
				SetKind("ClusterRole").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks metrics-server-auth-reader rb",
			item: test.NewItem().
				SetName("metrics-server-auth-reader").
				SetNamespace("kube-system").
				SetKind("RoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks metrics-server-binding rb",
			item: test.NewItem().
				SetName("metrics-server-binding").
				SetNamespace("kube-system").
				SetKind("RoleBinding").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks metrics-server-auth-reader role",
			item: test.NewItem().
				SetName("system:metrics-server").
				SetNamespace("kube-system").
				SetKind("Role").
				SetAPIVersion("rbac.authorization.k8s.io/v1").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out aks metrics-server service account",
			item: test.NewItem().
				SetName("metrics-server").
				SetNamespace("kube-system").
				SetKind("ServiceAccount").
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out item with azure managed by label",
			item: test.NewItem().
				SetName("bongo").
				SetLabels(map[string]string{"kubernetes.azure.com/managedby": "aks"}).
				Item(),
			bundle:       bundles.AksBundle(),
			shouldReturn: true,
		},
	}
	return cases
}
