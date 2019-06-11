package main

import (
	"fmt"
	"os"

	"github.com/jinghzhu/KubernetesPodOperator/pkg/operator"
	"github.com/jinghzhu/KubernetesPodOperator/pkg/types"
)

func main() {
	fmt.Println("Start Pod Operator")
	op, err := operator.New("", os.Getenv("KUBECONFIG"), os.Getenv("NAMESPACE"))
	if err != nil {
		panic(err)
	}
	go op.Start()

	types.MainStop = make(chan bool, 1)
	<-types.MainStop
	fmt.Println("Pod Operator is terminated now")
}
