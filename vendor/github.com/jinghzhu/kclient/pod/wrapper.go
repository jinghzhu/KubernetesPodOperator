package pod

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetLog returns Pod's log in string. This method is the wrapper of kube client creation and GetLogString.
func GetLog(podName, podNamespace, kubeconfigPath string) (string, error) {
	c, err := New("", kubeconfigPath)
	if err != nil {
		return "", fmt.Errorf("Fail to create Kubernetes client in GetLog: %v", err)
	}
	logOptions := &corev1.PodLogOptions{
		Timestamps: false,
	}

	return c.GetLogString(podNamespace, podName, logOptions)
}

// DeletePod deletes Pod. It accepts Pod name, namespace and delete options.
func DeletePod(podName, podNamespace, kubeconfigPath string, opts *metav1.DeleteOptions) error {
	if opts == nil {
		return fmt.Errorf("*metav1.DeleteOptions is nil in DeletePod")
	}
	c, err := New("", kubeconfigPath)
	if err != nil {
		return fmt.Errorf("Fail to create Kubernetes client in DeletePod: %v", err)
	}

	return c.DeletePod(podNamespace, podName, opts)
}

// DeletePodWithCheck delets Pod and starts a goroutine in background to check the delete operation.
func DeletePodWithCheck(podName, podNamespace, kubeconfigPath string, opts *metav1.DeleteOptions) error {
	if opts == nil {
		return fmt.Errorf("*metav1.DeleteOptions is nil in DeletePod")
	}
	c, err := New("", kubeconfigPath)
	if err != nil {
		return fmt.Errorf("Fail to create Kubernetes client in DeletePodWithCheck: %v", err)
	}

	return c.DeletePodWithCheck(podNamespace, podName, opts)
}

func GetPods(podNamespace, kubeconfigPath string, opts *metav1.ListOptions) (*corev1.PodList, error) {
	if opts == nil {
		return nil, fmt.Errorf("*metav1.ListOptions is nil in GetPods")
	}
	c, err := New("", kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("Fail to create Kubernetes client in GetPods: %v", err)
	}

	return c.GetPods(podNamespace, opts)
}
