package k8s

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConfigApi struct {
}

func (*ConfigApi) GetConfigList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "config",
	})
}
