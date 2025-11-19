package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

func KargoBundle() BundleFunc {
	return func() exclusions.Exclusions {
		return exclusions.Exclusions{
			Resources: []exclusions.ExcludedResource{
				{
					Name:    "kargo-admin",
					Version: "v1",
					Kind:    "ServiceAccount",
				},
				{
					Name:    "kargo-promoter",
					Version: "v1",
					Kind:    "ServiceAccount",
				},
				{
					Name:    "kargo-viewer",
					Version: "v1",
					Kind:    "ServiceAccount",
				},
				{
					Name:    "kargo-admin",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "Role",
				},
				{
					Name:    "kargo-promoter",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "Role",
				},
				{
					Name:    "kargo-viewer",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "Role",
				},
				{
					Name:    "kargo-admin",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "RoleBinding",
				},
				{
					Name:    "kargo-promoter",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "RoleBinding",
				},
				{
					Name:    "kargo-viewer",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "RoleBinding",
				},
				{
					Name:    "kargo-admin-read-secrets",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "RoleBinding",
				},
				{
					Name:    "kargo-project-admin",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "RoleBinding",
				},
				{
					Name:    "kargo-project-secrets-reader",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "RoleBinding",
				},
				{
					Name:    "kargo-controller-read-secrets",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "RoleBinding",
				},
				{
					Name:    "kargo-project-admin-.*",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "ClusterRole",
				},
				{
					Name:    "kargo-project-admin-.*",
					Version: "rbac.authorization.k8s.io/v1",
					Kind:    "ClusterRoleBinding",
				},
			},
			GroupVersionResources: []exclusions.ExcludedGroupVersionResource{
				{
					Group:    "kargo.akuity.io",
					Version:  "v1alpha1",
					Resource: "freights",
				},
			},
		}
	}
}
