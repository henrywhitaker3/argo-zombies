package detector

import (
	"regexp"

	"github.com/henrywhitaker3/argo-zombies/internal/config"
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Filter returns true if the item should be filtered out
type Filter func(item unstructured.Unstructured) bool

// BuildFilters returns the filters to pass items through
func BuildFilters() []Filter {
	return []Filter{
		ArgoZombiesAnnotationFilter(),
		ArgoLabelFilter(),
		ArgoAnnotationFilter(),
		ArgoRedisSecretFilter(),
		ServiceAccountSecretFilter(),
		HelmSecretFilter(),
		HasOwnerFilter(),
		KubernetesBootstrappingFilter(),
		ExcludedNamespacesFilter(),
		ExcludedSelectorsFilter(),
		ExcludedResourcessFilter(),
		ExcludeStatefulsetPersistentVolumeClaims(),
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

func ArgoAnnotationFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if _, present := item.GetAnnotations()["argocd.argoproj.io/tracking-id"]; present {
			return true
		}
		return false
	}
}

func ArgoRedisSecretFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		return item.GetName() == "argocd-redis" && item.GetKind() == "Secret"
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

func ArgoZombiesAnnotationFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if val, present := item.GetAnnotations()["argo-zombies/ignore"]; present && val == "true" {
			return true
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

func KubernetesBootstrappingFilter() Filter {
	return func(item unstructured.Unstructured) bool {
		if _, present := item.GetLabels()["kubernetes.io/bootstrapping"]; present {
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

func ExcludeStatefulsetPersistentVolumeClaims() Filter {
	return func(item unstructured.Unstructured) bool {
		if item.GetKind() != "PersistentVolumeClaim" {
			return false
		}
		match, err := regexp.MatchString("(.*)-[0-9]+", item.GetName())
		if err != nil {
			return false
		}
		return match
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
				return true
			}
		}
		return false
	}
}

func ExcludedResourcessFilter() Filter {
	argoResources := bundles.ArgoBundle()().Resources
	return func(item unstructured.Unstructured) bool {
		excl := append(config.Cfg.Exclusions.Resources, argoResources...)
		for _, r := range excl {
			if !resourceMatchesNamespace(r.Namespace, item) {
				continue
			}
			if !resourceMatchesAPIVersionKind(r.Version, r.Kind, item) {
				continue
			}
			if resourceMatchesName(r.Name, item) {
				return true
			}
		}
		return false
	}
}

func resourceMatchesNamespace(ns string, item unstructured.Unstructured) bool {
	if item.GetNamespace() == "" || ns == "" {
		return true
	}

	return ns == item.GetNamespace()
}

func resourceMatchesAPIVersionKind(version, kind string, item unstructured.Unstructured) bool {
	return item.GetAPIVersion() == version && item.GetKind() == kind
}

func resourceMatchesName(name string, item unstructured.Unstructured) bool {
	if item.GetName() == name {
		return true
	}

	match, err := regexp.MatchString(`^`+name+`$`, item.GetName())
	if err != nil {
		return false
	}

	return match
}
