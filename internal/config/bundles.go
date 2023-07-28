package config

type bundleFunc func() Exclusions

var bundles map[string]bundleFunc = map[string]bundleFunc{
	"k3s":      k3sBundle(),
	"longhorn": longhornBundle(),
}
