package bundles

import "github.com/henrywhitaker3/argo-zombies/internal/exclusions"

type bundleFunc func() exclusions.Exclusions

var Bundles map[string]bundleFunc = map[string]bundleFunc{
	"k8s":      K8sBundle(),
	"k3s":      K3sBundle(),
	"longhorn": LonghornBundle(),
	"aks":      AksBundle(),
}
