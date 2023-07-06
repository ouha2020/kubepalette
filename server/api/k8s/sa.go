package k8s

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaApi struct {
}

func (*SaApi) GetSaList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "sa",
	})
}
