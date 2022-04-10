package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"net/http"
	"tsukuyomi/controller/webhook/line"
)

func WebhookV1(c *gin.Context) {
	if line.BotError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to initialize 'Line Bot'",
		})
		return
	}
	events, err := line.Bot.ParseRequest(c.Request)
	if err != nil {
		if err != linebot.ErrInvalidSignature {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
		return
	}
	Selector(events)
	c.JSON(http.StatusOK, gin.H{})
}
