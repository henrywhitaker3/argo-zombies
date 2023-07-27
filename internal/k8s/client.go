package k8s

import (
	"os"
	"strings"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClient(kubeConfigPath string) (*kubernetes.Clientset, error) {
	config, err := getConfig(kubeConfigPath)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func NewDynamicClient(kubeConfigPath string) (*dynamic.DynamicClient, error) {
	config, err := getConfig(kubeConfigPath)
	if err != nil {
		return nil, err
	}
	return dynamic.NewForConfig(config)
}

func getConfig(kubeConfigPath string) (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		path, err := resolveConfigPath(kubeConfigPath)
		if err != nil {
			return nil, err
		}
		// Grab it from kubeconfig if we can't get an in-cluster config
		config, err = clientcmd.BuildConfigFromFlags("", path)

		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

func resolveConfigPath(path string) (string, error) {
	if strings.Contains(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = strings.Replace(path, "~", home, -1)
	}

	return path, nil
}
