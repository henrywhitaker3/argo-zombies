package detector_test

import (
	"testing"

	"github.com/henrywhitaker3/argo-zombies/internal/config"
	"github.com/henrywhitaker3/argo-zombies/internal/detector"
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestItFiltersOutItems(t *testing.T) {
	cfg := &config.Config{
		Exclusions: exclusions.Exclusions{
			Namespaces: []string{"bongo"},
			Selectors: []exclusions.ExcludedMetadata{
				{
					Annotations: map[string]string{
						"exclude-lone": "bingo",
					},
				},
				{
					Labels: map[string]string{
						"exclude-lone": "bingo",
					},
				},
				{
					Annotations: map[string]string{
						"exclude-combo": "bingo",
					},
					Labels: map[string]string{
						"exclude-combo": "bongo",
					},
				},
			},
			Resources: []exclusions.ExcludedResource{
				{
					Version:   "v1",
					Kind:      "Pod",
					Name:      "poddo",
					Namespace: "bongo",
				},
			},
		},
	}

	type test struct {
		name         string
		item         unstructured.Unstructured
		filter       detector.Filter
		shouldReturn bool
	}
	cases := []test{
		{
			name:         "filters out item with argo instance label",
			item:         itemWithArgoInstanceLabel(),
			filter:       detector.ArgoLabelFilter(),
			shouldReturn: true,
		},
		{
			name:         "doesn't filter out item without argo instance label",
			item:         baseItem(),
			filter:       detector.ArgoLabelFilter(),
			shouldReturn: false,
		},
		{
			name:         "filters out item with argo-zombies annotation",
			item:         itemWithArgoZombiesAnnotation(),
			filter:       detector.ArgoZombiesAnnotationFilter(),
			shouldReturn: true,
		},
		{
			name:         "doesn't filters out item without argo-zombies annotation",
			item:         baseItem(),
			filter:       detector.ArgoZombiesAnnotationFilter(),
			shouldReturn: false,
		},
		{
			name:         "filters out argocd-redis secret",
			item:         argoRedisSecret(),
			filter:       detector.ArgoRedisSecretFilter(),
			shouldReturn: true,
		},
		{
			name:         "doesnt filter out argocd-redis base item",
			item:         baseItem(),
			filter:       detector.ArgoRedisSecretFilter(),
			shouldReturn: false,
		},
		{
			name:         "filters out item with owner references",
			item:         itemWithOwnerReferences(),
			filter:       detector.HasOwnerFilter(),
			shouldReturn: true,
		},
		{
			name:         "doesn't filter out item without owner references",
			item:         baseItem(),
			filter:       detector.HasOwnerFilter(),
			shouldReturn: false,
		},
		{
			name:         "filters out service account secret",
			item:         serviceAccountSecret(),
			filter:       detector.ServiceAccountSecretFilter(),
			shouldReturn: true,
		},
		{
			name:         "doesn't filter out non service account sercet",
			item:         argoRedisSecret(),
			filter:       detector.ServiceAccountSecretFilter(),
			shouldReturn: false,
		},
		{
			name:         "filters out helm secret",
			item:         helmSecret(),
			filter:       detector.HelmSecretFilter(),
			shouldReturn: true,
		},
		{
			name:         "doesn't filter out non-helm secret",
			item:         argoRedisSecret(),
			filter:       detector.HelmSecretFilter(),
			shouldReturn: false,
		},
		{
			name:         "filters out a kubernetes bootstrapping labeled item",
			item:         kubernetesBootstrappingItem(),
			filter:       detector.KubernetesBootstrappingFilter(),
			shouldReturn: true,
		},
		{
			name:         "doesn't filters out an item without kubernetes bootstrapping label",
			item:         baseItem(),
			filter:       detector.KubernetesBootstrappingFilter(),
			shouldReturn: false,
		},
		{
			name:         "filters out an item in an exlcuded namespace",
			item:         itemInExcludedNamespace(),
			filter:       detector.ExcludedNamespacesFilter(cfg.Exclusions.Namespaces),
			shouldReturn: true,
		},
		{
			name:         "doesn't filter out an item not in an exlcuded namespace",
			item:         baseItem(),
			filter:       detector.ExcludedNamespacesFilter(cfg.Exclusions.Namespaces),
			shouldReturn: false,
		},
		{
			name:         "filters out item with lone annotation selector",
			item:         itemWithExcludedAnnotation(),
			filter:       detector.ExcludedSelectorsFilter(cfg.Exclusions.Selectors),
			shouldReturn: true,
		},
		{
			name:         "doesn't filter out item with lone annotation selector with wrong value",
			item:         itemWithExcludedAnnotationWrongValue(),
			filter:       detector.ExcludedSelectorsFilter(cfg.Exclusions.Selectors),
			shouldReturn: false,
		},
		{
			name:         "filters out item with excluded label",
			item:         itemWithExcludedLabel(),
			filter:       detector.ExcludedSelectorsFilter(cfg.Exclusions.Selectors),
			shouldReturn: true,
		},
		{
			name:         "doesn't filter out item with excluded label with wrong value",
			item:         itemWithExcludedLabelWrongValue(),
			filter:       detector.ExcludedSelectorsFilter(cfg.Exclusions.Selectors),
			shouldReturn: false,
		},
		{
			name:         "filters out an item with excluded label/anno combo",
			item:         itemWithExcludedAnnoAndLabel(),
			filter:       detector.ExcludedSelectorsFilter(cfg.Exclusions.Selectors),
			shouldReturn: true,
		},
		{
			name:         "doesn't filter out an item with excluded label/anno combo with wrong label value",
			item:         itemWithExcludedAnnoAndLabelWrongLabelValue(),
			filter:       detector.ExcludedSelectorsFilter(cfg.Exclusions.Selectors),
			shouldReturn: false,
		},
		{
			name:         "filters out excluded resource",
			item:         excludedResoureItem(),
			filter:       detector.ExcludedResourcessFilter(cfg.Exclusions.Resources),
			shouldReturn: true,
		},
		{
			name:         "doesn't filter out non-excluded resource",
			item:         baseItem(),
			filter:       detector.ExcludedResourcessFilter(cfg.Exclusions.Resources),
			shouldReturn: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.filter(c.item) != c.shouldReturn {
				t.Fail()
			}
		})
	}
}

