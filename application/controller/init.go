package controller

import (
	"github.com/gin-gonic/gin"
	"tsukuyomi/controller/webhook/discord"
	"tsukuyomi/controller/webhook/line"
)

func AttachEndpoint(router *gin.Engine) {
	router.GET("/", Index)
	router.GET("version", Version)
	webhook := router.Group("/webhook")
	{
		line.AttachEndpoint(webhook)
		discord.AttachEndpoint(webhook)
	}
}
