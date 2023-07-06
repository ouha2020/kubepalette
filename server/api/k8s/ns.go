package k8s

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NsApi struct {
}

func (*NsApi) GetNsList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ns",
	})
}
