package k8s

import (
	"github.com/gin-gonic/gin"
	"question.com/api"
)

type K8sRouter struct {
}

func (*K8sRouter) InitK8sRouter(r *gin.Engine) {

	ag := api.ApiGroupApp.K8sApiGroup
	rg := r.Group("/k8s")
	rg.GET("ping", ag.GetPodList)
}
