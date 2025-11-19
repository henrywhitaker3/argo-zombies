package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func IstioBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Kind:    "ConfigMap",
					Version: "v1",
					Name:    "istio-ca-crl",
				},
				{
					Kind:    "ConfigMap",
					Version: "v1",
					Name:    "istio-ca-root-cert",
				},
				{
					Kind:    "ConfigMap",
					Version: "v1",
					Name:    "istio-ip-autoallocate",
				},
				{
					Kind:    "ConfigMap",
					Version: "v1",
					Name:    "istio-leader",
				},
				{
					Kind:    "ConfigMap",
					Version: "v1",
					Name:    "istio-namespace-controller-election",
				},
				{
					Kind:    "ConfigMap",
					Version: "v1",
					Name:    ".*-gatewayapi-configmap",
				},
				{
					Kind:    "Secret",
					Version: "v1",
					Name:    "istio-ca-secret",
				},
				{
					Kind:    "GatewayClass",
					Version: "gateway.networking.k8s.io/v1",
					Name:    "istio",
				},
				{
					Kind:    "GatewayClass",
					Version: "gateway.networking.k8s.io/v1",
					Name:    "istio-waypoint",
				},
				{
					Kind:    "GatewayClass",
					Version: "gateway.networking.k8s.io/v1",
					Name:    "istio-remote",
				},
				{
					Kind:    "GatewayClass",
					Version: "gateway.networking.k8s.io/v1beta1",
					Name:    "istio",
				},
				{
					Kind:    "GatewayClass",
					Version: "gateway.networking.k8s.io/v1beta1",
					Name:    "istio-waypoint",
				},
				{
					Kind:    "GatewayClass",
					Version: "gateway.networking.k8s.io/v1beta1",
					Name:    "istio-remote",
				},
				{
					Kind:    "ClusterRole",
					Version: "rbac.authorization.k8s.io/v1",
					Name:    "kiali",
				},
				{
					Kind:    "ClusterRoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
					Name:    "kiali",
				},
			},
		}
	}
}
