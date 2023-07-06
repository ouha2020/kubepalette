package k8s

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceApi struct {
}

func (*ServiceApi) GetServiceList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "service",
	})
}
