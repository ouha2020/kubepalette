package global

import (
	"k8s.io/client-go/kubernetes"
	"question.com/config"
)

var (
	CONF          config.Server
	KubeConfigset *kubernetes.Clientset
)
