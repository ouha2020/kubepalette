package initializes

import (
	"flag"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"question.com/global"
)

func K8s() {
	println("kubeconfig-----------------")
	kubeconfig := ".kube/config"
	flag.Parse()
	println("okokokok")
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	global.KubeConfigset = clientSet
	println("global.KubeConfigset = clientSet")

}
