package operator

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func (op *Operator) onAdd(obj interface{}) {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		fmt.Println("Should be Pod object but encounter others in onAdd")

		return
	}
	copiedPod := pod.DeepCopy()
	fmt.Printf("Find a Pod add event.\n\tName: %s\n\tResource Version: %s\n", copiedPod.GetName(), copiedPod.GetResourceVersion())
}
