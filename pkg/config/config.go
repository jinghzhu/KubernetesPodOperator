package config

import "os"

func init() {
	initConfig()
}

func initConfig() {
	config = &Config{}
	config.podNamespace = os.Getenv("PO_NAMESPACE")
	if config.podNamespace == "" {
		config.podNamespace = DefaultPodNamespace
	}

	config.kubeconfigPath = os.Getenv("PO_KUBECONFIG")
	if config.kubeconfigPath == "" {
		config.kubeconfigPath = DefaultKubeconfigPath
	}
}

func GetConfig() *Config {
	return config
}
