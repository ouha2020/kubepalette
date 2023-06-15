package initializes

import (
	"question.com/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	sr := router.RouterGroupApp.SystemRouterGroup
	sr.InitSystem(r)
	return r
}
