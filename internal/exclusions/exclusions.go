package exclusions

type Exclusions struct {
	Resources             []ExcludedResource             `yaml:"resources"`
	Namespaces            []string                       `yaml:"namespaces"`
	Selectors             []ExcludedMetadata             `yaml:"selectors"`
	GroupVersionResources []ExcludedGroupVersionResource `yaml:"gvrs"`
	Bundles               []string                       `yaml:"bundles"`
}

type ExcludedResource struct {
	Version   string `yaml:"version"`
	Kind      string `yaml:"kind"`
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
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
