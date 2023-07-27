package detector

import (
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Returns true if the item should be filtered out
type Filter func(item unstructured.Unstructured) bool

// Get the filters to pass items through
func BuildFilters() []Filter {
	return []Filter{
		ArgoLabelFilter(),
		ArgoApplicationFilter(),
		ArgoApplicationSetFilter(),
		ServiceAccountSecretFilter(),
		HelmSecretFilter(),
		HasOwnerFilter(),
		NamespaceFilter(),
		EventFilter(),
		KubernetesBootstrappingFilter(),
		// TODO: add if in config for the longhorn ones
		LonghornBackupFilter(),
		LonghornBackupVolumeFilter(),
		LonghornSettingFilter(),
		LonghornVolumeFilter(),
		LonghornNodeFilter(),
		LonghornBackupTargetFilter(),
		LonghornEngineImageFilter(),
		LonghornSystemBackupFilter(),
		MetricsFilter(),
		NodeFilter(),
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

func ArgoApplicationFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "Application" && strings.Contains(item.GetAPIVersion(), "argoproj.io") {
			return true
		}
		return false
	}
}

func ArgoApplicationSetFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "ApplicationSet" && strings.Contains(item.GetAPIVersion(), "argoproj.io") {
			return true
		}
		return false
	}
}

func EventFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "Event" && (item.GetAPIVersion() == "v1" || strings.Contains(item.GetAPIVersion(), "events.k8s.io")) {
			return true
		}
		return false
	}
}

func NamespaceFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "Namespace" && item.GetAPIVersion() == "v1" {
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

func NodeFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() == "Node" && strings.Contains(item.GetAPIVersion(), "v1") {
			return true
		}
		return false
	}
}

func MetricsFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		return strings.Contains(item.GetAPIVersion(), "metrics.k8s.io")
	}
}
