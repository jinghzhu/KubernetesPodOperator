package operator

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinghzhu/KubernetesPodOperator/pkg/types"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Operator is a watch on Pods.
type Operator struct {
	kubeClient *kubernetes.Clientset
	indexer    cache.Indexer    // A cache of Pods.
	watcher    cache.Controller // Watch changes to all Pods.
	podsQueue  workqueue.RateLimitingInterface
	context    context.Context
}

// New returns an instance of Pod Operator.
func New(masterURL, kubeconfigPath, namespace string) (*Operator, error) {
	fmt.Println("Ready to new an Operator object for kubeconfig " + kubeconfigPath + " and namespace " + namespace)
	clientConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		fmt.Printf(
			"Fail to get RESTClient config by kubeconfig %s and ane namespace %s: %+v\n",
			kubeconfigPath,
			namespace,
			err,
		)

		return nil, err
	}
	client, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		fmt.Printf("Fail to create a new Kubernetes Clientset: %+v\n", err)

		return nil, err
	}

	id, _ := uuid.NewRandom()
	ctx := context.WithValue(types.ContextRoot, "operator-id", id)

	op := &Operator{
		kubeClient: client,
		podsQueue:  workqueue.NewNamedRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(100*time.Millisecond, 5*time.Second), "pods"),
		context:    ctx,
	}
	// fieldSelector := labels.Set{"keys": string(nodeName)}.AsSelector()
	// ListOption is used to get all the marked pod as created by Gulel
	ListOption := metav1.ListOptions{}
	op.indexer, op.watcher = cache.NewIndexerInformer(
		// cache.NewListWatchFromClient(client.CoreV1().RESTClient(), "pods", namespace, fieldSelector),
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				options = ListOption
				return op.kubeClient.CoreV1().Pods(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				options = ListOption
				return op.kubeClient.CoreV1().Pods(namespace).Watch(options)
			},
		},
		&corev1.Pod{},
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    op.onAdd,
			UpdateFunc: op.onUpdate,
			DeleteFunc: op.onDelete,
		},
		cache.Indexers{},
	)

	return op, nil
}

// GetContext returns the context of the Operator instance.
func (op *Operator) GetContext() context.Context {
	return op.context
}
