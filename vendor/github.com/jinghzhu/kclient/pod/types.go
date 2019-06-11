package pod

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	defaultDeletePeriod int64 = 2
)

// Client helps talk to Kubernetes objects.
type Client struct {
	kubeClient *kubernetes.Clientset
}

// New returns a pointer to Client object. If neither masterUrl or kubeconfigPath are passed in we fallback
// to inClusterConfig. If inClusterConfig fails, we fallback to the default config.
func New(masterURL, kubeconfigPath string) (*Client, error) {
	clientConfig, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigPath)
	if err != nil {
		return nil, err
	}
	clientSet, err := kubernetes.NewForConfig(clientConfig)

	return &Client{
		kubeClient: clientSet,
	}, err
}
