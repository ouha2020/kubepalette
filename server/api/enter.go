package api

import (
	"question.com/api/k8s"
	"question.com/api/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	K8sApiGroup    k8s.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
