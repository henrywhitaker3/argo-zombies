package detector

import (
	"github.com/henrywhitaker3/argo-zombies/internal/config"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func getSkipList() []schema.GroupVersionResource {
	list := []schema.GroupVersionResource{
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
			Version:  "v1beta2",
			Group:    "flowcontrol.apiserver.k8s.io",
			Resource: "prioritylevelconfigurations",
		},
		{
			Version:  "v1beta3",
			Group:    "flowcontrol.apiserver.k8s.io",
			Resource: "prioritylevelconfigurations",
		},
		{
			Version:  "v1",
			Group:    "flowcontrol.apiserver.k8s.io",
			Resource: "prioritylevelconfigurations",
		},
		{
			Version:  "v1",
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
		{
			Version:  "v1",
			Group:    "apiregistration.k8s.io",
			Resource: "apiservices",
		},
		{
			Version:  "v1",
			Group:    "node.k8s.io",
			Resource: "runtimeclasses",
		},
	}

	for _, gvr := range config.Cfg.Exclusions.GroupVersionResources {
		list = append(list, schema.GroupVersionResource{
			Version:  gvr.Version,
			Group:    gvr.Group,
			Resource: gvr.Resource,
		})
	}

	return list
}
