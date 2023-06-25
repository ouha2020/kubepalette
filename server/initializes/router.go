package initializes

import (
	"question.com/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	sr := router.RouterGroupApp.SystemRouterGroup
	kr := router.RouterGroupApp.K8sRouterGroup
	sr.InitSystem(r)
	kr.InitK8sRouter(r)
	return r
}
