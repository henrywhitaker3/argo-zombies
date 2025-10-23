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
	cases = append(cases, certManagerTestCases()...)
	cases = append(cases, datadogTestCases()...)
	cases = append(cases, ingressNginxTestCases()...)
	cases = append(cases, k3sTestCases()...)
	return cases
}

func TestBundleExclusions(t *testing.T) {
	for _, c := range testCases() {
		t.Run(c.name, func(t *testing.T) {
			cfg := &config.Config{
				Exclusions: c.bundle(),
			}
			filtered := false
			if detector.ExcludedNamespacesFilter(cfg.Exclusions.Namespaces)(c.item) {
				filtered = true
			}
			if detector.ExcludedResourcessFilter(cfg.Exclusions.Resources)(c.item) {
				filtered = true
			}
			if detector.ExcludedSelectorsFilter(cfg.Exclusions.Selectors)(c.item) {
				filtered = true
			}
			if c.shouldReturn != filtered {
				t.Fail()
			}
		})
	}
}
