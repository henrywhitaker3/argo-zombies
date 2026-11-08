package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func EnvoyGatewayBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Name:    "envoy",
					Kind:    "Secret",
					Version: "v1",
				},
				{
					Name:    "envoy-gateway",
					Kind:    "Secret",
					Version: "v1",
				},
				{
					Name:    "envoy-oidc-hmac",
					Kind:    "Secret",
					Version: "v1",
				},
				{
					Name:    "envoy-rate-limit",
					Kind:    "Secret",
					Version: "v1",
				},
			},
		}
	}
}
