package bundles_test

import (
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"github.com/henrywhitaker3/argo-zombies/internal/test"
)

func certManagerTestCases() []testCase {
	return []testCase{
		{
			name: "filters out issuer key secrets",
			item: test.NewItem().
				SetKind("Secret").
				SetName("bongo-issuer-key").
				SetNamespace("cert-manager").
				Item(),
			bundle:       bundles.CertManagerBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out cert-manager-webhook-ca secret",
			item: test.NewItem().
				SetKind("Secret").
				SetName("cert-manager-webhook-ca").
				SetNamespace("cert-manager").
				Item(),
			bundle:       bundles.CertManagerBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out item with cert-manager label",
			item: test.NewItem().
				SetLabels(map[string]string{"controller.cert-manager.io/fao": "true"}).
				Item(),
			bundle:       bundles.CertManagerBundle(),
			shouldReturn: true,
		},
	}
}
