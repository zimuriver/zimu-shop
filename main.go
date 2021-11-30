package main

import (
	"fmt"
	"zimu-shop/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("First test")
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.GET("api/user/register", controller.Register)
	r.Run()
}
