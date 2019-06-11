package operator

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func (op *Operator) onDelete(obj interface{}) {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		fmt.Println("Should be Pod object but encounter others in onDelete")

		return
	}
	copiedPod := pod.DeepCopy()
	fmt.Printf("Find a Pod delete event: %s\n", copiedPod.String())
}
