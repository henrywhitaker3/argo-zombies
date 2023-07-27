package config

type Exclusions struct {
	Resources  []ExcludedResource `yaml:"resources"`
	Namespaces []string           `yaml:"namespaces"`
	Selectors  []ExcludedMetadata `yaml:"selectors"`
}

type ExcludedResource struct {
	Annotations map[string]string `yaml:"annotations"`
	Labels      map[string]string `yaml:"labels"`
	Namespace   string            `yaml:"namespace"`
	Name        string            `yaml:"name"`
	Kind        string            `yaml:"kind"`
	Version     string            `yaml:"version"`
}

type ExcludedMetadata struct {
	Annotations map[string]string `yaml:"annotations"`
	Labels      map[string]string `yaml:"labels"`
}
