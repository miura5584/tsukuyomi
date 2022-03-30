package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"tsukuyomi/controller"
	"tsukuyomi/ent"
)

func init() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)
}

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
