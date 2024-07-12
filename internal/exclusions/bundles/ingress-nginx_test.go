package bundles_test

import (
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"github.com/henrywhitaker3/argo-zombies/internal/test"
)

func ingressNginxTestCases() []testCase {
	return []testCase{
		{
			name: "filters out ingress-nginx-admission secret",
			item: test.NewItem().
				SetName("ingress-nginx-admission").
				SetKind("Secret").
				Item(),
			bundle:       bundles.IngressNginxBundle(),
			shouldReturn: true,
		},
	}
}
