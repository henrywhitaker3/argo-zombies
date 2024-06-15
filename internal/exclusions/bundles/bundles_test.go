package bundles_test

import (
	"testing"

	"github.com/henrywhitaker3/argo-zombies/internal/config"
	"github.com/henrywhitaker3/argo-zombies/internal/detector"
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions/bundles"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type testCase struct {
	name         string
	item         unstructured.Unstructured
	bundle       bundles.BundleFunc
	shouldReturn bool
}

func testCases() []testCase {
	cases := []testCase{}
	cases = append(cases, aksTestCases()...)
	return cases
}

func baseItem() unstructured.Unstructured {
	item := unstructured.Unstructured{}
	item.SetAPIVersion("v1")
	item.SetKind("Pod")
	item.SetName("bongo")
	item.SetNamespace("default")
	return item
}

func TestBundleExclusions(t *testing.T) {
	for _, c := range testCases() {
		t.Run(c.name, func(t *testing.T) {
			config.Cfg = &config.Config{
				Exclusions: c.bundle(),
			}
			defer func() { config.Cfg = nil }()
			filtered := false
			if detector.ExcludedNamespacesFilter()(c.item) {
				filtered = true
			}
			if detector.ExcludedResourcessFilter()(c.item) {
				filtered = true
			}
			if detector.ExcludedSelectorsFilter()(c.item) {
				filtered = true
			}
			if c.shouldReturn != filtered {
				t.Fail()
			}
		})
	}
}
