package watcher

import (
	"errors"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CheckPods will firstly get a list of Pods by given list options. Then it will perform the action defined
// by the function parameter to deal with each Pod.
func CheckPods(kubeconfigPath, namespace string, opts *metav1.ListOptions, f func(*corev1.Pod) error) error {
	// Validate opts.
	if opts == nil {
		return errors.New("*metav1.ListOptions is nil in CheckPods")
	}
	// Build client.
	clientConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		fmt.Printf("Fail to init clientConfig in CheckPods: %+v\n", err)

		return err
	}
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		fmt.Printf("Fail to init Kubernetes API client in CheckPods: %+v\n", err)

		return err
	}
	pods, err := clientSet.CoreV1().Pods(namespace).List(*opts)
	if err != nil {
		fmt.Printf("Fail to list Pods: %+v\n", err)

		return err
	}
	fmt.Printf("Successfully list %d Pods\n", len(pods.Items))

	for k, v := range pods.Items {
		fmt.Printf("Ready to check Pod %s\n", v.GetName())
		go f(&pods.Items[k])
	}

	return nil
}

// processBadPendingPod deals with bad Pending Pods.
func processBadPendingPod(pod *corev1.Pod) error {
	fmt.Println("Process bad Pending Pod " + pod.GetName())

	return nil
}
