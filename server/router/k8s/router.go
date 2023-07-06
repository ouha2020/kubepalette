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
	rg.GET("ns", ag.GetNsList)
	rg.GET("sa", ag.GetSaList)
	rg.GET("pod", ag.GetPodList)
	rg.GET("nodeApi", ag.GetNodeList)
	rg.GET("configApi", ag.GetConfigList)
	rg.GET("serviceApi", ag.GetServiceList)
	rg.GET("storageApi", ag.GetStorageList)

}
