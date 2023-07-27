package config

func longhornBundle() bundleFunc {
	return func() Exclusions {
		longhorn := Exclusions{
			GroupVersionResources: []ExcludedGroupVersionResource{
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "backingimages",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "backingimages",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "backups",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "backups",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "backuptargets",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "backuptargets",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "backupvolumes",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "backupvolumes",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "engineimages",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "engineimages",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "engines",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "engines",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "instancemanagers",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "instancemanagers",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "nodes",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "nodes",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "orphans",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "recurringjobs",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "recurringjobs",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "replicas",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "replicas",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "settings",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "settings",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "sharemanagers",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "sharemanagers",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "snapshots",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "supportbundles",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "systembackups",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "systemrestores",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "volumeattachments",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta1",
					Resource: "volumes",
				},
				{
					Group:    "longhorn.io",
					Version:  "v1beta2",
					Resource: "volumes",
				},
			},
			Namespaces: []string{},
			Selectors:  []ExcludedMetadata{},
		}

		return longhorn
	}
}
