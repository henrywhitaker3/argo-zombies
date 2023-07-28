package config

type bundleFunc func() Exclusions

var bundles map[string]bundleFunc = map[string]bundleFunc{
	"k8s":      k8sBundle(),
	"k3s":      k3sBundle(),
	"longhorn": longhornBundle(),
}
