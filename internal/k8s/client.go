package k8s

import (
	"os"
	"strings"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func init() {
	rest.SetDefaultWarningHandler(rest.NoWarnings{})
}

type ConnectionParams struct {
	KubeconfigPath           string
	URL                      string
	ClientKeyPath            string
	ClientCertificatePath    string
	ClusterCACertificatePath string
}

func NewClient(params ConnectionParams) (*kubernetes.Clientset, error) {
	config, err := getConfig(params)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func NewDynamicClient(params ConnectionParams) (*dynamic.DynamicClient, error) {
	config, err := getConfig(params)
	if err != nil {
		return nil, err
	}
	return dynamic.NewForConfig(config)
}

func getConfig(params ConnectionParams) (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		builder := buildKubeconfigConfig
		if params.URL != "" {
			builder = buildDefinedConfig
		}

		// Grab it from kubeconfig if we can't get an in-cluster config
		config, err = builder(params)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

func buildDefinedConfig(params ConnectionParams) (*rest.Config, error) {
	return clientcmd.BuildConfigFromKubeconfigGetter("", func() (*api.Config, error) {
		return &api.Config{
			Clusters: map[string]*api.Cluster{
				"cluster": {
					Server:               params.URL,
					CertificateAuthority: params.ClusterCACertificatePath,
				},
			},
			Contexts: map[string]*api.Context{
				"cluster": {
					Cluster:  "cluster",
					AuthInfo: "cluster",
				},
			},
			AuthInfos: map[string]*api.AuthInfo{
				"cluster": {
					ClientCertificate: params.ClientCertificatePath,
					ClientKey:         params.ClientKeyPath,
				},
			},
			CurrentContext: "cluster",
		}, nil
	})
}

func buildKubeconfigConfig(params ConnectionParams) (*rest.Config, error) {
	path, err := resolveConfigPath(params.KubeconfigPath)
	if err != nil {
		return nil, err
	}
	return clientcmd.BuildConfigFromFlags("", path)
}

func resolveConfigPath(path string) (string, error) {
	if strings.Contains(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = strings.ReplaceAll(path, "~", home)
	}

	return path, nil
}
