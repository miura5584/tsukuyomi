package line

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"os"
	v1 "tsukuyomi/controller/webhook/line/v1"
	v2 "tsukuyomi/controller/webhook/line/v2"
)

var (
	Bot      *linebot.Client
	BotError error
)

func init() {
	Bot, BotError = linebot.New(os.Getenv("LINEBOT_SECRET_KEY"), os.Getenv("LINEBOT_CHANEL_ACCESS_TOKEN"))
}

func AttachEndpoint(webhook *gin.RouterGroup) {
	lineEndpoint := webhook.Group("line")
	{
		lineEndpoint.POST("/v1", v1.WebhookV1)
		lineEndpoint.POST("/v2", v2.WebhookV2)
	}
}
