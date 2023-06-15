package system

import (
	"github.com/gin-gonic/gin"
	"question.com/api"
)

type SystemRouter struct {
}

func (*SystemRouter) InitSystem(r *gin.Engine) {

	ag := api.ApiGroupApp.SystemApiGroup
	rg := r.Group("/system")
	rg.GET("ping", ag.SystemTest)
}
