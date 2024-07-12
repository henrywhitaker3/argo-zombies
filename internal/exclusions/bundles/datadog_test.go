package bundles_test

import (
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"github.com/henrywhitaker3/argo-zombies/internal/test"
)

func datadogTestCases() []testCase {
	return []testCase{
		{
			name: "filters out datadog-cluster-id cm",
			item: test.NewItem().
				SetName("datadog-cluster-id").
				SetKind("ConfigMap").
				Item(),
			bundle:       bundles.DatadogBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out datadog-leader-election cm",
			item: test.NewItem().
				SetName("datadog-leader-election").
				SetKind("ConfigMap").
				Item(),
			bundle:       bundles.DatadogBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out datadogtoken cm",
			item: test.NewItem().
				SetName("datadogtoken").
				SetKind("ConfigMap").
				Item(),
			bundle:       bundles.DatadogBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out webhook-certificate secret",
			item: test.NewItem().
				SetName("webhook-certificate").
				SetKind("Secret").
				Item(),
			bundle:       bundles.DatadogBundle(),
			shouldReturn: true,
		},
		{
			name: "filters out datadog-webhook",
			item: test.NewItem().
				SetName("datadog-webhook").
				SetKind("MutatingWebhookConfiguration").
				SetAPIVersion("admissionregistration.k8s.io/v1").
				Item(),
			bundle:       bundles.DatadogBundle(),
			shouldReturn: true,
		},
	}
}
