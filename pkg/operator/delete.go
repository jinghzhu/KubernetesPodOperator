package operator

import (
	"fmt"
	"time"

	"github.com/jinghzhu/KubernetesPodOperator/pkg/config"

	jinghzhuv1client "github.com/jinghzhu/KubernetesCRD/pkg/crd/jinghzhu/v1/client"

	corev1 "k8s.io/api/core/v1"
)

func (op *Operator) onDelete(obj interface{}) {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		fmt.Println("Should be Pod object but encounter others in onDelete")

		return
	}
	copiedPod := pod.DeepCopy()
	fmt.Printf("\nFind a Pod delete event.\n\tName: %s\n\tResource Version: %s\n", copiedPod.GetName(), copiedPod.GetResourceVersion())

	// Deal with ReplicaSet demo HA case.
	time.Sleep(3 * time.Second)
	cfg := config.GetConfig()
	crdClient, err := jinghzhuv1client.NewClient(op.GetContext(), cfg.GetKubeconfigPath(), "crd")
	if err != nil {
		fmt.Printf("Fail to create CRD client in onDelete: %+v\n", err)

		return
	}
	crdInstance, err := crdClient.GetDefault(copiedPod.Labels["crd"])
	if err != nil {
		fmt.Printf("Fail to get CRD instance %s in onDelete: %+v\n", copiedPod.Labels["crd"], err)

		return
	}
	i, l := 0, len(crdInstance.Spec.PodList)
	podName := copiedPod.GetName()
	for ; i < l; i++ {
		if crdInstance.Spec.PodList[i] == podName {
			break
		}
	}
	if i == l {
		return
	}
	crdInstance.Spec.PodList = append(crdInstance.Spec.PodList[:i], crdInstance.Spec.PodList[i+1:]...)
	crdInstance.Spec.Current--
	crdInstance.Status.Message = "Need to do HA"
	_, err = crdClient.PatchSpecAndStatus(copiedPod.Labels["crd"], &crdInstance.Spec, &crdInstance.Status)
	if err != nil {
		fmt.Printf("Fail to patch CRD %s status and spec for Pod %s in onDelete: %+v\n", copiedPod.Labels["crd"], podName, err)
	}
}
