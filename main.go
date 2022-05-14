package main

import (
	"github.com/krisdiano/autolabel/pkg"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//1. config
	//2. client
	//3. informer
	//4. add event handler
	//5. informer.Start

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			panic("can't get config")
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic("can't create client")
	}

	factory := informers.NewSharedInformerFactory(clientset, 0)
	podInformer := factory.Core().V1().Pods()

	controller := pkg.NewController(clientset, podInformer)
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)

	controller.Run(stopCh)

}
