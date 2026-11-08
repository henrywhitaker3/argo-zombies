package bundles_test

import (
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"github.com/henrywhitaker3/argo-zombies/internal/test"
)

func k8sTestCases() []testCase {
	return []testCase{
		{
			name: "filters out Kubernetes device taint eviction controller",
			item: test.NewItem().
				SetName("device-taint-eviction-controller").
				SetNamespace("kube-system").
				SetKind("ServiceAccount").
				Item(),
			bundle:       bundles.K8sBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out Kubernetes API server ClusterTrustBundle publisher",
			item: test.NewItem().
				SetName("kube-apiserver-serving-clustertrustbundle-publisher").
				SetNamespace("kube-system").
				SetKind("ServiceAccount").
				Item(),
			bundle:       bundles.K8sBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out Kubernetes legacy service account token tracking",
			item: test.NewItem().
				SetName("kube-apiserver-legacy-service-account-token-tracking").
				SetNamespace("kube-system").
				SetKind("ConfigMap").
				Item(),
			bundle:       bundles.K8sBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out Kubernetes pod certificate request cleaner",
			item: test.NewItem().
				SetName("podcertificaterequestcleaner").
				SetNamespace("kube-system").
				SetKind("ServiceAccount").
				Item(),
			bundle:       bundles.K8sBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out Kubernetes storage version migrator controller",
			item: test.NewItem().
				SetName("storage-version-migrator-controller").
				SetNamespace("kube-system").
				SetKind("ServiceAccount").
				Item(),
			bundle:       bundles.K8sBundle(),
			shouldReturn: true,
		},
	}
}
