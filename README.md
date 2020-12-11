# Kubernetes CRD Controller
This repository is to provide a sample about how to create a controller to watch all events (add/update/delete) of Pods so that we can implement the Operator pattern.

I develop it to work with followings together to demo a sample [ReplicaSet](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) scenario:
1. CRD [Jinghzhu v1](https://github.com/jinghzhu/KubernetesCRD)
2. [CRD Controller](https://github.com/jinghzhu/KubernetescrdOperator/)

Assume:
1. Go version > 1.9.0
2. A Kubernetes cluster is available.
3. Kubernetes version >= 1.18.0



# How It Works
There are many articles introducing the mechanism of controller in Kubernetes community. So, I wouldn't pay too much attention on it. In short, it leverage the etcd watch mechanism to catch all events.

In the method `New()` at `pkg/operator/types.go`, we define the major components of Operator:
1. queue - I implements an in-memory queue to process all events in concurrent way. By default, there is only one main goroutine to help the Operator get events from etcd and call add/update/delete handlers. Now, after receiving events, it puts them in the queue which can ensure the concurrent safe.
2. lister
3. informer
4. event handler callback

```go
    op := &Operator{
		kubeClient: client,
		podsQueue:  workqueue.NewNamedRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(100*time.Millisecond, 5*time.Second), "pods"),
		context:    ctx,
	}
	// fieldSelector := labels.Set{"keys": string(nodeName)}.AsSelector()
	ListOption := metav1.ListOptions{}
	op.indexer, op.watcher = cache.NewIndexerInformer(
		// cache.NewListWatchFromClient(client.CoreV1().RESTClient(), "pods", namespace, fieldSelector),
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				options = ListOption
				return op.kubeClient.CoreV1().Pods(namespace).List(ctx, options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				options = ListOption
				return op.kubeClient.CoreV1().Pods(namespace).Watch(ctx, options)
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
```

For details, please see `pkg/operator/operator.go` and `pkg/operator/watcher.go`.



# How It Looks Like
```bash
$ go run cmd/operator/main.go

Init Pod Operator...
Ready to new an Operator object for kubeconfig /home/jinghzhu/.kube/config and namespace worker
Ready to start Pod Operator...
Pod Operator is running...

Find a Pod add event.
        Name: jinghzhu-worker-hl7fc
        Resource Version: 2347337
```