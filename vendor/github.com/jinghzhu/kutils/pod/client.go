package pod

import (
	"bytes"
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
)

// GetPods returns a list of Pods by namespace and list options.
func (c *Client) GetPods(ctx context.Context, namespace string, opts *metav1.ListOptions) (*corev1.PodList, error) {
	return c.kubeClient.CoreV1().Pods(namespace).List(ctx, *opts)
}

// GetPod returns the Pod instance by namespace, Pod name and get options.
func (c *Client) GetPod(ctx context.Context, namespace, podName string, opts *metav1.GetOptions) (*corev1.Pod, error) {
	return c.kubeClient.CoreV1().Pods(namespace).Get(ctx, podName, *opts)
}

// IsExist returns false if the Pod doesn't exist in the specific namespace.
func (c *Client) IsExist(ctx context.Context, namespace, podName string) (bool, error) {
	_, err := c.GetPod(ctx, namespace, podName, &metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

// AddPodLabel adds a label field into Pod. If the key already exists, it'll overwrite it.
func (c *Client) AddPodLabel(ctx context.Context, pod *corev1.Pod, key string, value string) (*corev1.Pod, error) {
	if pod.GetLabels() == nil {
		pod.SetLabels(make(map[string]string))
	}
	pod.Labels[key] = value

	return c.UpdatePod(ctx, pod, pod.GetNamespace(), metav1.UpdateOptions{})
}

// AddAnnotation adds a new key-value pair into Pod annotation field.
func (c *Client) AddAnnotation(ctx context.Context, pod *corev1.Pod, key, value string) (*corev1.Pod, error) {
	if pod.GetAnnotations() == nil {
		pod.SetAnnotations(make(map[string]string))
	}
	pod.Annotations[key] = value

	return c.UpdatePod(ctx, pod, pod.GetNamespace(), metav1.UpdateOptions{})
}

// UpdatePod accepts a context, pod and namespace. It returns a pointer to pod and error.
func (c *Client) UpdatePod(ctx context.Context, pod *corev1.Pod, namespace string, opts metav1.UpdateOptions) (*corev1.Pod, error) {
	return c.kubeClient.CoreV1().Pods(namespace).Update(ctx, pod, opts)
}

// DeletePod talks to Kubernetes to delete a Pod by given delete options.
func (c *Client) DeletePod(ctx context.Context, namespace, podName string, opts metav1.DeleteOptions) error {
	return c.kubeClient.CoreV1().Pods(namespace).Delete(ctx, podName, opts)
}

// DeletePodWithCheck deletes the Pod and will start a goroutine in background
// to confirm whether the Pod is successfully deleted.
func (c *Client) DeletePodWithCheck(ctx context.Context, namespace, podName string, opts metav1.DeleteOptions) error {
	if *opts.GracePeriodSeconds < int64(0) {
		opts.GracePeriodSeconds = &defaultDeletePeriod
	}
	err := c.DeletePod(ctx, namespace, podName, opts)
	if err != nil {
		return err
	}
	go c.WaitForDeletion(ctx, namespace, podName, opts.GracePeriodSeconds)

	return nil
}

// WaitForDeletion will wait for a period to check if Pod is deleted.
func (c *Client) WaitForDeletion(ctx context.Context, namespace, podName string, period *int64) error {
	var waitTime time.Duration
	if period != nil {
		waitTime = time.Duration(*period) * time.Second
	} else {
		waitTime = time.Duration(defaultDeletePeriod) * time.Second
	}

	time.Sleep(waitTime) // Wait for gracefully deletion.

	// Check if the Pod is deleted. If it sill exits, we'll force to delete it. Then check it again.
	// This logic will be tried 3 times at max.
	err := wait.Poll(
		time.Duration(defaultDeletePeriod)*time.Second,
		3*time.Duration(defaultDeletePeriod)*time.Second,
		func() (bool, error) {
			exist, err := c.IsExist(ctx, namespace, podName)
			if err != nil {
				return false, err
			} else if exist {
				var gracePeriod int64
				c.DeletePod(ctx, namespace, podName, metav1.DeleteOptions{GracePeriodSeconds: &gracePeriod})

				return false, nil
			}

			return true, nil
		},
	)

	return err
}

// GetLogString returns the log of Pod in string.
func (c *Client) GetLogString(ctx context.Context, namespace, podName string, opts *corev1.PodLogOptions) (string, error) {
	stream, err := c.kubeClient.CoreV1().Pods(namespace).GetLogs(podName, opts).Stream(ctx)
	defer func() {
		if stream != nil {
			stream.Close()
		}
	}()
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)

	return buf.String(), nil
}

// GetEvents returns a EventList object for the given Pod name and list options.
func (c *Client) GetEvents(ctx context.Context, namespace, podName string, opts *metav1.ListOptions) (*corev1.EventList, error) {
	return c.kubeClient.CoreV1().Events(namespace).List(ctx, *opts)
}

// GetVersion returns the version of the the current REST client
func (c *Client) GetVersion() schema.GroupVersion {
	return c.kubeClient.CoreV1().RESTClient().APIVersion()
}
