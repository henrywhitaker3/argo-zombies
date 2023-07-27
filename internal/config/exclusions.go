package config

type Exclusions struct {
	Resources             []ExcludedResource             `yaml:"resources"`
	Namespaces            []string                       `yaml:"namespaces"`
	Selectors             []ExcludedMetadata             `yaml:"selectors"`
	GroupVersionResources []ExcludedGroupVersionResource `yaml:"gvrs"`
}

type ExcludedResource struct {
	Annotations map[string]string `yaml:"annotations"`
	Labels      map[string]string `yaml:"labels"`
	Namespace   string            `yaml:"namespace"`
	Name        string            `yaml:"name"`
	Kind        string            `yaml:"kind"`
	Version     string            `yaml:"version"`
}

type ExcludedGroupVersionResource struct {
	Version  string `yaml:"version"`
	Group    string `yaml:"group"`
	Resource string `yaml:"resource"`
}

type ExcludedMetadata struct {
	Annotations map[string]string `yaml:"annotations"`
	Labels      map[string]string `yaml:"labels"`
}
