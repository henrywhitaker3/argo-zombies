package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func LonghornBundle() BundleFunc {
	return func() exclusions.Exclusions {
		longhorn := exclusions.Exclusions{
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "backingimages",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "backingimages",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "backups",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "backups",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "backuptargets",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "backuptargets",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "backupvolumes",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "backupvolumes",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "engineimages",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "engineimages",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "engines",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "engines",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "instancemanagers",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "instancemanagers",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "nodes",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "nodes",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "orphans",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "recurringjobs",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "recurringjobs",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "replicas",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "replicas",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "settings",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "settings",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "sharemanagers",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "sharemanagers",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "snapshots",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "supportbundles",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "systembackups",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "systemrestores",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "volumeattachments",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "volumes",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "volumes",
				},
			},
			Namespaces: []string{},
			Selectors:  []exclusions.ExcludedMetadata{},
			Resources: []exclusions.ExcludedResource{
				{
					Name:      "csi-attacher",
					Namespace: "longhorn-system",
					Kind:      "Service",
					Version:   "v1",
				},
				{
					Name:      "csi-provisioner",
					Namespace: "longhorn-system",
					Kind:      "Service",
					Version:   "v1",
				},
				{
					Name:      "csi-resizer",
					Namespace: "longhorn-system",
					Kind:      "Service",
					Version:   "v1",
				},
				{
					Name:      "csi-snapshotter",
					Namespace: "longhorn-system",
					Kind:      "Service",
					Version:   "v1",
				},
				{
					Name:      "csi-attacher",
					Namespace: "longhorn-system",
					Kind:      "Deployment",
					Version:   "apps/v1",
				},
				{
					Name:      "csi-provisioner",
					Namespace: "longhorn-system",
					Kind:      "Deployment",
					Version:   "apps/v1",
				},
				{
					Name:      "csi-resizer",
					Namespace: "longhorn-system",
					Kind:      "Deployment",
					Version:   "apps/v1",
				},
				{
					Name:      "csi-snapshotter",
					Namespace: "longhorn-system",
					Kind:      "Deployment",
					Version:   "apps/v1",
				},
				{
					Name:      "csi-attacher",
					Namespace: "longhorn-system",
					Kind:      "PodDisruptionBudget",
					Version:   "policy/v1",
				},
				{
					Name:      "csi-provisioner",
					Namespace: "longhorn-system",
					Kind:      "PodDisruptionBudget",
					Version:   "policy/v1",
				},
				{
					Name:      "instance-manager-.*",
					Namespace: "longhorn-system",
					Kind:      "PodDisruptionBudget",
					Version:   "policy/v1",
				},
				{
					Name:      "longhorn-csi-plugin",
					Namespace: "longhorn-system",
					Kind:      "DaemonSet",
					Version:   "apps/v1",
				},
				{
					Name:      "longhorn-webhook-ca",
					Namespace: "longhorn-system",
					Kind:      "Secret",
					Version:   "v1",
				},
				{
					Name:      "longhorn-webhook-tls",
					Namespace: "longhorn-system",
					Kind:      "Secret",
					Version:   "v1",
				},
				{
					Name:      "recovery-backend-share-manager-.*",
					Namespace: "longhorn-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:    "driver.longhorn.io",
					Kind:    "CSIDriver",
					Version: "storage.k8s.io/v1",
				},
				{
					Name:    "longhorn",
					Kind:    "StorageClass",
					Version: "storage.k8s.io/v1",
				},
				{
					Name:    "longhorn-webhook-mutator",
					Kind:    "MutatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "longhorn-webhook-validator",
					Kind:    "ValidatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
			},
		}

		return longhorn
	}
}
