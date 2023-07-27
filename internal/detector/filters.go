package detector

import (
	"strings"

	"github.com/henrywhitaker3/argo-zombies/internal/config"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Returns true if the item should be filtered out
type Filter func(item unstructured.Unstructured) bool

// Get the filters to pass items through
func BuildFilters() []Filter {
	return []Filter{
		ArgoLabelFilter(),
		ServiceAccountSecretFilter(),
		HelmSecretFilter(),
		HasOwnerFilter(),
		KubernetesBootstrappingFilter(),
		ExcludedNamespacesFilter(),
		ExcludedSelectorsFilter(),
		// TODO: add if in config for the longhorn ones
		LonghornBackupFilter(),
		LonghornBackupVolumeFilter(),
		LonghornSettingFilter(),
		LonghornVolumeFilter(),
		LonghornNodeFilter(),
		LonghornBackupTargetFilter(),
		LonghornEngineImageFilter(),
		LonghornSystemBackupFilter(),
		CertManagerSecretFilter(),
	}
}

func ArgoLabelFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if _, present := item.GetLabels()["argocd.argoproj.io/instance"]; present {
			return true
		}
		return false
	}
}

func HasOwnerFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if refs := item.GetOwnerReferences(); len(refs) > 0 {
			return true
		}
		return false
	}
}

func ServiceAccountSecretFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "Secret" && item.GetAPIVersion() == "v1" {
			if _, present := item.GetAnnotations()["kubernetes.io/service-account.name"]; present {
				return true
			}
		}
		return false
	}
}

func HelmSecretFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "Secret" && item.GetAPIVersion() == "v1" {
			if val, present := item.GetLabels()["owner"]; present && val == "helm" {
				return true
			}
		}
		return false
	}
}

func CertManagerSecretFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "Secret" && item.GetAPIVersion() == "v1" {
			if val, present := item.GetLabels()["controller.cert-manager.io/fao"]; present && val == "true" {
				return true
			}
		}
		return false
	}
}

func KubernetesBootstrappingFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if _, present := item.GetLabels()["kubernetes.io/bootstrapping"]; present {
			return true
		}
		return false
	}
}

func LonghornBackupFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "Backup" && strings.Contains(item.GetAPIVersion(), "longhorn.io") {
			return true
		}
		return false
	}
}

func LonghornSystemBackupFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "SystemBackup" && strings.Contains(item.GetAPIVersion(), "longhorn.io") {
			return true
		}
		return false
	}
}

func LonghornSettingFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "Setting" && strings.Contains(item.GetAPIVersion(), "longhorn.io") {
			return true
		}
		return false
	}
}

func LonghornVolumeFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "Volume" && strings.Contains(item.GetAPIVersion(), "longhorn.io") {
			return true
		}
		return false
	}
}

func LonghornBackupVolumeFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "BackupVolume" && strings.Contains(item.GetAPIVersion(), "longhorn.io") {
			return true
		}
		return false
	}
}

func LonghornNodeFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "Node" && strings.Contains(item.GetAPIVersion(), "longhorn.io") {
			return true
		}
		return false
	}
}

func LonghornBackupTargetFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "BackupTarget" && strings.Contains(item.GetAPIVersion(), "longhorn.io") {
			return true
		}
		return false
	}
}

func LonghornEngineImageFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "EngineImage" && strings.Contains(item.GetAPIVersion(), "longhorn.io") {
			return true
		}
		return false
	}
}

func ExcludedNamespacesFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		for _, ns := range config.Cfg.Exclusions.Namespaces {
			if item.GetNamespace() == ns {
				return true
			}
		}
		return false
	}
}

func ExcludedSelectorsFilter() Filter {
	matches := func(selectors map[string]string, metadata map[string]string) bool {
		hit := 0
		for key, val := range selectors {
			if v, present := metadata[key]; present && val == v {
				hit++
			}
		}
		return len(selectors) == hit && hit != 0
	}

	return func(item unstructured.Unstructured) bool {
		for _, selector := range config.Cfg.Exclusions.Selectors {
			passed := false

			if len(item.GetAnnotations()) > 0 && len(selector.Annotations) > 0 {
				passed = matches(selector.Annotations, item.GetAnnotations())
			}
			if len(item.GetLabels()) > 0 && len(selector.Labels) > 0 {
				passed = matches(selector.Labels, item.GetLabels())
			}

			if passed {
				// fmt.Printf("Excluding argo label %s/%s/%s\n", item.GetAPIVersion(), item.GetKind(), item.GetName())
				return true
			}
		}
		return false
	}
}
