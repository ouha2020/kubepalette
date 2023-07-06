package k8s

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StorageApi struct {
}

func (*StorageApi) GetStorageList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "node",
	})
}
