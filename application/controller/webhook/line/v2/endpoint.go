package v2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebhookV2(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
