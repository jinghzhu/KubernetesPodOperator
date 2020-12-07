package watcher

import (
	"context"
	"fmt"
	"time"

	"github.com/jinghzhu/KubernetesPodOperator/pkg/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	utilspod "github.com/jinghzhu/kutils/pod"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CheckPods will firstly get a list of Pods by given list options. Then it will perform the action defined
// by the function parameter to deal with each Pod.
func CheckPods(
	ctx context.Context,
	kubeconfigPath,
	namespace string,
	opts metav1.ListOptions,
	f func(context.Context, *corev1.Pod) error) error {
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
	pods, err := clientSet.CoreV1().Pods(namespace).List(ctx, opts)
	if err != nil {
		fmt.Printf("Fail to list Pods: %+v\n", err)

		return err
	}
	fmt.Printf("Successfully list %d Pods\n", len(pods.Items))

	for k, v := range pods.Items {
		fmt.Printf("Ready to check Pod %s\n", v.GetName())
		go f(ctx, &pods.Items[k])
	}

	return nil
}

// processBadPendingPod deals with bad Pending Pods.
func processBadPendingPod(ctx context.Context, pod *corev1.Pod) error {
	fmt.Println("Process bad Pending Pod " + pod.GetName())

	cfg := config.GetConfig()
	err := utilspod.DeletePodWithCheck(
		ctx,
		pod.GetName(),
		pod.GetNamespace(),
		cfg.GetKubeconfigPath(),
		metav1.DeleteOptions{
			GracePeriodSeconds: &deleteGracePeriod,
		},
	)
	if err != nil {
		fmt.Println("Fail to delete bad Pending Pod " + pod.GetName())

		// Try again later.
		go func() {
			time.Sleep(5 * time.Minute)
			utilspod.DeletePodWithCheck(
				ctx,
				pod.GetName(),
				pod.GetNamespace(),
				cfg.GetKubeconfigPath(),
				metav1.DeleteOptions{
					GracePeriodSeconds: &deleteGracePeriod,
				},
			)
		}()
	}

	return err
}
