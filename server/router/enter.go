package router

import (
	"question.com/router/k8s"
	"question.com/router/system"
)

type RouterGroup struct {
	SystemRouterGroup system.SystemRouter
	K8sRouterGroup    k8s.K8sRouter
}

var RouterGroupApp = new(RouterGroup)
