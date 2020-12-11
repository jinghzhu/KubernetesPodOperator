package operator

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func (op *Operator) onUpdate(oldObj, newObj interface{}) {
	old, ok1 := oldObj.(*corev1.Pod)
	new, ok2 := newObj.(*corev1.Pod)
	if !ok1 || !ok2 {
		fmt.Println("Should be Pod object but encounter others in onUpdate")

		return
	}
	oldPod, newPod := old.DeepCopy(), new.DeepCopy()
	if old.ResourceVersion == new.ResourceVersion || old.Status.Phase == new.Status.Phase {
		// Periodic resync will send update events for all known pods.
		// Two different versions of the same pod will always have different RVs.
		// Beside, it currently only watches status change event.
		return
	}
	fmt.Printf(
		"\nFind a Pod update event.\n\tPod Name: %s\n\tOld Status: %s\n\tNew Status: %s\n",
		oldPod.GetName(),
		oldPod.Status.String(),
		newPod.Status.String(),
	)
}
