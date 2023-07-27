package detector

import "k8s.io/apimachinery/pkg/runtime/schema"

func getSkipList() []schema.GroupVersionResource {
	return []schema.GroupVersionResource{
		{
			Version:  "v1",
			Resource: "namespaces",
		},
		{
			Version:  "v1",
			Resource: "events",
		},
		{
			Version:  "v1",
			Resource: "endpoints",
		},
		{
			Version:  "v1",
			Resource: "componentstatuses",
		},
		{
			Version:  "v1",
			Resource: "persistentvolumes",
		},
		{
			Version:  "v1",
			Group:    "storage.k8s.io",
			Resource: "volumeattachments",
		},
		{
			Version:  "v1",
			Resource: "nodes",
		},
		{
			Version:  "v1beta1",
			Group:    "events.k8s.io",
			Resource: "events",
		},
		{
			Version:  "v1",
			Group:    "events.k8s.io",
			Resource: "events",
		},
		{
			Version:  "v1beta1",
			Group:    "metrics.k8s.io",
			Resource: "pods",
		},
		{
			Version:  "v1beta1",
			Group:    "metrics.k8s.io",
			Resource: "nodes",
		},
		{
			Version:  "v1",
			Group:    "coordination.k8s.io",
			Resource: "leases",
		},
		{
			Version:  "v1beta2",
			Group:    "flowcontrol.apiserver.k8s.io",
			Resource: "flowschemas",
		},
		{
			Version:  "v1beta3",
			Group:    "flowcontrol.apiserver.k8s.io",
			Resource: "flowschemas",
		},
		{
			Version:  "v1alpha1",
			Group:    "argoproj.io",
			Resource: "applications",
		},
		{
			Version:  "v1alpha1",
			Group:    "argoproj.io",
			Resource: "applicationsets",
		},
		{
			Version:  "v1alpha1",
			Group:    "argoproj.io",
			Resource: "appprojects",
		},
	}
}
