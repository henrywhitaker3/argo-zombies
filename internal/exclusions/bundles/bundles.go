// Package bundles
package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

type BundleFunc func() exclusions.Exclusions

var Bundles map[string]BundleFunc = map[string]BundleFunc{
	"k8s":           K8sBundle(),
	"k3s":           K3sBundle(),
	"longhorn":      LonghornBundle(),
	"aks":           AksBundle(),
	"ingress-nginx": IngressNginxBundle(),
	"cert-manager":  CertManagerBundle(),
	"datadog":       DatadogBundle(),
	"netbird":       NetbirdBundle(),
	"cilium":        CiliumBundle(),
	"kyverno":       KyvernoBundle(),
}
