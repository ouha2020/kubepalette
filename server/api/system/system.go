package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemApi struct {
}

func (*SystemApi) SystemTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
