package operator

import (
	"fmt"
	"os"

	"github.com/jinghzhu/KubernetesPodOperator/pkg/types"
	"github.com/jinghzhu/KubernetesPodOperator/pkg/watcher"
)

// Start starts an Operator to watch Pod lifecycle.
func (op *Operator) Start() {
	fmt.Println("Ready to start Pod Operator...")
	types.POStop = make(chan bool, 1)
	go op.watch(types.POStop)
	<-types.POStop
	fmt.Println("Pod Operator is existing...")
}

// Stop terminates the Pod Operator.
func (op *Operator) Stop() {
	types.POStop <- true
}

func (op *Operator) watch(stopCh <-chan bool) {
	ctx := op.GetContext()
	go op.watcher.Run(ctx.Done())
	// For Pods in Pending status.
	go watcher.PendingPodsWatcher(
		ctx,
		os.Getenv("NAMESPACE"),
		types.PendingWatchInterval,
	)
}
