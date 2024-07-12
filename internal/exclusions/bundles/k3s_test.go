package bundles_test

import (
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"github.com/henrywhitaker3/argo-zombies/internal/test"
)

func k3sTestCases() []testCase {
	return []testCase{
		{
			name: "filters out k3s coredns service account",
			item: test.NewItem().
				SetName("coredns").
				SetNamespace("kube-system").
				SetKind("ServiceAccount").
				Item(),
			bundle:       bundles.K3sBundle(),
			shouldReturn: true,
		},
	}
}
