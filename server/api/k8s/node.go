package k8s

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NodeApi struct {
}

func (*NodeApi) GetNodeList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "node",
	})
}
