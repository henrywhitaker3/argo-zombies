package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func ArgoBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Name:    "argo(cd){0,}-redis-ha.*",
					Kind:    "Service",
					Version: "v1",
				},
				{
					Name:    "argo(cd){0,}-redis-ha.*",
					Kind:    "RoleBinding",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "argo(cd){0,}-redis-ha.*",
					Kind:    "Role",
					Version: "rbac.authorization.k8s.io/v1",
				},
				{
					Name:    "argo(cd){0,}-redis-ha.*",
					Kind:    "ServiceAccount",
					Version: "v1",
				},
				{
					Name:    "argo(cd){0,}-redis-ha.*",
					Kind:    "ConfigMap",
					Version: "v1",
				},
				{
					Name:    "argo(cd){0,}-redis-ha.*",
					Kind:    "StatefulSet",
					Version: "apps/v1",
				},
				{
					Name:    "argo(cd){0,}-redis-ha.*",
					Kind:    "Deployment",
					Version: "apps/v1",
				},
			},
		}
	}
}
