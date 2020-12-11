package config

var (
	config *Config
)

const (
	// DefaultPodNamespace is the default namespace where the Pod Operator is watching.
	DefaultPodNamespace string = "worker"
	// DefaultKubeconfigPath is the default local path of kubeconfig file.
	DefaultKubeconfigPath string = "/home/jinghzhu/.kube/config"
)

type Config struct {
	podNamespace   string
	kubeconfigPath string
}

func (c *Config) GetPodNamespace() string {
	return c.podNamespace
}

func (c *Config) GetKubeconfigPath() string {
	return c.kubeconfigPath
}
