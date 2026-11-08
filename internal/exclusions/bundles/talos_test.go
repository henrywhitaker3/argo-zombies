package bundles_test

import (
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"github.com/henrywhitaker3/argo-zombies/internal/test"
)

func talosTestCases() []testCase {
	return []testCase{
		{
			name: "filters out Talos bootstrap manifest owned resource",
			item: test.NewItem().
				SetAnnotations(map[string]string{
					"config.k8s.io/owning-inventory": "talos-bootstrap-manifests-inventory",
				}).
				Item(),
			bundle:       bundles.TalosBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out Talos IPv4 endpoint slice",
			item: test.NewItem().
				SetName("talos-ipv4").
				SetNamespace("default").
				SetKind("EndpointSlice").
				SetAPIVersion("discovery.k8s.io/v1").
				Item(),
			bundle:       bundles.TalosBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out Talos bootstrap manifest inventory",
			item: test.NewItem().
				SetName("talos-bootstrap-manifests-inventory").
				SetNamespace("kube-system").
				SetKind("ConfigMap").
				Item(),
			bundle:       bundles.TalosBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out Talos service",
			item: test.NewItem().
				SetName("talos").
				SetNamespace("default").
				SetKind("Service").
				Item(),
			bundle:       bundles.TalosBundle(),
			shouldReturn: true,
		},
	}
}