func excludedResoureItem() unstructured.Unstructured {
	item := baseItem()
	item.SetName("poddo")
	item.SetNamespace("bongo")
	return item
}

func itemWithExcludedAnnoAndLabelWrongLabelValue() unstructured.Unstructured {
	item := baseItem()
	item.SetLabels(map[string]string{"exclude-combo": "wrongo"})
	item.SetAnnotations(map[string]string{"exclude-combo": "bingo"})
	return item
}

func itemWithExcludedAnnoAndLabel() unstructured.Unstructured {
	item := baseItem()
	item.SetLabels(map[string]string{"exclude-combo": "bongo"})
	item.SetAnnotations(map[string]string{"exclude-combo": "bingo"})
	return item
}

func itemWithExcludedLabelWrongValue() unstructured.Unstructured {
	item := baseItem()
	item.SetLabels(map[string]string{"exclude-lone": "wrongo"})
	return item
}

func itemWithExcludedLabel() unstructured.Unstructured {
	item := baseItem()
	item.SetLabels(map[string]string{"exclude-lone": "bingo"})
	return item
}

func itemWithExcludedAnnotationWrongValue() unstructured.Unstructured {
	item := baseItem()
	item.SetAnnotations(map[string]string{"exclude-lone": "wrongo"})
	return item
}

func itemWithExcludedAnnotation() unstructured.Unstructured {
	item := baseItem()
	item.SetAnnotations(map[string]string{"exclude-lone": "bingo"})
	return item
}

func itemInExcludedNamespace() unstructured.Unstructured {
	item := baseItem()
	item.SetNamespace("bongo")
	return item
}

func kubernetesBootstrappingItem() unstructured.Unstructured {
	item := baseItem()
	item.SetKind("ConfigMap")
	item.SetLabels(map[string]string{"kubernetes.io/bootstrapping": "true"})
	return item
}

func helmSecret() unstructured.Unstructured {
	item := baseItem()
	item.SetKind("Secret")
	item.SetLabels(map[string]string{"owner": "helm"})
	return item
}

func serviceAccountSecret() unstructured.Unstructured {
	item := baseItem()
	item.SetKind("Secret")
	item.SetAnnotations(map[string]string{"kubernetes.io/service-account.name": "bongo"})
	return item
}

func itemWithOwnerReferences() unstructured.Unstructured {
	item := baseItem()
	item.SetOwnerReferences([]v1.OwnerReference{
		{
			APIVersion: "v1",
			Kind:       "Pod",
			Name:       "bongo",
		},
	})
	return item
}

func argoRedisSecret() unstructured.Unstructured {
	item := baseItem()
	item.SetKind("Secret")
	item.SetName("argocd-redis")
	return item
}

func itemWithArgoInstanceLabel() unstructured.Unstructured {
	item := baseItem()
	item.SetLabels(map[string]string{"argocd.argoproj.io/instance": "argocd"})
	return item
}

func itemWithArgoZombiesAnnotation() unstructured.Unstructured {
	item := baseItem()
	item.SetAnnotations(map[string]string{"argo-zombies/ignore": "true"})
	return item
}

func baseItem() unstructured.Unstructured {
	item := unstructured.Unstructured{}
	item.SetName("item")
	item.SetAPIVersion("v1")
	item.SetKind("Pod")
	item.SetNamespace("default")
	return item
}
