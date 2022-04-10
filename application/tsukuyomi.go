package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"tsukuyomi/controller"
)

func main() {
	// Logger setting
	var err error
	router := gin.Default()
	err = router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}
	router.Use()
	controller.AttachEndpoint(router)
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
