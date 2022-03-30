package main

import (
	"github.com/gin-gonic/gin"
	"tsukuyomi/controller"
)

func main() {
	router := gin.Default()
	router.Use()
	router.GET("/", controller.Index)
	router.GET("/ping", controller.Ping)
	webhook := router.Group("/webhook")
	{
		v1 := webhook.Group("/v1")
		{
			v1.POST("/line")
			v1.POST("/discord")
		}

	}
	err := router.Run()
	if err != nil {
		return
	}
}

type EmailLoginRequest struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
