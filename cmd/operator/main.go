package main

import (
	"fmt"

	"github.com/jinghzhu/KubernetesPodOperator/pkg/config"
	"github.com/jinghzhu/KubernetesPodOperator/pkg/operator"
	"github.com/jinghzhu/KubernetesPodOperator/pkg/types"
)

func main() {
	fmt.Println("Start Pod Operator")
	cfg := config.GetConfig()
	op, err := operator.New("", cfg.GetKubeconfigPath(), cfg.GetPodNamespace())
	if err != nil {
		panic(err)
	}
	go op.Start()

	types.MainStop = make(chan bool, 1)
	<-types.MainStop
	fmt.Println("Pod Operator is terminated now")
}
