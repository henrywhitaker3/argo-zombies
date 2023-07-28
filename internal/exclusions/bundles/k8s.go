package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func K8sBundle() bundleFunc {
	return func() exclusions.Exclusions {
		e := exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
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
				{
					Name:      "coredns",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "coredns-autoscaler",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "coredns-custom",
					Namespace: "kube-system",
					Kind:      "ConfigMap",
					Version:   "v1",
				},
				{
					Name:      "coredns",
					Namespace: "kube-system",
					Kind:      "Deployment",
					Version:   "apps/v1",
				},
				{
					Name:      "coredns-autoscaler",
					Namespace: "kube-system",
					Kind:      "Deployment",
					Version:   "apps/v1",
				},
				{
					Name:      "coredns",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
				{
					Name:      "coredns-autoscaler",
					Namespace: "kube-system",
					Kind:      "ServiceAccount",
					Version:   "v1",
				},
			},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{},
			Namespaces:            []string{},
			Selectors:             []exclusions.ExcludedMetadata{},
		}

		return e
	}
}
