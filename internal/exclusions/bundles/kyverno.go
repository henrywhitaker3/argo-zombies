package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func KyvernoBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Name:    "kyverno-policy-mutating-webhook-cfg",
					Kind:    "MutatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "kyverno-resource-mutating-webhook-cfg",
					Kind:    "MutatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "kyverno-verify-mutating-webhook-cfg",
					Kind:    "MutatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "kyverno-cel-exception-validating-webhook-cfg",
					Kind:    "ValidatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "kyverno-cleanup-validating-webhook-cfg",
					Kind:    "ValidatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "kyverno-exception-validating-webhook-cfg",
					Kind:    "ValidatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "kyverno-global-context-validating-webhook-cfg",
					Kind:    "ValidatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "kyverno-policy-validating-webhook-cfg",
					Kind:    "ValidatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "kyverno-resource-validating-webhook-cfg",
					Kind:    "ValidatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "kyverno-ttl-validating-webhook-cfg",
					Kind:    "ValidatingWebhookConfiguration",
					Version: "admissionregistration.k8s.io/v1",
				},
				{
					Name:    "kyverno-cleanup-controller.kyverno.svc.kyverno-tls-ca",
					Kind:    "Secret",
					Version: "v1",
				},
				{
					Name:    "kyverno-cleanup-controller.kyverno.svc.kyverno-tls-pair",
					Kind:    "Secret",
					Version: "v1",
				},
				{
					Name:    "kyverno-svc.kyverno.svc.kyverno-tls-ca",
					Kind:    "Secret",
					Version: "v1",
				},
				{
					Name:    "kyverno-svc.kyverno.svc.kyverno-tls-pair",
					Kind:    "Secret",
					Version: "v1",
				},
			},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{
				{
					Group:    "wgpolicyk8s.io",
					Version:  "v1alpha2",
					Resource: "policyreports",
				},
				{
					Group:    "reports.kyverno.io",
					Version:  "v1",
					Resource: "ephemeralreports",
				},
			},
		}
	}
}
